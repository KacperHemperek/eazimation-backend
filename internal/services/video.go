package services

import (
	"eazimation-backend/internal/database"
)

type VideoService interface {
	Create(userID int, videoID string, data any) (*database.VideoModel, error)
	GetByID(id int) (*database.VideoModel, error)
	GetUserVideos(userID int) ([]*database.VideoModel, error)
}

type PGVideoService struct {
	db database.Store
}

func (s *PGVideoService) Create(userID int, videoID string, data any) (*database.VideoModel, error) {
	row := s.db.Client.QueryRow(
		"insert into rendered_videos (user_id, video_id, video_data) values($1, $2, $3) returning id, user_id, video_id, video_data, created_at, updated_at;",
		userID, videoID, data,
	)

	return scanVideo(row)
}

func (s *PGVideoService) GetUserVideos(userID int) ([]*database.VideoModel, error) {
	rows, err := s.db.Client.Query(
		"select id, user_id, video_id, video_data, created_at, updated_at from rendered_videos where user_id = $1;",
		userID,
	)

	videos := make([]*database.VideoModel, 0)

	if err != nil {
		return videos, err
	}

	for rows.Next() {
		video, err := scanVideo(rows)

		if err != nil {
			return make([]*database.VideoModel, 0), err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func (s *PGVideoService) GetByID(videoID int) (*database.VideoModel, error) {
	row := s.db.Client.QueryRow("select id, user_id, video_id, video_data, created_at, updated_at from rendered_videos where id = $1", videoID)
	return scanVideo(row)
}

// NewPGVideoService is a constructor for VideoService implementation
// using postgres as a persistent storage
func NewPGVideoService(db database.Store) *PGVideoService {
	return &PGVideoService{
		db: db,
	}
}

// scanVideo scans sql query result in order:
// id, user_id, video_id, video_data, created_at, updated_at
func scanVideo(row SqlScanner) (*database.VideoModel, error) {
	video := &database.VideoModel{}
	err := row.Scan(&video.ID, &video.UserID, &video.VideoID, &video.VideoData, &video.CreatedAt, &video.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return video, nil
}
