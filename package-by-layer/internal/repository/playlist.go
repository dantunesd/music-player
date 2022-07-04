package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type Playlist struct {
	Database Database
}

func (p *Playlist) Create(playlist *domain.Playlist) (string, error) {
	return p.Database.Create(&playlist)
}

func (p *Playlist) Get(id string) (*domain.Playlist, error) {
	var playlist domain.Playlist
	return &playlist, p.Database.Get("_id", id, &playlist)
}

func (p *Playlist) GetAll() ([]*domain.Playlist, error) {
	var playlist []*domain.Playlist
	return playlist, p.Database.GetAll(&playlist)
}
