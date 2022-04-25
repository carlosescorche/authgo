package user

import (
	"fmt"
	"time"

	"github.com/carlosescorche/authgo/utils/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstname" validate:"required,min=1,max=100"`
	LastName  string             `bson:"lastname" validate:"required,min=1,max=100"`
	Username  string             `bson:"username" validate:"required,min=5,max=100"`
	Email     string             `bson:"email" validate:"required,email"`
	Password  string             `bson:"password" validate:"required"`
	Roles     []string           `bson:"roles"`
	Enabled   bool               `bson:"enabled"`
	Created   time.Time          `bson:"created"`
	Updated   time.Time          `bson:"updated"`
}

func NewUser() *User {
	return &User{
		ID:      primitive.NewObjectID(),
		Enabled: true,
		Roles:   []string{"user"},
		Created: time.Now(),
		Updated: time.Now(),
	}
}

func (u *User) SetPassword(pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateSchema() error {
	if errs, ok := validator.ValidateStruct(u); !ok {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
