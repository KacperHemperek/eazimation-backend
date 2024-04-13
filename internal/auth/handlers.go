package auth

import (
	"eazimation-backend/internal/api"
	"errors"
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
)

var (
	frontendUrl = os.Getenv("FRONTEND_URL")
)

func HandleAuthCallback(sessionStore SessionStore) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("%s/auth/failed", frontendUrl), http.StatusFound)
			return nil
		}
		sessionID, err := sessionStore.AddSession(user.UserID)
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

func HandleAuth(sessionStore SessionStore) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if user, err := gothic.CompleteUserAuth(w, r); err == nil {
			sessionID, err := sessionStore.AddSession(user.UserID)
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

func HandleGetUser(sessionStore SessionStore) api.HandlerFunc {
	type response struct {
		UserID string `json:"userId"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		sessionCookie, err := r.Cookie(SessionCookieName)
		if err != nil {
			return NewUnauthorizedApiError(err)
		}
		session, err := sessionStore.GetSession(sessionCookie.Value)

		if err != nil {
			return NewUnauthorizedApiError(err)
		}

		return api.WriteJSON(w, http.StatusOK, &response{UserID: session})

	}
}

func HandleLambdaAuth(sessionStore SessionStore) api.HandlerFunc {
	type response struct {
		UserID string `json:"userId"`
	}

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
		return api.WriteJSON(w, http.StatusOK, &response{UserID: session})

	}
}

func NewUnauthorizedApiError(err error) *api.Error {
	return &api.Error{
		Message: "Unautorized",
		Code:    http.StatusUnauthorized,
		Cause:   err,
	}
}
