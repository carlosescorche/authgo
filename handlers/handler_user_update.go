package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/carlosescorche/authgo/model/token"
	"github.com/carlosescorche/authgo/model/user"
	"github.com/carlosescorche/authgo/utils/api"
	e "github.com/carlosescorche/authgo/utils/errors"
	"github.com/carlosescorche/authgo/utils/validator"
)

type HandlerUserUpdatePayload struct {
	FirstName string `json:"firstname" validate:"required,max=100"`
	LastName  string `json:"lastname" validate:"required,max=100"`
	Username  string `json:"username" validate:"required,max=100"`
	Email     string `json:"email" validate:"required,email"`
}

func HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	var payload HandlerUserUpdatePayload

	token, err := token.Validate(api.GetTokenString(r))
	if err != nil {
		api.Error(w, e.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		api.Error(w, e.ErrInvalidPayload, http.StatusBadRequest)
		return
	}

	if errs, ok := validator.ValidateStruct(payload); !ok {
		api.Error(w, e.NewPayloadError(errs), http.StatusBadRequest)
		return
	}

	u, err := user.Get(token.UserID.Hex())

	u.FirstName = payload.FirstName
	u.LastName = payload.LastName
	u.Username = payload.Username
	u.Email = payload.Email

	err = user.Update(u)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserExist):
			err := map[string][]string{"username": {"The username is registered"}}
			api.Error(w, e.NewPayloadError(err), http.StatusBadRequest)
			return
		case errors.Is(err, user.ErrUserEmailExist):
			err := map[string][]string{"email": {"The email is registered"}}
			api.Error(w, e.NewPayloadError(err), http.StatusBadRequest)
			return
		default:
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}
	}

	api.Success(w, nil, http.StatusAccepted)
}
