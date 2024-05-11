package services

import "eazimation-backend/internal/database"

type VideoService interface {
	Create(userID, videoID string, data any) (*database.VideoModel, error)
	GetByID(id int) (*database.VideoModel, error)
}

type PGVideoService struct {
	db database.Store
}

func (s *PGVideoService) Create(userID, videoID string, data any) (*database.VideoModel, error) {
	return nil, nil
}

func (s *PGVideoService) GetByID(id int) (*database.VideoModel, error) {
	return nil, nil
}

func NewPGVideoService(db database.Store) *PGVideoService {
	return &PGVideoService{
		db: db,
	}
}
