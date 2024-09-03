package productcontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/models/productmodels"
	"github.com/need/go-backend/queries/productquery"
)

func GetProductByCategory(c *fiber.Ctx) error {

	// queryParams := make(map[string]string)
	// queryParams["categ"] = c.Query("categ")
	// queryParams["limit"] = c.Query("limit")
	// queryParams["offset"] = c.Query("offset")
	var (
		params     = c.Queries()
		categ      = params["categ"]
		limitQuery = params["limit"]
		pageQuery  = params["page"]
	)

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid params please review your url",
			"data":    err.Error(),
		})
	}
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid params please review your url",
			"data":    err.Error(),
		})
	}
	offset := (page - 1) * limit

	conn := dbconfig.DB

	var TotalRows int
	err = conn.QueryRow(productquery.CountProductsByCategory, categ).Scan(&TotalRows)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Can not count all page, Database server error",
			"data":    err.Error(),
		})
	}
	TotalPage := (TotalRows + limit - 1) / limit
	if page > TotalPage {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Page not found",
		})
	}

	rows, err := conn.Query(productquery.GetProductsByCategory, categ, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while executing query",
			"data":   err.Error(),
		})
	}
	defer rows.Close()

	product := []productmodels.Product{}
	for rows.Next() {
		var p productmodels.Product
		if err := rows.Scan(&p.ProductName, &p.Cost, &p.Price, &p.Category, &p.ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  "Error while scanning product rows",
				"data":   err.Error(),
			})
		}
		product = append(product, p)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Error while iterating product rows",
			"data":   err.Error(),
		})
	}

	if len(product) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No products found for the given category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Products query by category successfully",
		"products":   product,
		"totalPages": TotalPage,
		"totalRows":  TotalRows,
		"page":       page,
	})
}
