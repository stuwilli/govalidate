package rules

import (
	"regexp"

	"github.com/stuwilli/govalidate/helper"
)

//Alphanumeric Validates that a string only contains alphabetic or numeric characters
func Alphanumeric(data ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString(v) {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "contains non-alphanumeric characters",
			Message:        data.Message,
		}
	}

	return nil
}
