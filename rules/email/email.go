package email

import (
	"regexp"

	"github.com/amasses/govalidate/helper"
	"github.com/amasses/govalidate/rules"
)

func init() {
	rules.Add("Email", Email)
}

func Email(data rules.ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if IsEmail(v) {
		return
	}

	return rules.ErrInvalid{
		ValidationData: data,
		Failure:        "is not a valid email address",
		Message:        data.Message,
	}
}

func IsEmail(str string) bool {
	return regexp.MustCompile(`(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`).MatchString(str)
}
