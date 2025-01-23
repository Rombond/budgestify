package token

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(login string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   login,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_KEY")))
}

func IsTokenValid(tokenHeader string) (bool, error) {
	if tokenHeader == "" {
		return false, jwt.ErrTokenUnverifiable
	}

	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
	if tokenString == tokenHeader {
		return false, jwt.ErrTokenMalformed
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("API_KEY")), nil
		},
	)
	return token.Valid, err
}
