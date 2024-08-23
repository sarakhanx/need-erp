package purchasevalidation

import (
	"errors"

	"github.com/need/go-backend/models/departmentmodels/purchase"
)

func ValidatePurchaseOrderInput(data purchase.PurchaseOrderDocument) error {
	if data.DocId == 0 {
		return errors.New("DocId is required and cannot be zero")
	}
	if data.DocStatus == 0 {
		return errors.New("DocStatus is required and cannot be zero")
	}
	if data.DocNote == "" {
		return errors.New("DocNote is required")
	}
	if data.BranchID == 0 {
		return errors.New("BranchID is required and cannot be zero")
	}
	if data.UserID == 0 {
		return errors.New("UserID is required and cannot be zero")
	}
	if data.DepartmentID == 0 {
		return errors.New("DepartmentID is required and cannot be zero")
	}
	if data.VendorData == "" {
		return errors.New("VendorData is required")
	}
	if data.DocPrefixID == 0 {
		return errors.New("DocPrefixID is required and cannot be zero")
	}
	return nil
}
