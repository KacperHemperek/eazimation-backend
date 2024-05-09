package authhandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"net/http"
)

func HandleLambdaAuth(sessionStore auth.SessionStore) api.HandlerFunc {
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
			return auth.NewUnauthorizedApiError(err)
		}
		return api.WriteJSON(w, http.StatusOK, session)

	}
}
