package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/carlosescorche/authgo/model/token"
	"github.com/carlosescorche/authgo/model/user"
	"github.com/carlosescorche/authgo/utils/api"
	e "github.com/carlosescorche/authgo/utils/errors"
)

type Output struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Enabled   bool      `json:"enabled"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

func HandlerUserGet(w http.ResponseWriter, r *http.Request) {

	token, err := token.Validate(api.GetTokenString(r))
	if err != nil {
		api.Error(w, e.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	u, err := user.Get(token.UserID.Hex())
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserId):
			api.Error(w, e.ErrUnauthorized, http.StatusInternalServerError)
			return
		default:
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}
	}

	output := Output{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Enabled:   u.Enabled,
		Created:   u.Created,
		Updated:   u.Updated,
	}

	api.Success(w, output, http.StatusOK)
}
