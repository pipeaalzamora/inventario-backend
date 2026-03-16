package dto

import "time"

type DtoInventoryCountDetail struct {
	ID              string                  `json:"id"`
	DisplayID       string                  `json:"displayId"`
	StoreID         string                  `json:"storeId"`
	StoreName       string                  `json:"storeName"`
	CompanyID       string                  `json:"companyId"`
	WarehouseID     string                  `json:"warehouseId"`
	WarehouseName   string                  `json:"warehouseName"`
	CreatedBy       string                  `json:"createdBy"`
	CreatedByName   string                  `json:"createdByName"`
	AssignedTo      *string                 `json:"assignedTo"`
	AssignedToName  *string                 `json:"assignedToName"`
	Status          string                  `json:"status"`
	ScheduledAt     time.Time               `json:"scheduledAt"`
	CompletedAt     *time.Time              `json:"completedAt"`
	CreatedAt       time.Time               `json:"createdAt"`
	UpdatedAt       time.Time               `json:"updatedAt"`
	MovementTrackId string                  `json:"movementTrackId"`
	CountItems      []DtoInventoryCountItem `json:"countItems"`
}

type DtoInventoryCountGeneral struct {
	ID              string     `json:"id"`
	DisplayID       string     `json:"displayId"`
	StoreID         string     `json:"storeId"`
	StoreName       string     `json:"storeName"`
	CompanyID       string     `json:"companyId"`
	WarehouseID     string     `json:"warehouseId"`
	WarehouseName   string     `json:"warehouseName"`
	CreatedBy       string     `json:"createdBy"`
	CreatedByName   string     `json:"createdByName"`
	AssignedTo      *string    `json:"assignedTo"`
	AssignedToName  *string    `json:"assignedToName"`
	Status          string     `json:"status"`
	ScheduledAt     time.Time  `json:"scheduledAt"`
	CompletedAt     *time.Time `json:"completedAt"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	MovementTrackId string     `json:"movementTrackId"`
	TotalItems      int        `json:"totalItems"`
	CompletedItems  int        `json:"completedItems"`
}

type DtoInventoryCountItem struct {
	ProductID            string          `json:"productId"`
	ProductSKU          string          `json:"productSku"`
	ProductName         string          `json:"productName"`
	ProductImage        *string         `json:"productImage"`
	Completed           bool            `json:"completed"`
	Total               float32         `json:"total"`
	UnitsCount          []DtoUnitsCount `json:"unitsCount"`
	IncidenceImageURL   *string         `json:"incidenceImageUrl"`
	IncidenceObservation *string         `json:"incidenceObservation"`
}

type DtoUnitsCount struct {
	UnitId  int     `json:"unitId"`
	UnitAbv string  `json:"unitAbv"`
	Count   float32 `json:"count"`
	Factor  float32 `json:"factor"`
}
