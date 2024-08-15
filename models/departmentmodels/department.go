package departmentmodels

import (
	"database/sql"
)

type Department struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Position    string        `json:"position"`
	Member      sql.NullInt64 `json:"member"`
	CreatedDate sql.NullTime  `json:"created_date"`
	UpdatedDate sql.NullTime  `json:"updated_date"`
}
