package service

import "music-player/package-by-layer/internal/domain"

type SongRepository interface {
	Create(*domain.Song) (string, error)
	GetAll() ([]*domain.Song, error)
	Get(id string) (*domain.Song, error)
}

type Song struct {
	Repository SongRepository
}

func (s *Song) Create(name, artistName, albumName string, number int) (*domain.Song, error) {
	song := domain.Song{
		Name:       name,
		ArtistName: artistName,
		AlbumName:  albumName,
		Number:     number,
	}

	id, err := s.Repository.Create(&song)
	if err != nil {
		return nil, err
	}

	song.ID = id
	return &song, nil
}

func (s *Song) Get(id string) (*domain.Song, error) {
	return s.Repository.Get(id)
}

func (s *Song) GetAll() ([]*domain.Song, error) {
	return s.Repository.GetAll()
}
