package purchase

const (
	CreateStatusTable = `
CREATE TABLE IF NOT EXISTS doc_status (
	doc_status_id SERIAL PRIMARY KEY,
	status_name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

	CreateDefaultDocPrefix = `
CREATE TABLE IF NOT EXISTS doc_prefix (
    doc_prefix_id SERIAL PRIMARY KEY,
    prefix_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

	CreateDocumentTable = `
CREATE TABLE IF NOT EXISTS Documents ( -- //NOTE
    doc_id SERIAL PRIMARY KEY,
    doc_log_id INT,
    doc_status_id INT NOT NULL,
    ex_vat NUMERIC(12, 2) NOT NULL,
    vat NUMERIC(12, 2) NOT NULL,
    in_vat NUMERIC(12, 2) NOT NULL,
    doc_discount NUMERIC(12, 2) NOT NULL,
    doc_note TEXT, -- //NOTE
    doc_prefix_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (doc_status_id) REFERENCES doc_status(doc_status_id),
    FOREIGN KEY (doc_prefix_id) REFERENCES doc_prefix(doc_prefix_id)
);
`

	InsertDefaultDocStatus = `
INSERT INTO doc_status (doc_status_id, status_name) VALUES
(1, 'Draft'),
(2, 'Waiting'),
(3, 'Validation'),
(4, 'Success'),
(5, 'Pending'),
(6, 'Done')
ON CONFLICT DO NOTHING;
`

	CreateDocLog = `
CREATE TABLE IF NOT EXISTS Documentlog (
	doc_log_id SERIAL PRIMARY KEY,
	doc_id INT,
	date TIMESTAMP NOT NULL,
	doc_action VARCHAR(50) NOT NULL,
	qty INT NOT NULL,
	user_id INT,
	department_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (doc_id) REFERENCES Documents(doc_id),
	FOREIGN KEY (user_id) REFERENCES Users(user_id),
	FOREIGN KEY (department_id) REFERENCES Departments(department_id)
);
`
	CreateDocumentHeader = `
CREATE TABLE IF NOT EXISTS DocumentHeader ( -- //NOTE
    doc_header_id SERIAL PRIMARY KEY,
    doc_id INTEGER NOT NULL REFERENCES Documents(doc_id),
    branch_id INTEGER NOT NULL REFERENCES Branches(branch_id),
    user_id INTEGER NOT NULL REFERENCES Users(user_id),
    department_id INTEGER NOT NULL REFERENCES Departments(department_id),
    vendor_data TEXT, -- //NOTE
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
	CreateSaleOrder = `
CREATE TABLE IF NOT EXISTS sale_order (
    sale_order_id SERIAL PRIMARY KEY,
    sale_order_discount NUMERIC(10, 2),
    sale_order_price NUMERIC(10, 2),
    sale_order_price_total NUMERIC(10, 2),
    sale_order_qty INTEGER,
    product_id INTEGER NOT NULL REFERENCES Products(product_id),
    doc_id INTEGER NOT NULL REFERENCES Documents(doc_id)
);
`
)
