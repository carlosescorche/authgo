package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(userID primitive.ObjectID) (string, error) {
	token, err := insert(userID)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrTokenInsert, err)
	}

	tokenString, err := token.Encode()
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrTokenInternal, err)
	}

	return tokenString, nil
}

func Validate(tokenString string) (*Token, error) {

	tokenID, err := extractPayload(tokenString)
	if err != nil {
		return nil, err
	}

	token, err := findByID(tokenID)
	if err != nil || !token.Enabled {
		return nil, ErrTokenUnauthorized
	}

	if time.Now().After(token.Expiry) {
		return nil, ErrTokenExpired
	}

	return token, nil
}

func extractPayload(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("secret")), nil
	})

	if err != nil || !token.Valid {
		return "", ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", ErrTokenInvalid
	}

	tokenID := claims["tokenID"].(string)

	return tokenID, nil
}
