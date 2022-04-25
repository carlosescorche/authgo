package handlers

import (
	"net/http"

	"github.com/carlosescorche/authgo/model/user"
	"github.com/carlosescorche/authgo/utils/api"
)

func HandlerUserGet(w http.ResponseWriter, r *http.Request) {

	user := user.NewUser()

	api.Success(w, user, http.StatusOK)
}
