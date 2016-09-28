package rules

import "fmt"

//ErrInvalid ...
type ErrInvalid struct {
	ValidationData
	Failure string
	Message string
}

func (t ErrInvalid) Error() string {
	if len(t.Message) > 0 {
		return t.Message
	}
	return fmt.Sprintf("Field '%s' %s", t.Field, t.Failure)
}

//ErrNoValidationMethod ...
type ErrNoValidationMethod struct {
	Tag string
}

func (t ErrNoValidationMethod) Error() string {
	return fmt.Sprintf("No validation method for '%s' has been registered", t.Tag)
}
