package service

import (
	"auth-server/api/repository"
	"auth-server/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
