package service

import "music-player/package-by-layer/internal/domain"

type songRepository interface {
	Create(*domain.Song) (string, error)
	GetAll() ([]*domain.Song, error)
	Get(id string) (*domain.Song, error)
}

type song struct {
	repository songRepository
}

func NewSong(repository songRepository) *song {
	return &song{
		repository: repository,
	}
}

func (s *song) Create(name, artistName, albumName string, number int) (*domain.Song, error) {
	song := domain.Song{
		Name:       name,
		ArtistName: artistName,
		AlbumName:  albumName,
		Number:     number,
	}

	id, err := s.repository.Create(&song)
	if err != nil {
		return nil, err
	}

	song.ID = id
	return &song, nil
}

func (s *song) Get(id string) (*domain.Song, error) {
	return s.repository.Get(id)
}

func (s *song) GetAll() ([]*domain.Song, error) {
	return s.repository.GetAll()
}
