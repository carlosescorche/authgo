package token

import "errors"

var ErrTokenInsert error = errors.New("token could not be inserted")
var ErrTokenInternal error = errors.New("token internal error")
var ErrTokenUnauthorized error = errors.New("token unauthorize")
var ErrTokenExpired error = errors.New("token expired")
var ErrTokenInvalid error = errors.New("token invalid")
