package rules

import (
	"regexp"

	"github.com/stuwilli/govalidate/helper"
)

//Alpha Validates that a string only contains alphabetic characters
func Alpha(data ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if regexp.MustCompile(`[^a-zA-Z]+`).MatchString(v) {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "contains non-alphabetic characters",
			Message:        data.Message,
		}
	}

	return nil
}
