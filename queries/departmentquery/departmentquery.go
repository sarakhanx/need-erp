package departmentquery

const (
	CreateDepartment = `
	INSERT INTO Departments (name, position , created_at, updated_at) VALUES ($1, $2 , Now(), Now());
	`
	GetAllDepartments = ` SELECT * FROM Departments;`
	GetDepartment     = `SELECT * FROM Departments WHERE name = $1;`
	DeleteDepartment  = `DELETE FROM Departments WHERE name = $1;`
)
