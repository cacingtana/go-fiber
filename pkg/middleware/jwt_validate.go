package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenData struct {
	Expires int64
}

func ExtractTokenData(f *fiber.Ctx) (*TokenData, error) {
	token, err := verifyToken(f)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		return &TokenData{
			Expires: expires,
		}, nil
	}
	return nil, err

}

func verifyToken(f *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(f)
	token, err := jwt.Parse(tokenString, jwtKey)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(f *fiber.Ctx) string {
	bearer := f.Get("Authorization")

	onlyToken := strings.Split(bearer, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func jwtKey(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
