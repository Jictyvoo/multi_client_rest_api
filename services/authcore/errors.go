package authcore

import "errors"

var (
	ErrExpiredToken       = errors.New("expired token")
	ErrInvalidMissingKey  = errors.New("invalid token: missing key")
	ErrInvalidMissingName = errors.New("invalid token: missing name")
)
