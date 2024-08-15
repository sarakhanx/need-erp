package dbconfig

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	//loading .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error while loading `.env` file : ", err)
	}

	//connecting to database
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("DB_URI is not set please review your .env file")
	}

	var err error
	DB, err = sql.Open("postgres", dbURI+"?sslmode=disable")
	if err != nil {
		log.Fatal("Errror while connecting to database : ", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Error while pinging database : ", err)
	}
	log.Println("Database connection established successfully")

	//สร้าง Log file
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	//Initialize Entity
	debugDb()
	initializeEntity()
}
