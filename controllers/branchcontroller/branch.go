package branchcontroller

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"

	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/middlewares/branchesvalidator"
	"github.com/need/go-backend/models/branchmodels"
	"github.com/need/go-backend/queries/branchesquery"
)

func CreateBranch(c *fiber.Ctx) error {

	var BranchFromData branchmodels.Branch
	if err := c.BodyParser(&BranchFromData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Data is invalid, Please Review your input.",
			"data":    err,
		})
	}
	if BranchFromData.BranchName == "" ||
		BranchFromData.BranchAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "All fields are required.",
		})
	}
	exist, err := branchesvalidator.CheckBranchExisted(BranchFromData.BranchName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error checking branch existence",
			"data":    err.Error(),
		})
	}
	if exist {
		log.Println("Branch name is existed, Please use another name")
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Branch name is existed, Please use another name",
		})
	}
	conn := dbconfig.DB
	result, err := conn.Exec(branchesquery.CreateBranch, BranchFromData.BranchName, BranchFromData.BranchAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create Branch",
			"data":    err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"result": result,
		"data":   BranchFromData,
	})
}

func GetAllBranches(c *fiber.Ctx) error {
	conn := dbconfig.DB

	rows, err := conn.Query(branchesquery.GetAllBranches)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":   "error",
			"response": "Error",
			"data":     err.Error(),
		})
	}
	Branches := []branchmodels.Branch{}
	for rows.Next() {
		var Branch branchmodels.Branch
		if err := rows.Scan(&Branch.ID, &Branch.BranchName, &Branch.BranchAddress, &Branch.Updated_at); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":   "error",
				"response": "Error",
				"data":     err.Error(),
			})
		}
		Branches = append(Branches, Branch)
	}
	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":   "error",
			"response": "Error",
			"data":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"response": "Success",
		"data":     Branches,
	})
}

func GetBranch(c *fiber.Ctx) error {
	conn := dbconfig.DB
	Params := c.Params("branch")

	var Branch branchmodels.Branch
	if err := conn.QueryRow(branchesquery.GetBranchByName, Params).Scan(
		&Branch.ID,
		&Branch.BranchName,
		&Branch.BranchAddress,
		&Branch.Updated_at,
	); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Not found",
				"data":    err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get Branch successful",
		"data":    Branch,
	})
}

func DeleteBranch(c *fiber.Ctx) error {
	conn := dbconfig.DB
	Params := c.Params("branch")

	result, err := conn.Exec(branchesquery.DeleteBranch, Params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"data":    err.Error(),
		})
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"data":    err.Error(),
		})
	}
	if RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Delete Branch successful",
		"result":  result,
	})
}

func UpdateBranch(c *fiber.Ctx) error {
	var BranchFromData branchmodels.Branch
	Params := c.Params("branch")
	if err := c.BodyParser(&BranchFromData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error While Parsing",
			"data":    err.Error(),
		})
	}

	if BranchFromData.BranchAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Branch Address cannot be empty",
			"data":    BranchFromData.BranchAddress,
		})
	}
	conn := dbconfig.DB

	err := conn.QueryRow(branchesquery.UpdateBranch, BranchFromData.BranchName, BranchFromData.BranchAddress, Params).Scan(
		&BranchFromData.ID,
		&BranchFromData.BranchName,
		&BranchFromData.BranchAddress,
		&BranchFromData.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Not found",
				"data":    err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error While Updating",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Update Branch successful",
		"result":  BranchFromData,
	})
}
