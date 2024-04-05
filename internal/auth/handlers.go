package auth

import (
	"eazimation-backend/internal/api"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func HandleAuthCallback(sessionStore SessionStore) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Redirect(w, r, "http://localhost:5173/auth/failed", http.StatusFound)
			return nil
		}

		newSess := NewSessionUser(user.UserID, user.Email, user.AvatarURL)
		sessionID := sessionStore.AddSession(newSess)

		SetSessionCookie(w, sessionID)

		http.Redirect(w, r, "http://localhost:5173/auth/success", http.StatusFound)
		return nil
	}
}

func HandleLogout() api.HandlerFunc {
	// TODO: remove session from sessionStore when user logs out
	return func(w http.ResponseWriter, r *http.Request) error {
		err := gothic.Logout(w, r)
		if err != nil {
			return err
		}
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuth(sessionStore SessionStore) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if user, err := gothic.CompleteUserAuth(w, r); err == nil {
			session := NewSessionUser(user.UserID, user.Email, user.AvatarURL)
			sessionID := sessionStore.AddSession(session)

			SetSessionCookie(w, sessionID)
			http.Redirect(w, r, "http://localhost:5173/auth/success", http.StatusFound)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
		return nil
	}
}

func HandleGetUser(sessionStore SessionStore) api.HandlerFunc {
	type response struct {
		User *SessionUser `json:"user"`
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

		return api.WriteJSON(w, http.StatusOK, &response{User: session})

	}
}

func HandleLambdaAuth(sessionStore SessionStore) api.HandlerFunc {
	type response struct {
		User *SessionUser `json:"user"`
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
		return api.WriteJSON(w, http.StatusOK, &response{User: session})

	}
}

func NewUnauthorizedApiError(err error) *api.Error {
	return &api.Error{
		Message: "Unautorized",
		Code:    http.StatusUnauthorized,
		Cause:   err,
	}
}
