package service

import "music-player/package-by-layer/internal/domain"

type UserRepository interface {
	Create(user *domain.User) (string, error)
	GetAll() ([]*domain.User, error)
	Get(id string) (*domain.User, error)
}

type User struct {
	Repository UserRepository
}

func (s *User) Create(name string) (*domain.User, error) {
	user := domain.User{Name: name}

	id, err := s.Repository.Create(&user)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

func (s *User) Get(id string) (*domain.User, error) {
	return s.Repository.Get(id)
}

func (s *User) GetAll() ([]*domain.User, error) {
	return s.Repository.GetAll()
}
