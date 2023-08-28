package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwtSecretKey  = "eu amo minha namorada"
	jwtExpiration = 24 * time.Hour
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(jwtExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
