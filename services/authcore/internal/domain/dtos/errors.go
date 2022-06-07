package dtos

import "errors"

var (
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrInvalidCustomerKey = errors.New("invalid customer key")
	ErrExpiredToken       = errors.New("expired token")
	ErrInvalidMissingName = errors.New("invalid token: missing name")
)
