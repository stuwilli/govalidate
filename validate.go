package validate

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/stuwilli/govalidate/rules"
)

//ValidationError ...
type ValidationError struct {
	Err map[string]string `json:"error"`
}

func (ve *ValidationError) addFailure(field, msg string) {

	// a := []rune(field)
	// a[0] = unicode.ToLower(a[0])
	// field = string(a)
	ve.Err[field] = msg
}

//Merge ...
func (ve *ValidationError) Merge(other ValidationError) {

	for k, v := range other.Err {
		ve.Err[k] = v
	}
}

//Error ...
func (ve ValidationError) Error() string {

	var str = "The following errors occured during validation: "
	for _, e := range ve.Err {
		str += e + ". "
	}
	return str
}

//Errors ...
func (ve ValidationError) Errors() map[string]string {

	return ve.Err
}

//Run ...
func Run(object interface{}, fieldsSlice ...string) error {

	pass := true
	err := ValidationError{Err: make(map[string]string)}

	// If we have been passed a slice of fields to valiate - to check only a
	// subset of fields - change the slice into a map for O(1) lookups instead
	// of O(n).
	fields := map[string]struct{}{}
	for _, field := range fieldsSlice {
		fields[field] = struct{}{}
	}

	// If we're passed a pointer to a struct we need to dereference the pointer before
	// inspecting its tags
	value := reflect.ValueOf(object)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Iterate through each field of the struct and validate
	typ := value.Type() // A Type's Field method returns StructFields
	for i := 0; i < value.NumField(); i++ {
		var validateTag string
		var validateError error
		var fieldName string

		// Is this an anonymous struct? If so, we also need to validate on this.
		if typ.Field(i).Anonymous == true {
			if anonErr := Run(value.Field(i).Interface(), fieldsSlice...); anonErr != nil {
				// The validation failed: set pass to false and merge the anonymous struct's
				// validation errors with our current validation error above to give a complete
				// error message.
				pass = false

				// A non validation error occurred: return this immediately
				if _, ok := anonErr.(ValidationError); !ok {
					return anonErr
				}

				err.Merge(anonErr.(ValidationError))
			}
		}

		if len(fields) > 0 {
			// We're only checking for a subset of fields; if this field isn't
			// included in the subset of fields to validate we can skip.
			if _, ok := fields[typ.Field(i).Name]; !ok {
				continue
			}
		}

		if fieldName = typ.Field(i).Tag.Get("json"); fieldName != "" {
			fieldName = strings.Split(fieldName, ",")[0]
		} else {
			fieldName = typ.Field(i).Name
		}

		if validateTag = typ.Field(i).Tag.Get("validate"); validateTag == "" {
			continue
		}

		// Validate this particular field against the options in our tag
		if validateError = validateField(value.Field(i).Interface(), fieldName, validateTag); validateError == nil {
			continue
		}

		// If there was no validation rule defined for the given tag return
		// that error immediately.
		if _, ok := validateError.(rules.ErrNoValidationMethod); ok {
			return validateError
		}

		pass = false
		err.addFailure(fieldName, validateError.Error())
	}

	if pass {
		return nil
	}

	return err
}

//CastError ...
func CastError(err error) ValidationError {

	return err.(ValidationError)
}

var rxRegexp = regexp.MustCompile(`Regexp:\/.+/`)
var rxMessage = regexp.MustCompile(`Message:.+(,\s)?`)

// Takes a field's value and the validation tag and applies each check
// until either a test fails or all tests pass.
func validateField(data interface{}, fieldName, tag string) (err error) {
	// Grab any custom messages
	var message string
	message = rxMessage.FindString(tag)
	tag = rxMessage.ReplaceAllString(tag, "")
	message = strings.Replace(message, "Message:", "", -1)

	// A tag can specify multiple validation rules which are delimited via ','.
	// However, because we allow regular expressions we can't split the tag field
	// via all commas to find our validation rules: we need to extract the regular expression
	// first (in case it specifies a comma), and *then* run through our validation rules.
	if match := rxRegexp.FindString(tag); match != "" {
		// If we fail validating the regexp we can break here
		if err := validateRule(data, fieldName, match, message); err != nil {
			return err
		}
		// Now we need to replace our regular expression from the tag list to continue
		// normally.
		tag = rxRegexp.ReplaceAllString(tag, "")
	}

	for tag != "" {
		var next string

		i := strings.Index(tag, ",")
		if i >= 0 {
			tag, next = tag[:i], tag[i+1:]
		}

		if err := validateRule(data, fieldName, tag, message); err != nil {
			return err
		}

		// Continue with the next tag
		tag = next
	}

	return nil
}

// Given a validation rule from a tag, run the associated validation methods and return
// the result.
func validateRule(data interface{}, fieldName, rule, message string) error {

	var args []string
	var err error
	var method rules.ValidatorFunc
	// Remove any preceeding spaces from comma separated tags
	rule = strings.TrimLeft(rule, " ")

	// If the rule is empty we don't need to process anything. This only happens
	// if we have a regex followed by another rule:
	//   `validate:"Regexp:/.+/, NotEmpty"`
	// Becomes:
	//   `validate:", NotEmpty"`
	// After processing in validateField()
	if rule == "" {
		return nil
	}

	// rule is the method we want to call. If it has a colon we need to further
	// process the rule to extract arguments to our validation method.
	i := strings.Index(rule, ":")
	if i > 0 {
		var arg string
		rule, arg = rule[:i], rule[i+1:]
		args = append(args, arg)
	}

	// Attempt to validate the data using methods registered with the rules
	// sub package
	if method, err = rules.Get(rule); err != nil {
		return err
	}

	var out = rules.ValidationData{
		Field:   fieldName,
		Value:   data,
		Args:    args,
		Message: message,
	}
	return method(out)
}
