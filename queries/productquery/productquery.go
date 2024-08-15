package productquery

const (
	CreateProductsTable = `
	CREATE TABLE IF NOT EXISTS Products (
    	product_id SERIAL PRIMARY KEY,
    	product_name VARCHAR(255) NOT NULL,
    	cost NUMERIC(10, 2) NOT NULL,
    	price NUMERIC(10, 2) NOT NULL,
		category VARCHAR(255) NOT NULL
	);
`

	CreateProductLog = `
	CREATE TABLE IF NOT EXISTS ProductLog (
		log_id SERIAL PRIMARY KEY,
		product_id INT,
		date TIMESTAMP NOT NULL,
		action VARCHAR(50) NOT NULL,
		qty INT NOT NULL,
		user_id INT,
		branch_id INT,
		FOREIGN KEY (product_id) REFERENCES Products(product_id),
		FOREIGN KEY (user_id) REFERENCES Users(user_id),
		FOREIGN KEY (branch_id) REFERENCES Branches(branch_id)
	);
`

	CreateProductStock = `
	CREATE TABLE IF NOT EXISTS ProductStock (
		product_id INT PRIMARY KEY,
		total_qty INT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES Products(product_id)
	);
`
	//NOTE - ออาจจะ Deprecate ในอนาคต
	CreateProduct = `INSERT INTO Products (product_name, cost, price, category) VALUES ($1, $2, $3, $4) RETURNING product_id;`

	DeleteProduct = `DELETE FROM Products WHERE product_id = $1;`

	DeleteProductFromBranch = `DELETE FROM BranchProductQty WHERE product_id = $1;`

	DeleteProductFromStock = `DELETE FROM ProductStock WHERE product_id = $1;`

	DeleteProductFromLog = `DELETE FROM ProductLog WHERE product_id = $1;`
	//NOTE - ออาจจะ Deprecate ในอนาคต
	InsertQtyToBranch = `INSERT INTO BranchProductQty (branch_id, product_id, qty) VALUES ($1, $2, $3);`
	//NOTE - ออาจจะ Deprecate ในอนาคต
	InsertLogToProductLog = `INSERT INTO ProductLog (product_id, action, qty, user_id, branch_id , date) VALUES ($1, $2, $3, $4, $5, Now());`
	//NOTE - ออาจจะ Deprecate ในอนาคต
	SumProductStock = `
	UPDATE ProductStock
	SET total_qty = (
		SELECT SUM(qty)
		FROM BranchProductQty
		WHERE product_id = $1
	)
	WHERE product_id = $1;`

	GetAllProducts = `SELECT * FROM Products Limit $1 OFFSET $2;`

	GetAProductById = `
SELECT 
    pqty.total_qty,
    p.product_id,
    p.product_name,
    p.cost,
    p.price,
    p.category,
    b.branch_name,
    b.branch_address,
    COALESCE(bqty.qty, 0) AS branch_qty,
    log.date AS log_date,
    log.action AS log_action,
    log.qty AS log_qty,
    log.user_id
FROM 
    branchproductqty bqty
JOIN 
    products p ON p.product_id = bqty.product_id
JOIN 
    branches b ON b.branch_id = bqty.branch_id
JOIN 
    productstock pqty ON pqty.product_id = bqty.product_id
LEFT JOIN 
    productlog log ON log.product_id = bqty.product_id
    AND log.date = (
        SELECT MAX(date) 
        FROM productlog 
        WHERE product_id = bqty.product_id
    )
WHERE 
    bqty.product_id = $1;
	`

	InsertProduct = `CALL insert_or_update_product($1, $2, $3, $4, $5);`

	CreateNewProduct = `CALL create_product($1, $2, $3, $4, $5, $6, $7, $8)`
)
