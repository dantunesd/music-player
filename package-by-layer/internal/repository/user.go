package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type user struct {
	database database
}

func NewUser(database database) *user {
	return &user{
		database: database,
	}
}

func (u *user) Create(user *domain.User) (string, error) {
	return u.database.Create(&user)
}

func (u *user) Get(id string) (*domain.User, error) {
	var user domain.User
	return &user, u.database.Get("_id", id, &user)
}

func (u *user) GetAll() ([]*domain.User, error) {
	var user []*domain.User
	return user, u.database.GetAll(&user)
}
