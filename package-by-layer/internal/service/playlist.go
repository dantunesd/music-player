package service

import "music-player/package-by-layer/internal/domain"

type playlistRepository interface {
	Create(*domain.Playlist) (string, error)
	GetAll() ([]*domain.Playlist, error)
	Get(id string) (*domain.Playlist, error)
}

type playlist struct {
	repository playlistRepository
}

func NewPlaylist(repository playlistRepository) *playlist {
	return &playlist{
		repository: repository,
	}
}

func (s *playlist) Create(userId, name string, songs []string) (*domain.Playlist, error) {
	Playlist := domain.Playlist{
		UserID: userId,
		Name:   name,
		Songs:  songs,
	}

	id, err := s.repository.Create(&Playlist)
	if err != nil {
		return nil, err
	}

	Playlist.ID = id
	return &Playlist, nil
}

func (s *playlist) Get(id string) (*domain.Playlist, error) {
	return s.repository.Get(id)
}

func (s *playlist) GetAll() ([]*domain.Playlist, error) {
	return s.repository.GetAll()
}
