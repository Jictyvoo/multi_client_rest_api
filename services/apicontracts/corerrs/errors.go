package corerrs

import "errors"

var (
	ErrInvalidPhone         = errors.New("phone number provided is invalid")
	ErrContactAlreadyExists = errors.New("contact already exists")
)
