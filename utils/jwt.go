package utils

import (
	"santrikoding/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// getJwtSecret retrieves the JWT secret from the configuration.
func GetJwtSecret() []byte {
	secret := config.AppConfig.GetString("jwt.secret")
	if secret == "" {
		panic("JWT secret is not set in the configuration")
	}
	return []byte(secret)
}

// GenerateJWT generates a JWT token for the given user ID.
func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJwtSecret())
}
