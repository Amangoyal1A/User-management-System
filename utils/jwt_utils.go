package utils

import (
	"time"
	"user-management/config"

	"github.com/golang-jwt/jwt/v4"
)
var jwtSecret = config.LoadConfig().JWTKey;

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}