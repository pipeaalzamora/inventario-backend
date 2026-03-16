package recipe

type WarehouseRecipe struct {
	CompanyId        string `json:"companyId" binding:"required"`
	StoreId          string `json:"storeId" binding:"required"`
	WarehouseName    string `json:"warehouseName" binding:"required"`
	Description      string `json:"description"`
	WarehouseAddress string `json:"warehouseAddress"`
	//DeliveryInstructions string        `json:"delivery_instructions"`
	//WorkingHours         *WorkingHours `json:"working_hours"`
	//WorkingTimezone      string        `json:"working_timezone"`
}

/*
type WorkingHours struct {
	Monday    []string `json:"monday"`
	Tuesday   []string `json:"tuesday"`
	Wednesday []string `json:"wednesday"`
	Thursday  []string `json:"thursday"`
	Friday    []string `json:"friday"`
	Saturday  []string `json:"saturday"`
	Sunday    []string `json:"sunday"`
}*/
