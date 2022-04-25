package errors

import (
	"fmt"
)

var ErrInternal error = NewCustomError("errInternal", "internal error", nil)
var ErrInvalidPayload error = NewCustomError("errInvalidPayload", "the payload is invalid", nil)
var ErrForbidden error = NewCustomError("errForbidden", "Does not have permissions", nil)
var ErrUnauthorized error = NewCustomError("errUnauthorized", "Unauthorized", nil)
var ErrTokenExpired error = NewCustomError("errTokenExpired", "The token is expired", nil)

type CustomError struct {
	Code    string
	Message string
	Extra   interface{}
}

type PayloadError struct {
	Message     string
	Validations interface{}
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Message)
}

func NewCustomError(code string, message string, extra interface{}) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Extra:   extra,
	}
}

func NewPayloadError(errs interface{}) *CustomError {
	return NewCustomError("errInvalidPayload", "invalid payload", errs)
}
