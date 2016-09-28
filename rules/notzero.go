package rules

import "github.com/stuwilli/govalidate/helper"

//NotZero Checks whether a float or int type is 0. This could mean the data is above *or* below 0.
// Fails if the data isn't a float/int type, or the data is exactly 0.
func NotZero(data ValidationData) error {
	v, err := helper.ToFloat64(data.Value)
	if err != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not numeric",
			Message:        data.Message,
		}
	}

	if v == 0 {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is 0",
			Message:        data.Message,
		}
	}

	return nil
}
