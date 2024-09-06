package userrouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/essential-functions/usercontroller"
	"github.com/need/go-backend/middlewares/auth"
)

func UserRouter(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/debug-user", auth.IsAdmin, usercontroller.DebugUser)
	api.Post("/signup", usercontroller.SignupUser)
	api.Post("/signin", usercontroller.SigninUser)
	api.Post("/signout", usercontroller.SignOut)
	api.Put("/reset-password", usercontroller.ResetPassUser)
	api.Delete("/delete-user", usercontroller.DeleteUser)
}
