package repository

import (
	"music-player/package-by-layer/internal/domain"
)

type User struct {
	Database Database
}

func (u *User) Create(user *domain.User) (string, error) {
	return u.Database.Create(&user)
}

func (u *User) Get(id string) (*domain.User, error) {
	var user domain.User
	return &user, u.Database.Get("_id", id, &user)
}

func (u *User) GetAll() ([]*domain.User, error) {
	var user []*domain.User
	return user, u.Database.GetAll(&user)
}
