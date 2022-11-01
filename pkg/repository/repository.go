package repository

import (
	federacion "github.com/Serelllka/Federacion"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user federacion.User) (int, error)
	GetUser(username, password string) (federacion.User, error)
}

type Articles interface {
	CreateArticle(user federacion.User) (int, error)
}

type Repository struct {
	Authorization
	Articles
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
