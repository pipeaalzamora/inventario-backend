package recipe

type AssignSuppliersToCompanyRecipe struct {
	SupplierIds []string `json:"supplierIds" binding:"required" errMsg:"Los IDs de proveedores son obligatorios"`
}
