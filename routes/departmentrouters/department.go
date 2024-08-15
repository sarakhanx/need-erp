package departmentrouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/departmentcontroller"
)

func DepartmentRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-department", departmentcontroller.CreateDepartment)
	api.Get("/get-department/:department", departmentcontroller.GetDepartment)
	api.Get("/get-departments", departmentcontroller.GetAllDepartments)
	api.Delete("/del-department/:department", departmentcontroller.DeleteDepartments)
}
