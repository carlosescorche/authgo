package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/carlosescorche/authgo/model/user"
	"github.com/carlosescorche/authgo/utils/api"
	e "github.com/carlosescorche/authgo/utils/errors"
	"github.com/carlosescorche/authgo/utils/validator"
)

type HandlerUserLoginPayload struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func HandlerUserLogin(w http.ResponseWriter, r *http.Request) {

	var payload HandlerUserLoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		api.Error(w, e.ErrInvalidPayload, http.StatusBadRequest)
		return
	}

	if errs, ok := validator.ValidateStruct(payload); !ok {
		api.Error(w, e.NewPayloadError(errs), http.StatusBadRequest)
		return
	}

	token, err := user.Login(payload.Login, payload.Password)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserNotFound):
			err := map[string][]string{"Login": {"The username or email is invalid"}}
			api.Error(w, e.NewPayloadError(err), http.StatusBadRequest)
			return
		case errors.Is(err, user.ErrUserInvalidPassword):
			err := map[string][]string{"Password": {"The password is invalid"}}
			api.Error(w, e.NewPayloadError(err), http.StatusBadRequest)
			return
		default:
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}
	}

	api.Success(w, token, http.StatusCreated)
}
