package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type song struct {
	database database
}

func NewSong(database database) *song {
	return &song{
		database: database,
	}
}

func (s *song) Create(song *domain.Song) (string, error) {
	return s.database.Create(&song)
}

func (s *song) Get(id string) (*domain.Song, error) {
	var song domain.Song
	return &song, s.database.Get("_id", id, &song)
}

func (s *song) GetAll() ([]*domain.Song, error) {
	var song []*domain.Song
	return song, s.database.GetAll(&song)
}
