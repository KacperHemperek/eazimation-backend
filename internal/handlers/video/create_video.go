package videohandlers

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/services"
	"net/http"
)

func HandleCreateVideo(videoService services.VideoService) api.HandlerFunc {

	type request struct {
		VideoData map[string]any `json:"videoData"`
		VideoID   string         `json:"videoId"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		session, err := auth.GetSessionFromRequest(*r)

		if err != nil {
			return err
		}

		return api.WriteJSON(w, http.StatusCreated, map[string]any{
			"message": "video created successfully",
		})
	}
}
