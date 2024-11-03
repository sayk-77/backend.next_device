package tools

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	JwtKey := os.Getenv("JWT_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JwtKey))
}
