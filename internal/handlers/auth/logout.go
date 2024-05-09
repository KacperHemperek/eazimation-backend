package authhandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"errors"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func HandleLogout(sessionStore auth.SessionStore) api.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		sessionCookie, err := r.Cookie(auth.SessionCookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return &api.Error{
					Message: "User is not logged in",
					Code:    http.StatusUnauthorized,
					Cause:   err,
				}
			}
			return err
		}
		err = gothic.Logout(w, r)
		if err != nil {
			return err
		}

		err = sessionStore.RemoveSession(sessionCookie.Value)

		if err != nil {
			return err
		}

		auth.RemoveSessionCookie(w)
		return api.WriteJSON(w, http.StatusOK, &response{Message: "User logged out successfully"})
	}
}
