package recipe

import "time"

type RecipeCreateInventoryCount struct {
	StoreID      string     `json:"storeId" binding:"required"`
	WarehouseID  string     `json:"warehouseId" binding:"required"`
	DueDate      *time.Time `json:"dueDate" binding:"omitempty"` // Optional
	Observations string     `json:"observations"`
	AssignedTo   *string    `json:"assignedTo"`
	ProductsID   []string   `json:"productsId" binding:"required"`
}

type RecipeInventoryCount struct {
	StoreID     string            `json:"storeId" binding:"required,uuid"`
	WarehouseID string            `json:"warehouseId" binding:"required,uuid"`
	Items       []RecipeCountItem `json:"items" binding:"required"`
}

type RecipeCountItem struct {
	ProductID  string             `json:"productId" binding:"required"`
	Completed  bool               `json:"completed" binding:"required"`
	UnitsCount []RecipeUnitsCount `json:"unitsCount" binding:"required"`
}

type RecipeUnitsCount struct {
	UnitId  int     `json:"unitId" binding:"required"`
	UnitAbv string  `json:"unitAbv" binding:"required"`
	Count   float32 `json:"count" binding:"required"`
}

type RecipeChangeInventoryCountAssigned struct {
	NewId *string `json:"newId"`
}

type RecipeChangeInventoryCountDate struct {
	NewDate time.Time `json:"newDate" binding:"required"`
}
