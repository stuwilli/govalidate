package alphanumeric

import (
	"regexp"

	"github.com/amasses/govalidate/helper"
	"github.com/amasses/govalidate/rules"
)

func init() {
	rules.Add("Alphanumeric", Alphanumeric)
}

// Validates that a string only contains alphabetic or numeric characters
func Alphanumeric(data rules.ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString(v) {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "contains non-alphanumeric characters",
			Message:        data.Message,
		}
	}

	return nil
}
