package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(level int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["level"] = level
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
