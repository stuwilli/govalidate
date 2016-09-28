package rules

import (
	"fmt"
	"strconv"

	"github.com/stuwilli/govalidate/helper"
)

//GreaterThan Passes if the data is a float/int and is greater than the specified integer.
// Note that this is *not* a greater than or equals check, and this only comapres
// floats/ints to a predefined integer specified in your tag.
// Fails if the data is not a float/int or the data is less than or equals the comparator
func GreaterThan(data ValidationData) error {
	v, err := helper.ToFloat64(data.Value)
	if err != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not numeric",
			Message:        data.Message,
		}
	}

	// We should always be provided with a length to validate against
	if len(data.Args) == 0 {
		return fmt.Errorf("No argument found in the validation struct (eg 'GreaterThan:5')")
	}

	// Typecast our argument and test
	var min float64
	if min, err = strconv.ParseFloat(data.Args[0], 64); err != nil {
		return err
	}

	if v < min {
		return ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf("must be greater than %f", min),
			Message:        data.Message,
		}
	}

	return nil
}
