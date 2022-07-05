package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type playlist struct {
	database database
}

func NewPlaylist(database database) *playlist {
	return &playlist{
		database: database,
	}
}

func (p *playlist) Create(playlist *domain.Playlist) (string, error) {
	return p.database.Create(&playlist)
}

func (p *playlist) Get(id string) (*domain.Playlist, error) {
	var playlist domain.Playlist
	return &playlist, p.database.Get("_id", id, &playlist)
}

func (p *playlist) GetAll() ([]*domain.Playlist, error) {
	var playlist []*domain.Playlist
	return playlist, p.database.GetAll(&playlist)
}
