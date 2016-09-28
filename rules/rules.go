package rules

import "fmt"

// This maps all validation tags to the corresponding validation methods
var rules map[string]ValidatorFunc

func init() {
	rules = map[string]ValidatorFunc{}
	Add("Alpha", Alpha)
	Add("Alphanumeric", Alphanumeric)
	Add("Email", Email)
	Add("GreaterThan", GreaterThan)
	Add("Length", Length)
	Add("LessThan", LessThan)
	Add("MaxLength", MaxLength)
	Add("MinLength", MinLength)
	Add("NotEmpty", NotEmpty)
	Add("NotZero", NotZero)
	Add("NotZeroTime", NotZeroTime)
	Add("Regexp", Regexp)
	Add("URL", URL)
	Add("UUID", UUID)
}

//ValidationData ...
type ValidationData struct {
	// The name of the field being validated
	Field string

	// The override message for validation failure
	Message string

	// The value of the struct field being validated
	Value interface{}

	// Arguments from the validation tags. For example, in the following
	// definition Args will will contain a single "5":
	//
	// struct {
	//     Age `validate:"GreaterThan:5"`
	// }
	//
	// Unfortunately, due to the nature of tags these will always be strings.
	Args []string
}

//ValidatorFunc All validation methods must return an ErrInvalid error type if the data
// is invalid, or nil if the data is valid
type ValidatorFunc func(ValidationData) error

// Add a new validation method for a given struct tag. If a validation method
// already exists this will return an error
func Add(tag string, method ValidatorFunc) (err error) {

	if _, ok := rules[tag]; ok {
		return fmt.Errorf("Validation method for '%s' already exists", tag)
	}

	rules[tag] = method
	return
}

//Get Return a registered validation method for a given tag
func Get(tag string) (method ValidatorFunc, err error) {

	m, ok := rules[tag]
	if !ok {
		return nil, ErrNoValidationMethod{Tag: tag}
	}

	return m, nil
}
