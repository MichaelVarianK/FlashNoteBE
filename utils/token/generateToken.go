package utils_token

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(email string) (string, error) {
	var SecretKey = []byte(os.Getenv("JWT_SECRET"))

	if len(SecretKey) == 0 {
		log.Fatal("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
}

func GenerateRefreshToken() string {
	return uuid.New().String()
}
