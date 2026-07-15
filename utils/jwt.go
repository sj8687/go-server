package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("super-secret-key")

func GenerateJWT(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret)
}