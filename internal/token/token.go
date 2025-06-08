package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprint(id),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_KEY")))
}

func IsTokenValid(tokenHeader string, id int) (bool, error) {
	if tokenHeader == "" {
		return false, jwt.ErrTokenUnverifiable
	}

	// Supprime le préfixe "Bearer " si présent
	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
	if tokenString == tokenHeader {
		return false, jwt.ErrTokenMalformed
	}

	// Définir une structure pour les claims
	claims := &jwt.RegisteredClaims{}

	// Parse le token et les claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("API_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	// Vérifie si le token est valide et si le subject correspond à l'id fourni
	if token.Valid && claims.Subject == fmt.Sprint(id) {
		return true, nil
	}

	return false, jwt.ErrTokenInvalidClaims
}
