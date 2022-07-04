package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type Song struct {
	Database Database
}

func (s *Song) Create(song *domain.Song) (string, error) {
	return s.Database.Create(&song)
}

func (s *Song) Get(id string) (*domain.Song, error) {
	var song domain.Song
	return &song, s.Database.Get("_id", id, &song)
}

func (s *Song) GetAll() ([]*domain.Song, error) {
	var song []*domain.Song
	return song, s.Database.GetAll(&song)
}
