package branchesvalidator

import (
	"database/sql"
	"log"

	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/queries/branchesquery"
)

func CheckBranchExisted(branchName string) (bool, error) {
	conn := dbconfig.DB
	log.Println("Checking existence for branch:", branchName)

	var exist bool
	err := conn.QueryRow(branchesquery.GetBranchByName, branchName).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	return exist, err
}
