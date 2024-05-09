package authhandlers

import (
	"database/sql"
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	services "eazimation-backend/internal/services/user"
	"errors"
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
)

var (
	frontendUrl = os.Getenv("FRONTEND_URL")
)

func HandleAuthCallback(sessionStore auth.SessionStore, userService services.UserService) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		authUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("%s/auth/failed", frontendUrl), http.StatusFound)
			return nil
		}

		user, err := userService.GetByEmail(authUser.Email)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		if err != nil && errors.Is(err, sql.ErrNoRows) {
			createdUser, createUserError := userService.Create(authUser.Email, authUser.AvatarURL)

			if createUserError != nil {
				return createUserError
			}

			user = createdUser
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
		return nil
	}
}
