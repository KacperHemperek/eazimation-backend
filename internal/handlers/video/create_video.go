package videohandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"net/http"
)

func HandleCreateVideo() api.HandlerFunc {

	//type request struct {
	//	VideoID string `json:"videoId"`
	//}

	return func(w http.ResponseWriter, r *http.Request) error {
		_, err := auth.GetSessionFromRequest(*r)

		if err != nil {
			return err
		}

		return api.WriteJSON(w, http.StatusCreated, map[string]any{
			"message": "video created successfully",
		})
	}
}
