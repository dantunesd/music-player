package service

import "music-player/package-by-layer/internal/domain"

type PlaylistRepository interface {
	Create(*domain.Playlist) (string, error)
	GetAll() ([]*domain.Playlist, error)
	Get(id string) (*domain.Playlist, error)
}

type Playlist struct {
	Repository PlaylistRepository
}

func (s *Playlist) Create(userId, name string, songs []string) (*domain.Playlist, error) {
	Playlist := domain.Playlist{
		UserID: userId,
		Name:   name,
		Songs:  songs,
	}

	id, err := s.Repository.Create(&Playlist)
	if err != nil {
		return nil, err
	}

	Playlist.ID = id
	return &Playlist, nil
}

func (s *Playlist) Get(id string) (*domain.Playlist, error) {
	return s.Repository.Get(id)
}

func (s *Playlist) GetAll() ([]*domain.Playlist, error) {
	return s.Repository.GetAll()
}
