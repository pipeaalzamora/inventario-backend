package recipe

type StoreRecipe struct {
	CompanyId   string `json:"companyId" binding:"required" errMsg:"El ID de la empresa es obligatorio"`
	Name        string `json:"name" binding:"required" errMsg:"El nombre es obligatorio"`
	Description string `json:"description" binding:"required" errMsg:"La descripción es obligatoria"`
	Address     string `json:"address" binding:"required" errMsg:"La dirección es obligatoria"`
	CostCenter  string `json:"costCenter" binding:"required" errMsg:"El centro de costo es obligatorio"`
}

type UpdateStoreRecipe struct {
	Name        string `json:"name" binding:"required" errMsg:"El nombre es obligatorio"`
	Description string `json:"description" binding:"required" errMsg:"La descripción es obligatoria"`
	Address     string `json:"address" binding:"required" errMsg:"La dirección es obligatoria"`
	CostCenter  string `json:"costCenter" binding:"required" errMsg:"El centro de costo es obligatorio"`
}

type UpdateStoreSuppliersRecipe struct {
	SupplierIds []string `json:"supplierIds" binding:"required" errMsg:"Los IDs de proveedores son obligatorios"`
}
