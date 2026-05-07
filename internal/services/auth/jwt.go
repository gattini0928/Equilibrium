package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int, expirationSec int64) (string, error) {
	expiration := time.Second * time.Duration(expirationSec)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(secret string, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	userIDFloat := claims["userID"].(float64)
	return int(userIDFloat), nil
}

