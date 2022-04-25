package user

import (
	"fmt"
	"strings"

	"github.com/carlosescorche/authgo/model/token"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(login string, password string) (string, error) {
	user, err := findByLogin(login)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}

	if !user.Enabled {
		return "", fmt.Errorf("%w: %v", ErrUserUnauthorized, err)
	}

	if err = user.ValidatePassword(password); err != nil {
		return "", fmt.Errorf("%w: %v", ErrUserInvalidPassword, err)
	}

	token, err := token.Create(user.ID)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrUserToken, err)
	}

	return token, nil
}

func Insert(user *User) error {
	err := user.ValidateSchema()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUserInvalid, err)
	}

	_, err = insert(user)
	if err != nil {
		switch {
		case mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "username_1"):
			return fmt.Errorf("%w: %v", ErrUserExist, err)
		case mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "email_1"):
			return fmt.Errorf("%w: %v", ErrUserEmailExist, err)
		default:
			return fmt.Errorf("%w: %v", ErrUserInternal, err)
		}
	}

	return nil
}
