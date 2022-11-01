package repository

import (
	"fmt"

	federacion "github.com/Serelllka/Federacion"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user federacion.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := ap.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (federacion.User, error) {
	var user federacion.User
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := ap.db.Get(&user, query, username, password)

	return user, err
}
