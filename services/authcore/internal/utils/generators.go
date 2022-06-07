package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenExpirationTime = time.Hour * 72
	passwordCost        = 11
)

func GenerateClaimToken(namespace string) Claims {
	return Claims{
		Namespace: namespace,
		ExpiresAt: time.Now().Add(tokenExpirationTime),
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
