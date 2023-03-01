package repository

import (
	"auth-server/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string, password string) (model.User, error)
}

type Repository struct {
	Authorization
}

func (r Repository) CreateUser(user model.User) (int, error) {
	return 0, nil
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthBD(db),
	}
}
