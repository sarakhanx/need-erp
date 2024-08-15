package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var SecretKey []byte

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenToString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenToString, nil
}
