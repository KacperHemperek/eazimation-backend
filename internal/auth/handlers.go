package auth

import (
	"database/sql"
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/services/user"
	"errors"
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
)

var (
	frontendUrl = os.Getenv("FRONTEND_URL")
)

func HandleAuthCallback(sessionStore SessionStore, userService services.UserService) api.HandlerFunc {
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

		sessionID, err := sessionStore.AddSession(Session{
			UserID: user.ID,
			Email:  user.Email,
			Avatar: user.Avatar,
		})

		if err != nil {
			return err
		}

		SetSessionCookie(w, sessionID)

		http.Redirect(w, r, frontendUrl, http.StatusFound)
		return nil
	}
}

func HandleLogout(sessionStore SessionStore) api.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		sessionCookie, err := r.Cookie(SessionCookieName)
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

		RemoveSessionCookie(w)
		return api.WriteJSON(w, http.StatusOK, &response{Message: "User logged out successfully"})
	}
}

func HandleAuth(sessionStore SessionStore, userService services.UserService) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if authUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			user, err := userService.GetByEmail(authUser.Email)

			if err != nil {
				return err
			}

			sessionID, err := sessionStore.AddSession(Session{
				UserID: user.ID,
				Email:  user.Email,
				Avatar: user.Avatar,
			})
			if err != nil {
				return err
			}

			SetSessionCookie(w, sessionID)
			http.Redirect(w, r, frontendUrl, http.StatusFound)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
		return nil
	}
}

func HandleGetUser() api.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		session, err := GetSessionFromRequest(*r)

		if err != nil {
			return NewUnauthorizedApiError(err)
		}

		return api.WriteJSON(w, http.StatusOK, session)
	}
}

func HandleLambdaAuth(sessionStore SessionStore) api.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		sessionID := r.URL.Query().Get("session_id")
		if sessionID == "" {
			return &api.Error{
				Message: "session_id is missing from query params",
				Code:    http.StatusBadRequest,
			}
		}
		session, err := sessionStore.GetSession(sessionID)
		if err != nil {
			return NewUnauthorizedApiError(err)
		}
		return api.WriteJSON(w, http.StatusOK, session)

	}
}

func NewUnauthorizedApiError(err error) *api.Error {
	return &api.Error{
		Message: "Unauthorized",
		Code:    http.StatusUnauthorized,
		Cause:   err,
	}
}
