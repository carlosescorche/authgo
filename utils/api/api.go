package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosescorche/authgo/utils/errors"
)

type HTTPResponseEnvelope struct {
	HTTPStatus int         `json:"httpStatus"`
	Data       interface{} `json:"data"`
}

type HTTPResponseError struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Extra   interface{} `json:"extra"`
}

type ErrorInfo struct {
	Code    string
	Message interface{}
	Extra   interface{}
}

func GetTokenString(r *http.Request) string {
	var tokenString string

	authValue := r.Header.Get("Authorization")
	split := strings.Split(authValue, "Bearer ")

	if len(split) > 1 {
		tokenString = split[1]
	}

	return tokenString
}

func Error(w http.ResponseWriter, err error, code int) {
	envelope := HTTPResponseError{}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)

	if e, ok := err.(*errors.CustomError); ok {
		envelope = HTTPResponseError{
			Code:    e.Code,
			Message: e.Message,
			Extra:   e.Extra,
			Status:  code,
		}
	}

	json.NewEncoder(w).Encode(envelope)
	return
}

func Success(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)

	envelope := HTTPResponseEnvelope{
		HTTPStatus: code,
		Data:       response,
	}

	json.NewEncoder(w).Encode(envelope)
	return
}
