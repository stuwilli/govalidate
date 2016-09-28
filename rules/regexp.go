package rules

import (
	"fmt"
	"regexp"

	"github.com/stuwilli/govalidate/helper"
)

//Regexp Validates that a string only contains alphabetic characters
func Regexp(data ValidationData) (err error) {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	// We should always be provided with a length to validate against
	if len(data.Args) == 0 {
		return fmt.Errorf("No argument found in the validation struct (eg 'Regexp:/^\\s+$/')")
	}

	// Remove the trailing slashes from our regex string. Regexps must be enclosed
	// within two "/" characters.
	re := data.Args[0]
	re = re[1 : len(re)-1]
	if regexp.MustCompile(re).MatchString(v) == false {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "doesn't match regular expression",
			Message:        data.Message,
		}
	}

	return nil
}
