package dtos

import "errors"

var (
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrInvalidCustomerKey = errors.New("invalid customer key")
)
