package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenExpirationTime = time.Hour * 72
	passwordCost        = 11
)

func GenerateClaimToken(name string) Claims {
	token, _ := uuid.NewRandom()
	return Claims{
		Name:      name,
		ExpiresAt: time.Now().Add(tokenExpirationTime),
		Uuid:      token,
	}
}

func CreateJWT(claims Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(secretKey))
}

func EncryptPassword(password string) ([]byte, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return []byte{}, err
	}
	return encryptedPassword, nil
}
