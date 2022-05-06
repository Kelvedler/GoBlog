package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(issuer string, subject string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	now := time.Now().Unix()
	claims := &jwt.StandardClaims{
		ExpiresAt: now + (60 * 60 * 24),
		IssuedAt:  now,
		Issuer:    issuer,
		Subject:   subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
