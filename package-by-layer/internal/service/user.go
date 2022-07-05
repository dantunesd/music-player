package service

import "music-player/package-by-layer/internal/domain"

type userRepository interface {
	Create(user *domain.User) (string, error)
	GetAll() ([]*domain.User, error)
	Get(id string) (*domain.User, error)
}

type user struct {
	repository userRepository
}

func NewUser(repository userRepository) *user {
	return &user{
		repository: repository,
	}
}

func (s *user) Create(name string) (*domain.User, error) {
	user := domain.User{Name: name}

	id, err := s.repository.Create(&user)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

func (s *user) Get(id string) (*domain.User, error) {
	return s.repository.Get(id)
}

func (s *user) GetAll() ([]*domain.User, error) {
	return s.repository.GetAll()
}
