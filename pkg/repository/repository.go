package repository

type Authorization interface {
}

type Friends interface {
}

type Repository struct {
	Authorization
	Friends
}

func NewRepository() *Repository {
	return &Repository{}
}
