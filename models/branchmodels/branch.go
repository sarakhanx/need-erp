package branchmodels

import "database/sql"

type Branch struct {
	ID            int          `json:"branch_id"`
	BranchName    string       `json:"branch_name"`
	BranchAddress string       `json:"branch_address"`
	Updated_at    sql.NullTime `json:"updated_at"`
}
type BracnhProductQty struct {
	ID         int `json:"branch_id"`
	Product_ID int `json:"product_id"`
	QTY        int `json:"qty"`
}
