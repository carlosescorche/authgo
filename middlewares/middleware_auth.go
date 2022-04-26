package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/carlosescorche/authgo/model/token"
	"github.com/carlosescorche/authgo/utils/api"
	e "github.com/carlosescorche/authgo/utils/errors"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string

		authValue := r.Header.Get("Authorization")
		split := strings.Split(authValue, "Bearer ")

		if len(split) > 1 {
			tokenString = split[1]
		}

		_, err := token.Validate(tokenString)
		if err != nil {
			switch {
			case errors.Is(err, token.ErrTokenUnauthorized):
				api.Error(w, e.ErrForbidden, http.StatusForbidden)
				return
			case errors.Is(err, token.ErrTokenExpired):
				api.Error(w, e.ErrTokenExpired, http.StatusForbidden)
				return
			default:
				api.Error(w, e.ErrUnauthorized, http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
