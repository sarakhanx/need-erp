package purchasing

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/middlewares/departmentvalidator/purchasevalidation"
	"github.com/need/go-backend/models/departmentmodels/purchase"
	"github.com/need/go-backend/models/productmodels"

	purchasequery "github.com/need/go-backend/queries/documents/purchase"
)

func Purchase(c *fiber.Ctx) error {
	conn := dbconfig.DB
	var DocHeaderData purchase.PurchaseOrderDocument

	_, err := conn.Exec(`INSERT INTO Document (doc_status) VALUES ('Draft')`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while insert data",
			"data":   err.Error(),
		})
	}
	// fmt.Printf("here is model from form %+v\n", DocHeaderData)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created a document successfully",
		"data":    DocHeaderData,
	})
}

func GetPurchasingOrderByID(c *fiber.Ctx) error {
	conn := dbconfig.DB
	Params := c.Params("id")
	var result purchase.PurchaseOrderDocument
	rows, err := conn.Query(purchasequery.GetDocumentByID, Params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"data":    err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var product productmodels.Product
		err := rows.Scan(
			&result.DocPrefix,
			&result.DocId,
			&result.BranchID,
			&result.UserID,
			&result.VendorData,
			&result.BranchName,
			&result.BranchAddress,
			&result.Username,
			&result.DocNote,
			&result.CreatedAt,
			&result.ExVat,
			&result.InVat,
			&result.Vat,
			&product.ProductName,
			&product.Cost,
			&product.Price,
			&product.ID,
			&product.SaleOrderQty,
			&product.SaleOrderPrice,
			&product.SaleOrderPriceTotal,
			&product.SaleOrderDiscount,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Error scanning row",
				"data":    err.Error(),
			})
		}
		result.Products = append(result.Products, product)
	}

	if len(result.Products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Purchasing order not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get purchasing order by ID successfully",
		"data":    result,
	})
}

func CreatePurchaseOrder(c *fiber.Ctx) error {
	conn := dbconfig.DB
	FormData := purchase.PurchaseOrderDocument{}
	err := c.BodyParser(&FormData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while parsing form data from body",
			"data":    err.Error(),
		})
	}

	if err := purchasevalidation.ValidatePurchaseOrderInput(FormData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation error",
			"data":    err.Error(),
		})
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Println("Error while starting a transection : ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while insert data on Document table",
			"data":   err.Error(),
		})
	}

	result, err := tx.Exec(purchasequery.CreatePODocument,
		&FormData.DocId,
		&FormData.DocStatus,
		&FormData.ExVat,
		&FormData.Vat,
		&FormData.InVat,
		&FormData.DocDiscount,
		&FormData.DocNote,
		&FormData.BranchID,
		&FormData.UserID,
		&FormData.DepartmentID,
		&FormData.VendorData,
		&FormData.DocPrefixID,
	)
	if err != nil {
		tx.Rollback()
		log.Println("Error while create document:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while insert data",
			"data":   err.Error(),
		})
	}

	Products := FormData.Products
	if len(Products) > 0 {
		for _, product := range Products {
			log.Printf("Inserting product: %+v", product)
			_, err = tx.Exec(purchasequery.InsertProdToPO,
				FormData.DocId,
				product.ProductName,
				product.Cost,
				product.Price,
				product.Category,
				product.BranchID,
				product.Branch_QTY,
				product.Action,
				product.UserID,
				product.SaleOrderDiscount,
				product.SaleOrderPrice,
				product.SaleOrderPriceTotal,
			)
			if err != nil {
				log.Printf("Error inserting product: %+v\nError: %v", product, err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status": "error",
					"error":  "Internal server error while insert data on Product table",
					"data":   err.Error(),
				})
			}
		}
	}
	DocumentLog := FormData.DocumentLog
	_, err = tx.Exec(purchasequery.InsertDocLog, DocumentLog.DocAction, FormData.DocId, FormData.UserID, DocumentLog.DocQty, FormData.DepartmentID)
	if err != nil {
		tx.Rollback()
		log.Println("Error while insert document log: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Internal server error while insert data into document log",
			"data":   err.Error(),
		})
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while commit data. See more detail in logs",
			"data":    err.Error(),
		})
	}

	rowsAffected, _ := result.RowsAffected()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       "success",
		"message":      "Created a document successfully",
		"rowsAffected": rowsAffected,
	})
}

func GetDocumentPagination(c *fiber.Ctx) error {
	conn := dbconfig.DB
	Page := c.Query("page")
	page, err := strconv.Atoi(Page)
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid params please review your url",
			"data":    err.Error(),
		})
	}

	limit := 100
	offset := (page - 1) * limit

	var TotalRows int
	err = conn.QueryRow(`select count(*) from documentheader;`).Scan(&TotalRows)
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

	rows, err := conn.Query(purchasequery.GetDocumentPagination, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while excuting query",
			"data":    err.Error(),
		})
	}
	defer rows.Close()

	DocumentData := []purchase.PurchaseOrderDocument{}
	for rows.Next() {
		var po purchase.PurchaseOrderDocument
		var product productmodels.Product

		if err := rows.Scan(
			&po.DocPrefix,
			&po.DocId,
			&po.BranchID,
			&po.UserID,
			&po.VendorData,
			&po.BranchName,
			&po.BranchAddress,
			&po.Username,
			&po.DocNote,
			&po.CreatedAt,
			&po.ExVat,
			&po.InVat,
			&po.Vat,
			&product.ProductName,
			&product.Cost,
			&product.Price,
			&product.ID,
			&product.SaleOrderQty,
			&product.SaleOrderPrice,
			&product.SaleOrderPriceTotal,
			&product.SaleOrderDiscount,
		); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Error occured while scanning data",
				"data":    err.Error(),
			})
		}
		DocumentData = append(DocumentData, po)

	}
	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while iterating data",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Get document successfully",
		"data":       DocumentData,
		"page":       page,
		"totalPages": TotalPage,
		"totalRows":  TotalRows,
	})
}

func DeletePoDoc(c *fiber.Ctx) error {
	var (
		TypeOfBody purchase.DocIdentity
		conn       = dbconfig.DB
	)

	if err := c.BodyParser(&TypeOfBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid body request",
			"data":    err.Error(),
		})
	}

	DocId := TypeOfBody.DocId

	_, err := conn.Exec(purchasequery.DeletePoDocument, DocId)
	if err != nil {
		log.Printf("Error while deleting document with DocId %d: %v", DocId, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error, cannot delete document",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Document deleted successfully",
	})
}
