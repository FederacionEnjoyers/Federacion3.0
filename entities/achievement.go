package entities

import "database/sql"

type Achievement struct {
	Id          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name" binding:"required"`
	Description string         `json:"description" db:"description" binding:"required"`
	Image       sql.NullString `json:"image" db:"image"`
}
