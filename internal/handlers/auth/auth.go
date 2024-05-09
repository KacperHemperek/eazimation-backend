package authhandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	services "eazimation-backend/internal/services/user"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func HandleAuth(sessionStore auth.SessionStore, userService services.UserService) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if authUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			user, err := userService.GetByEmail(authUser.Email)

			if err != nil {
				return err
			}

			sessionID, err := sessionStore.AddSession(auth.Session{
				UserID: user.ID,
				Email:  user.Email,
				Avatar: user.Avatar,
			})
			if err != nil {
				return err
			}

			auth.SetSessionCookie(w, sessionID)
			http.Redirect(w, r, frontendUrl, http.StatusFound)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
		return nil
	}
}
