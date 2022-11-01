package federacion

type Article struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	CurrentCost string `json:"current_cost" binding:"required"`
	Description string `json:"description"`
}
