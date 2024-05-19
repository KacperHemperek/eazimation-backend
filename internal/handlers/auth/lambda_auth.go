package authhandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"net/http"
)

func HandleLambdaAuth() api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		session, err := auth.GetSessionFromRequest(*r)
		if err != nil {
			return api.NewUnauthorizedApiError(err)
		}
		return api.WriteJSON(w, http.StatusOK, session)

	}
}
