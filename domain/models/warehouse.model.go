package models

import "time"

type ModelWarehouse struct {
	ID                  string `json:"id" bson:"id"`
	StoreId             string `json:"storeId" bson:"storeId"`
	WarehouseName       string `json:"warehouseName" bson:"warehouseName"`
	Description         string `json:"description" bson:"description"`
	WarehouseAddress    string `json:"warehouseAddress" bson:"warehouseAddress"`
	WarehousePhone      string `json:"warehousePhone" bson:"warehousePhone"`
	IsMomeventWarehouse bool   `json:"isMomeventWarehouse" bson:"isMomeventWarehouse"`
	//DeliveryInstructions string            `json:"delivery_instructions" bson:"delivery_instructions"`
	//WorkingHours         ModelWorkingHours `json:"working_hours" bson:"working_hours"`
	//WorkingTimezone      string            `json:"working_timezone" bson:"working_timezone"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

/*
type ModelWorkingHours struct {
	Monday    []string `json:"monday"`
	Tuesday   []string `json:"tuesday"`
	Wednesday []string `json:"wednesday"`
	Thursday  []string `json:"thursday"`
	Friday    []string `json:"friday"`
	Saturday  []string `json:"saturday"`
	Sunday    []string `json:"sunday"`
}
*/
