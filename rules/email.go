package rules

import (
	"regexp"

	"github.com/stuwilli/govalidate/helper"
)

//Email Validates email addresses
func Email(data ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if IsEmail(v) {
		return
	}

	return ErrInvalid{
		ValidationData: data,
		Failure:        "is not a valid email address",
		Message:        data.Message,
	}
}

//IsEmail Email address validation
func IsEmail(str string) bool {
	return regexp.MustCompile(`(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`).MatchString(str)
}
