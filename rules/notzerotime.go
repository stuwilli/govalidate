package rules

import "time"

//NotZeroTime Checks whether a float or int type is 0. This could mean the data is above *or* below 0.
// Fails if the data isn't a float/int type, or the data is exactly 0.
func NotZeroTime(data ValidationData) error {
	if _, ok := data.Value.(time.Time); !ok {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a Time type",
			Message:        data.Message,
		}
	}

	if data.Value.(time.Time).Equal(time.Time{}) == true {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "has a zero value",
			Message:        data.Message,
		}
	}

	return nil
}
