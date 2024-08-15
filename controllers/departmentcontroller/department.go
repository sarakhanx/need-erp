package departmentcontroller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/need/go-backend/config/db-config"
	"github.com/need/go-backend/models/departmentmodels"
	"github.com/need/go-backend/queries/departmentquery"
)

// TODO : Waiting to make the Validation
func CreateDepartment(c *fiber.Ctx) error {
	conn := dbconfig.DB
	var departmentFormData departmentmodels.Department
	//EXPLAIN - ถ้าหากไม่มี Body ให้ส่ง Error 400
	if err := c.BodyParser(&departmentFormData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// EXPLAIN - ตรวจสอบว่ามีค่าว่างหรือไม่
	if departmentFormData.Name == "" ||
		departmentFormData.Position == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "All fields are required"})
	}

	//EXPLAIN - หากเกิด Error ในขั้นตอนบันทึก Database ให้ส่ง Error 400
	_, err := conn.Exec(departmentquery.CreateDepartment, departmentFormData.Name, departmentFormData.Position)
	if err != nil {
		log.Println("Error creating department:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error creating department", "data": err})
	}
	//EXPLAIN - หากบันทึกสำเร็จ ให้ส่ง Response 201 พร้อม Data
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "Success", "message": "Departments was created", "data": departmentFormData})
}

func GetAllDepartments(c *fiber.Ctx) error {
	conn := dbconfig.DB

	row, err := conn.Query(departmentquery.GetAllDepartments)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error", "data": err})
	}
	data := []departmentmodels.Department{}
	for row.Next() {
		var department departmentmodels.Department
		if err := row.Scan(&department.ID, &department.Name, &department.Position, &department.Member, &department.CreatedDate, &department.UpdatedDate); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error", "data": err.Error()})
		}
		data = append(data, department)
	}
	if err := row.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error", "data": err.Error()})
	}
	// defer dbconfig.DB.Close()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": data})
}
func GetDepartment(c *fiber.Ctx) error {
	conn := dbconfig.DB
	//EXPLAIN - ความสำคัญของ Param ต้องเป็นตัวเล็ก ตัวใหญ่ตาม SQL นะ เรา Value ที่ Name
	Params := c.Params("department")
	var data departmentmodels.Department
	err := conn.QueryRow(departmentquery.GetDepartment, Params).Scan(&data.ID, &data.Name, &data.Position, &data.Member, &data.CreatedDate, &data.UpdatedDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error", "data": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": data})
}
func DeleteDepartments(c *fiber.Ctx) error {
	conn := dbconfig.DB

	//EXPLAIN - ความสำคัญของ Param ต้องเป็นตัวเล็ก ตัวใหญ่ตาม SQL นะ เรา Value ที่ Name
	Params := c.Params("department")
	_, err := conn.Exec(departmentquery.DeleteDepartment, Params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Not found content", "data": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "message": "Department was deleted"})
}
