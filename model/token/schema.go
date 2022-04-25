package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  primitive.ObjectID `bson:"userId"`
	Enabled bool               `bson:"enabled"`
	Expiry  time.Time          `bson:"expiry"`
}

func newToken(userID primitive.ObjectID) *Token {
	return &Token{
		ID:      primitive.NewObjectID(),
		UserID:  userID,
		Enabled: true,
		Expiry:  time.Now().Add(time.Hour * 6),
	}
}

func (t *Token) Encode() (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"tokenID": t.ID.Hex(),
		"userID":  t.UserID.Hex(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("secret")))

	fmt.Println(tokenString, err)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
