package departmentrouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/moderatorcontroller"
)

func DepartmentRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-department", moderatorcontroller.CreateDepartment)
	api.Get("/get-department/:department", moderatorcontroller.GetDepartment)
	api.Get("/get-departments", moderatorcontroller.GetAllDepartments)
	api.Delete("/del-department/:department", moderatorcontroller.DeleteDepartments)
}
