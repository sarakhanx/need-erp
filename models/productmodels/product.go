package productmodels

type Product struct {
	ID            int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Cost          float64 `json:"cost"`
	Price         float64 `json:"price"`
	Total_QTY     int     `json:"total_qty"`
	Log_QTY       int     `json:"log_qty"`
	BranchID      int     `json:"branch_id"`
	Action        string  `json:"action"`
	UserID        int     `json:"user_id"`
	Category      string  `json:"category"`
	BranchName    string  `json:"branch_name"`
	Updated_at    string  `json:"updated_at"`
	BranchAddress string  `json:"branch_address"`
	Branch_QTY    int     `json:"branch_qty"`
}

type ProductLog struct {
	ID        int    `json:"log_id"`
	ProductID int    `json:"product_id"`
	Date      string `json:"date"`
	Action    string `json:"action"`
	Qty       int    `json:"qty"`
	UserId    int    `json:"user_id"`
	BranchId  int    `json:"branch_id"`
}
