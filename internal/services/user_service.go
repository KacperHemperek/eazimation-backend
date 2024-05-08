package services

import "eazimation-backend/internal/database"

type UserService interface {
	Create(name, email string) (*database.UserModel, error)
	GetByID(id int) (*database.UserModel, error)
	GetByEmail(email string) (*database.UserModel, error)
}

type PGUserService struct {
	db database.Service
}

func (s *PGUserService) Create(name, email string) (*database.UserModel, error) {
	return nil, nil
}

func (s *PGUserService) GetByID(id int) (*database.UserModel, error) {
	return nil, nil
}

func (s *PGUserService) GetByEmail(email string) (*database.UserModel, error) {
	return nil, nil
}

func NewPGUserService(db database.Service) *PGUserService {
	return &PGUserService{
		db: db,
	}
}
