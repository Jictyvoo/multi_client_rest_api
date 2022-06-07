package dtos

import "errors"

var (
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrInvalidCustomerKey    = errors.New("invalid customer key")
	ErrExpiredToken          = errors.New("expired token")
	ErrInvalidMissingName    = errors.New("invalid token: missing name")
	ErrInvalidMissingUUID    = errors.New("invalid token: missing uuid")
)
