package models

import "time"

type StoreModel struct {
	ID           string    `json:"id"`
	CompanyID    string    `json:"companyId"`
	StoreName    string    `json:"storeName"`
	StoreAddress string    `json:"storeAddress"`
	Description  string    `json:"description"`
	IDCostCenter string    `json:"idCostCenter"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	WareHouses []ModelWarehouse `json:"warehouses"`

	//SupplierApplied []StoreSupplierModel `json:"supplierApplied"`
}

/*
type StoreSupplierModel struct {
	SupplierID   string `json:"id"`
	SupplierName string `json:"supplier_name"`
	Available    bool   `json:"available"`
}
*/
