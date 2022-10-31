package service

import "github.com/Serelllka/Federacion/pkg/repository"

type Authorization interface {
}

type Friends interface {
}

type Service struct {
	Authorization
	Friends
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
