package productrouters

import (
	"github.com/gofiber/fiber/v2"
	productcontroller "github.com/need/go-backend/controllers/essential-functions/productcontrollers"
)

func ProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create-product", productcontroller.CreateProduct)
	api.Post("/insert-product", productcontroller.ReplenishProduct)
	api.Post("/create-product-procedure", productcontroller.CreateProductWithProcedure)

	api.Delete("/del-product/:product_id", productcontroller.DeleteProduct)

	api.Get("/product/:product_id", productcontroller.GetProductByID)
	api.Get("/products", productcontroller.GetHundredProducts)
	api.Get("/product", productcontroller.GetProductByCategory)
}
