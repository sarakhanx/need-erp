package purchasing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/departmentsController/purchasing"
)

func PurchasingRouters(app *fiber.App) {
	api := app.Group("/api/purchasing")

	api.Get("/purchase-debug", purchasing.Purchase)
}
