package branchesquery

const (
	CreateBranchesTable = `
	CREATE TABLE IF NOT EXISTS Branches (
		branch_id SERIAL PRIMARY KEY,
		branch_name VARCHAR(255) NOT NULL,
		branch_address VARCHAR(255) NOT NULL
	);
`
	CreateBranchProductQty = `
	CREATE TABLE IF NOT EXISTS BranchProductQty (
		branch_id INT,
		product_id INT,
		qty INT NOT NULL,
		PRIMARY KEY (branch_id, product_id),
		FOREIGN KEY (branch_id) REFERENCES Branches(branch_id),
		FOREIGN KEY (product_id) REFERENCES Products(product_id)
	);`
	CreateBranch    = `INSERT INTO Branches (branch_name, branch_address) VALUES ($1, $2);`
	GetAllBranches  = `SELECT * FROM Branches;`
	GetBranchByName = `SELECT * FROM Branches WHERE branch_name = $1;`
	DeleteBranch    = `DELETE FROM Branches WHERE branch_name = $1;`
	UpdateBranch    = `UPDATE Branches SET branch_name = $1, branch_address = $2, updated_at = NOW() WHERE branch_name = $3 RETURNING branch_id, branch_name, branch_address, updated_at;`
)
