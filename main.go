package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/routes/branchrouters"
	"github.com/need/go-backend/routes/departmentrouters"
	"github.com/need/go-backend/routes/departmentsRouters/purchasing"
	"github.com/need/go-backend/routes/productrouters"
	"github.com/need/go-backend/routes/userrouters"
)

func main() {
	app := fiber.New()
	dbconfig.InitDB()

	log.Println("Server is starting...")
	defer dbconfig.DB.Close()
	//Routers
	userrouters.UserRouter(app)
	departmentrouters.DepartmentRoutes(app)
	branchrouters.BranchRouter(app)
	productrouters.ProductRoutes(app)
	purchasing.PurchasingRouters(app)

	log.Println("Server is running on port 8000...")

	log.Println(app.Listen(":8000"))
}
