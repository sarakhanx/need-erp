package uservalidator

import (
	"database/sql"
	"log"
	"regexp"

	_ "github.com/lib/pq"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/queries/userquery"
)

func IsValidEmail(email string) bool {
	//NOTE -  Simple regex for email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
func IsExistUser(email string) (bool, error) {

	conn := dbconfig.DB
	log.Println("Checking existence for email:", email)

	query := userquery.CheckExistingUser

	var exist bool
	err := conn.QueryRow(query, email).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	return exist, err
}
