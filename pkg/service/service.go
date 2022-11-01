package service

import (
	federacion "github.com/Serelllka/Federacion"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type Authorization interface {
	CreateUser(user federacion.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Friends interface {
}

type Service struct {
	Authorization
	Friends
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
