package utils

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore"
	"time"
)

type Claims struct {
	Namespace string    `json:"name"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (c Claims) Valid() error {
	if time.Now().After(c.ExpiresAt) {
		return authcore.ErrExpiredToken
	}

	if len(c.Namespace) < 5 {
		return authcore.ErrInvalidMissingName
	}
	return nil
}
