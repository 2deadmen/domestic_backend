package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a JWT token
func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
