package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	// set .env secret key
	secret := os.Getenv("JWT_SECRET_KEY")
	// set expiration count in minutes for JWT_SECRET_KEY
	minCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRATION_MINUTES"))
	// create new claims
	claims := jwt.MapClaims{}
	// set public claims
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minCount)).Unix()
	// create new jwt token w/ claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// generate token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %e", err)
	}
	return t, nil
}
