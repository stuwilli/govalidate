package rules

import "github.com/stuwilli/govalidate/helper"

//NotEmpty Checks whether a string is empty.
// Passes if the data is a non-empty string. Fails if the data isn't a string, or the data is an empty string.
func NotEmpty(data ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}
	if v == "" {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is empty",
			Message:        data.Message,
		}
	}
	return nil
}
