package healthhandlers

import (
	"eazimation-backend/internal/api"
	"net/http"
)

func GetHealthHandler() api.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		return api.WriteJSON(w, 200, &response{Message: "OK"})
	}
}
