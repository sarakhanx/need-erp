package purchasing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/departmentsController/purchase-services/purchasing_order"
)

func PurchasingRouters(app *fiber.App) {
	api := app.Group("/api/purchasing")

	api.Get("/purchase-debug", purchasing_order.Purchase)
	api.Get("/getpo/:id", purchasing_order.GetPurchasingOrderByID)
	api.Post("/createpurchaseorder", purchasing_order.CreatePurchaseOrder)
	api.Get("/purchaseorder", purchasing_order.GetDocumentPagination)
	api.Delete("/purchaseorder", purchasing_order.DeletePoDoc)
}
