package usermodels

import "database/sql"

type User struct {
	ID        int          `json:"user_id"`
	Name      string       `json:"name"`
	LastName  string       `json:"lastname"`
	Password  string       `json:"password"`
	Mobile    string       `json:"mobile"`
	LastLogin sql.NullTime `json:"last_login"`
	Email     string       `json:"email"`
	Role      string       `json:"role"`
	Position  string       `json:"position"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
