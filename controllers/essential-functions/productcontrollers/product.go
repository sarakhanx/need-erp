package productcontroller

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/models/branchmodels"
	"github.com/need/go-backend/models/productmodels"
	"github.com/need/go-backend/queries/productquery"
)

func DeleteProduct(c *fiber.Ctx) error {
	//TODO - กลบัมาทำเรื่องของ Procedure ใส่อีกครั้ง
	Params := c.Params("product_id")

	conn := dbconfig.DB
	tx, err := conn.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while begin transaction",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เริ่มลบสินค้าจาก Branch
	//TODO - ทำลบจากสาขาไหน และลบในจำนวนเท่าไหร่ด้วย
	_, err = tx.Exec(productquery.DeleteProductFromBranch, Params)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while delete product from branch",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เริ่มลบ Logของสินค้า
	_, err = tx.Exec(productquery.DeleteProductFromLog, Params)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while delete product from log",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เริ่มลบ Product
	//TODO - ทำลบในจำนวนเท่าไหร่ด้วย
	_, err = tx.Exec(productquery.DeleteProduct, Params)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while delete product from product",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เริ่มลบ ProductStock
	//TODO - ทำลบในจำนวนเท่าไหร่ด้วย
	_, err = tx.Exec(productquery.DeleteProductFromStock, Params)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error while delete product from stock",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - สำเร็จ Commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to commit transaction",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Product Successfully",
		"data":    nil,
	})
}

func GetHundredProducts(c *fiber.Ctx) error {
	conn := dbconfig.DB
	Page := c.Query("page")
	page, err := strconv.Atoi(Page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with query in params",
			"data":    err.Error(),
		})
	}
	limit := 100
	offset := (page - 1) * limit

	var TotalRows int
	err = conn.QueryRow(`select count(*) from products;`).Scan(&TotalRows)
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

	rows, err := conn.Query(productquery.GetAllProducts, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while excuting query",
			"data":    err.Error(),
		})
	}
	defer rows.Close()

	product := []productmodels.Product{}
	for rows.Next() {

		var p productmodels.Product

		if err := rows.Scan(&p.ID, &p.ProductName, &p.Cost, &p.Price, &p.Category); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed while scanning products",
				"data":    err.Error(),
			})
		}

		product = append(product, p)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while iterating products",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Get Products Successfully",
		"products":   product,
		"page":       page,
		"totalPages": TotalPage,
		"totalRows":  TotalRows,
	})
}

func GetProductByID(c *fiber.Ctx) error {
	//TODO : need t have information of product with created | updated | prod_log | qty | how many each branch have
	Params := c.Params("product_id")
	conn := dbconfig.DB

	var product_item productmodels.Product

	err := conn.QueryRow(productquery.GetAProductById, Params).Scan(
		&product_item.Total_QTY,     //total_qty
		&product_item.ID,            //Product_ID
		&product_item.ProductName,   //product_name
		&product_item.Cost,          //product_cost
		&product_item.Price,         //product_price
		&product_item.Category,      //product_category
		&product_item.BranchName,    //branch_name
		&product_item.BranchAddress, //branch_address
		&product_item.Branch_QTY,    //branch_qty
		&product_item.Updated_at,    //log_date
		&product_item.Action,        //log_action
		&product_item.Log_QTY,       //log_qty
		&product_item.UserID,        //user_id
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Product not found",
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while excuting query",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get Products Successfully",
		"data":    product_item,
	})
}

// NOTE - :ยังไม่่ได้ทำ Route
func GetProductWithProcedure(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get Products Successfully",
		"data":    nil,
	})
}

func CreateProductWithProcedure(c *fiber.Ctx) error {
	conn := dbconfig.DB
	var ProductFormData productmodels.Product

	err := c.BodyParser(&ProductFormData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while parsing data",
			"data":    err.Error(),
		})
	}

	if ProductFormData.Cost == 0 ||
		ProductFormData.Price == 0 ||
		ProductFormData.Category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "All fields are required",
		})
	}

	_, err = conn.Exec(productquery.CreateNewProduct,
		ProductFormData.ProductName,
		ProductFormData.Cost,
		ProductFormData.Price,
		ProductFormData.Category,
		ProductFormData.BranchID,
		ProductFormData.Total_QTY,
		ProductFormData.Action,
		ProductFormData.UserID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create product",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created Product Successfully",
	})
}

func ReplenishProduct(c *fiber.Ctx) error {
	//EXPLAIN - เอาไว้ทำการเพิ่มสินค้าเข้ามาในระบบ แบบ click เลือกที่ตรง Card Product เลย
	conn := dbconfig.DB
	var product_data productmodels.Product
	err := c.BodyParser(&product_data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Data is invalid, Please Review your input.",
			"data":    err,
		})
	}

	if product_data.ID == 0 ||
		product_data.BranchID == 0 ||
		product_data.Branch_QTY == 0 ||
		product_data.Action == "" ||
		product_data.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "All fields are required.",
		})
	}

	var result sql.Result
	result, err = conn.Exec(productquery.InsertProduct,
		product_data.ID,
		product_data.BranchID,
		product_data.Branch_QTY,
		product_data.Action,
		product_data.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to Insert product",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Replenish products successfully",
		"data":    result,
	})
}

// NOTE - อาจจะ deprecate ในอนาคต เพราะมี Procedure มาใช้งานแทน
func CreateProduct(c *fiber.Ctx) error {
	conn := dbconfig.DB
	tx, err := conn.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to start transaction",
			"data":    err.Error(),
		})
	}
	var ProductFormData productmodels.Product
	//EXPLAIN - เริ่มดึงค่าจาก Body
	err = c.BodyParser(&ProductFormData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while parsing data",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เช็คว่ามีค่าว่างหรือไม่
	if ProductFormData.Cost == 0 ||
		ProductFormData.Price == 0 ||
		ProductFormData.Category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "All fields are required",
		})
	}
	//EXPLAIN - เริ่ม Execute SQL สำหรับสร้าง Prod ใหม่พร้อม ReturnID
	err = tx.QueryRow(productquery.CreateProduct, ProductFormData.ProductName, ProductFormData.Cost, ProductFormData.Price, ProductFormData.Category).Scan(&ProductFormData.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error creating product:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating product",
			"data":    err.Error(),
		})
	}
	//EXPLAIN - เริ่ม Execute SQL สำหรับสร้าง BranchQTY โดยรับ RetuningID มา
	var BranchQTY branchmodels.BracnhProductQty
	BranchQTY.ID = ProductFormData.BranchID
	BranchQTY.QTY = ProductFormData.Total_QTY
	BranchQTY.Product_ID = ProductFormData.ID

	_, err = tx.Exec(productquery.InsertQtyToBranch, BranchQTY.ID, BranchQTY.Product_ID, BranchQTY.QTY)
	if err != nil {
		tx.Rollback()
		log.Println("Error while insert qty in branch:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while insert qty in branch",
			"data":    err.Error(),
		})
	}

	//EXPLAIN - เริ่ม Execute SQL สำหรับสร้าง ProductLog
	var ProductLog productmodels.ProductLog
	ProductLog.ProductID = ProductFormData.ID
	ProductLog.Action = ProductFormData.Action
	ProductLog.Qty = ProductFormData.Total_QTY
	ProductLog.UserId = ProductFormData.UserID
	ProductLog.BranchId = ProductFormData.BranchID

	_, err = tx.Exec(productquery.InsertLogToProductLog, ProductLog.ProductID, ProductLog.Action, ProductLog.Qty, ProductLog.UserId, ProductLog.BranchId)
	if err != nil {
		tx.Rollback()
		log.Println("Error while insert log in productlog:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while insert log in productlog",
			"data":    err.Error(),
		})
	}

	// Insert a new row into ProductStock if it doesn't exist
	_, err = tx.Exec(`INSERT INTO ProductStock (product_id, total_qty) VALUES ($1, 0) ON CONFLICT (product_id) DO NOTHING`, ProductFormData.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error while insert into ProductStock:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while insert into ProductStock",
			"data":    err.Error(),
		})
	}

	//TODO : - สั่ง Run Manual sum productstock
	_, err = tx.Exec(productquery.SumProductStock, ProductFormData.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error while summarize product stock:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while summarize product stock",
			"data":    err.Error(),
		})
	}

	//EXPLAIN - สำเร็จ Commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to commit transaction",
			"data":    err.Error(),
		})
	}

	//NOTE - เพิ่มสินค้าสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created Product Successfully",
		"data":    ProductFormData,
	})
}
