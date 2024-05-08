package services

import "eazimation-backend/internal/database"

type UserService interface {
	Create(name, email string) (*database.UserModel, error)
	GetByID(id int) (*database.UserModel, error)
	GetByEmail(email string) (*database.UserModel, error)
}

type PGUserService struct {
	store database.Store
}

func (s *PGUserService) Create(email, avatar string) (*database.UserModel, error) {
	row := s.store.Client.QueryRow("insert into users (email, avatar) values($1, $2) returning id, email, avatar", email, avatar)
	return scanUser(*row)
}

func (s *PGUserService) GetByID(id int) (*database.UserModel, error) {
	return nil, nil
}

func (s *PGUserService) GetByEmail(email string) (*database.UserModel, error) {
	row := s.store.Client.QueryRow("select id, email, avatar from users where email = $1", email)
	return scanUser(*row)
}

func NewPGUserService(db database.Store) *PGUserService {
	return &PGUserService{
		store: db,
	}
}
