package userquery

const UserQueryDebug = `
CREATE TABLE IF NOT EXISTS DebugUser (
id SERIAL PRIMARY KEY,
name VARCHAR(255),
date TIMESTAMP
)
`

const (
	CreateDepartmentTable = `
	CREATE TABLE IF NOT EXISTS Departments (
		department_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		position VARCHAR(255) UNIQUE,
		member INTEGER,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
		
	)`
	CreateUserTable = `
	CREATE TABLE IF NOT EXISTS Users (
		user_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		lastname VARCHAR(255),
		mobile VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		role JSONB,
		position VARCHAR(255),
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		last_login TIMESTAMP,
		CONSTRAINT fk_department_position FOREIGN KEY (position) REFERENCES Departments(position)
	)
	`
	CheckExistingUser = `SELECT * FROM Users WHERE email = $1`
	InsertNewUser     = `INSERT INTO Users (name, lastname, mobile, email, password, role, position , created_at, updated_at)
                      VALUES ($1, $2, $3, $4, $5, $6, $7 , Now(), Now()) RETURNING user_id`
	SigninUser = `
					  SELECT user_id, name, lastname, mobile, email, password, role, position, last_login
					  FROM Users
					  WHERE email = $1
				  `
	UpdateLastLogin = `
	UPDATE Users
	SET last_login = $1
	WHERE user_id = $2
`
)
