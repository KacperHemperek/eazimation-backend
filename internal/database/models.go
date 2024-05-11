package database

import "time"

type UserModel struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type VideoModel struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	VideoID   string    `json:"videoId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	VideoData any       `json:"videoData"`
}
