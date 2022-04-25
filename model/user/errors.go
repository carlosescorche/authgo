package user

import "errors"

var ErrUserNotFound error = errors.New("user not found")
var ErrUserExist error = errors.New("user the username is registered")
var ErrUserEmailExist error = errors.New("user the email is registered")
var ErrUserUnauthorized error = errors.New("user unauthorized")
var ErrUserInternal error = errors.New("user internal error")
var ErrUserInvalidPassword error = errors.New("user invalid password")
var ErrUserToken error = errors.New("user there was an error creating the token")
var ErrUserInvalid error = errors.New("user is invalid")
