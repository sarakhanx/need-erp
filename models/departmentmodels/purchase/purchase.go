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
	Products []productmodels.Product `json:"products"`
}

type DocIdentity struct {
	DocId         int            `json:"doc_id"`
	DocStatus     int            `json:"doc_status_id"`
	DocStatusName string         `json:"status_name"`
	ExVat         float64        `json:"ex_vat"`
	Vat           float64        `json:"vat"`
	InVat         float64        `json:"in_vat"`
	DocDiscount   float64        `json:"doc_discount"`
	DocNote       string         `json:"doc_note"`
	DocPrefix     sql.NullString `json:"doc_prefix_name"`
	DocPrefixID   int            `json:"doc_prefix_id"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}
type DocumentHeader struct {
	DocHeaderID   string `json:"doc_header_id"`
	BranchID      int    `json:"branch_id"`
	UserID        int    `json:"user_id"`
	Username      string `json:"user_name"`
	DepartmentID  int    `json:"department_id"`
	VendorData    string `json:"vendor_data"`
	BranchName    string `json:"branch_name"`
	BranchAddress string `json:"branch_address"`
}

type SaleOrder struct {
	SaleOrderID         int                     `json:"sale_order_id"`
	SaleOrderDiscount   float64                 `json:"sale_order_discount,omitempty"`
	SaleOrderPrice      float64                 `json:"sale_order_price,omitempty"`
	SaleOrderPriceTotal float64                 `json:"sale_order_price_total,omitempty"`
	Products            []productmodels.Product `json:"products"`
}
type DocumentLog struct {
	DocLogID   int    `json:"doc_log_id"`
	DocDateLog string `json:"doc_log_date"`
	DocAction  string `json:"doc_log_action"`
	DocQty     int    `json:"doc_log_qty"`
}
