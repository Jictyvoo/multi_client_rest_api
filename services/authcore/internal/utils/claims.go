package utils

import (
	"github.com/google/uuid"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"time"
)

type Claims struct {
	Name      string `json:"name"`
	Uuid      uuid.UUID
	ExpiresAt time.Time `json:"expires_at"`
}

func (c Claims) Valid() error {
	if time.Now().After(c.ExpiresAt) {
		return dtos.ErrExpiredToken
	}

	if len(c.Uuid.String()) < 5 {
		return dtos.ErrInvalidMissingUUID
	}

	if len(c.Name) < 5 {
		return dtos.ErrInvalidMissingName
	}
	return nil
}
