package usercontroller

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/need/go-backend/config/db-config"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/middlewares/uservalidator"
	"github.com/need/go-backend/models/usermodels"
	"github.com/need/go-backend/queries/userquery"
	"github.com/need/go-backend/utils/bcrypt"
	"github.com/need/go-backend/utils/jwt"
)

func DebugUser(c *fiber.Ctx) error {

	return c.SendString("Hello World")
}

func SignupUser(c *fiber.Ctx) error {
	conn := dbconfig.DB
	var signupData usermodels.User

	if err := c.BodyParser(&signupData); err != nil {
		return c.Status(500).JSON(fiber.Map{"error occurred when parsing body": err})
	}
	// EXPLAIN -  Validate data
	if strings.TrimSpace(signupData.Name) == "" ||
		strings.TrimSpace(signupData.LastName) == "" ||
		strings.TrimSpace(signupData.Password) == "" ||
		strings.TrimSpace(signupData.Mobile) == "" ||
		strings.TrimSpace(signupData.Email) == "" ||
		strings.TrimSpace(signupData.Role) == "" ||
		strings.TrimSpace(signupData.Position) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No Blank Spaces Allowed"})
	}
	if !uservalidator.IsValidEmail(signupData.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Email"})
	}
	if len(signupData.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password too short"})
	}

	exist, err := uservalidator.IsExistUser(signupData.Email)
	if err != nil {
		log.Println("Error checking user existence:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User already exist" + err.Error()})
	}
	if exist {
		log.Println("User already exists with email:", signupData.Email)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exist"})
	}
	log.Println("User registration successful for email:", signupData.Email)
	//NOTE : เริ่มการบีนทึก
	hashedPassword, err := bcrypt.HashedPassword(signupData.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error hashing password" + err.Error()})
	}

	query := userquery.InsertNewUser
	err = conn.QueryRow(query, signupData.Name, signupData.LastName, signupData.Mobile, signupData.Email, hashedPassword, signupData.Role, signupData.Position).Scan(&signupData.ID)
	if err != nil {
		log.Println("Error inserting new user to database :", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting new user" + err.Error()})
	}
	return c.JSON(signupData)
}

func SigninUser(c *fiber.Ctx) error {
	conn := dbconfig.DB

	var signinData usermodels.User
	if err := c.BodyParser(&signinData); err != nil {
		log.Println("Can not read data in body ", err)
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	var UserSession usermodels.User
	query := userquery.SigninUser
	if err := conn.QueryRow(query, signinData.Email).Scan(&UserSession.ID, &UserSession.Name, &UserSession.LastName, &UserSession.Mobile, &UserSession.Email,
		&UserSession.Password, &UserSession.Role, &UserSession.Position, &UserSession.LastLogin); err != nil {
		log.Println("Not Found user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Found user", "message": err.Error()})
	}

	if err := bcrypt.ComparePasswords(UserSession.Password, signinData.Password); err != nil {
		log.Println("Incorrect Password")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect Password", "message": err.Error()})
	}

	//NOTE : แจก Token จาก JWT
	token, err := jwt.GenerateToken(signinData.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating token", "message": err.Error()})
	}
	LastLogin := userquery.UpdateLastLogin
	if _, err := conn.Exec(LastLogin, time.Now(), UserSession.ID); err != nil {
		log.Println("Error updating last login:", err)
	}

	UserSession.Password = ""
	return c.JSON(fiber.Map{"token": token, "user": UserSession})
}

func SignOut(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return c.JSON(fiber.Map{"Status": "Sign Out Successful"})
}

func ResetPassUser(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
