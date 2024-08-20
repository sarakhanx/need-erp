package purchasing

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/models/departmentmodels/purchase"
)

func Purchase(c *fiber.Ctx) error {
	conn := dbconfig.DB
	var DocHeaderData purchase.PurchaseOrderDocument

	_, err := conn.Exec(`INSERT INTO Document (doc_status)`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while insert data",
			"data":   err.Error(),
		})
	}
	fmt.Printf("here is model from form %+v\n", DocHeaderData)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created a document successfully",
		"data":    DocHeaderData,
	})
}
