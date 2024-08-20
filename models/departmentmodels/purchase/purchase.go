package purchase

import (
	"database/sql"

	"github.com/need/go-backend/models/productmodels"
)

type PurchaseOrderDocument struct {
	DocIdentity
	DocumentHeader
	SaleOrder
	DocumentLog
}

type DocIdentity struct {
	DocId       int          `json:"doc_id"`
	Doc_Status  string       `json:"doc_status"`
	ExVat       float64      `json:"ex_vat"`
	Vat         float64      `json:"vat"`
	InVat       float64      `json:"in_vat"`
	DocDiscount float64      `json:"doc_discount"`
	DocNote     string       `json:"doc_note"`
	DocPrefix   string       `json:"doc_prefix"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
type DocumentHeader struct {
	DocHeaderID  string `json:"doc_header_id"`
	BranchID     int    `json:"branch_id"`
	UserID       int    `json:"user_id"`
	DepartmentID int    `json:"department_id"`
	VendorData   string `json:"vendor_data"`
}

type SaleOrder struct {
	SaleOrderID         string  `json:"sale_order_id"`
	SaleOrderDiscount   float64 `json:"sale_order_discount"`
	SaleOrderPrice      float64 `json:"sale_order_price"`
	SaleOrderPriceTotal float64 `json:"sale_order_price_total"`
	productmodels.Product
}
type DocumentLog struct {
	DocLogID   int    `json:"doc_log_id"`
	DocDateLog string `json:"doc_log_date"`
	DocAction  string `json:"doc_log_action"`
	DocQty     int    `json:"doc_log_qty"`
}
