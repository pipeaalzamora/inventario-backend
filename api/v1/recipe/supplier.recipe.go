package recipe

type RecipeCreateSupplier struct {
	Name          string                  `json:"name" binding:"required"`
	Description   *string                 `json:"description"`
	FiscalName    string                  `json:"fiscalName" binding:"required"`
	IDFiscal      string                  `json:"idFiscal" binding:"required"`
	FiscalAddress string                  `json:"fiscalAddress" binding:"required"`
	FiscalState   string                  `json:"fiscalState" binding:"required"`
	FiscalCity    string                  `json:"fiscalCity" binding:"required"`
	Email         string                  `json:"email" binding:"required,email"`
	Contacts      []RecipeSupplierContact `json:"contacts"`
	Available     bool                    `json:"available"`
}

type RecipeSupplierContact struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required"`
}

type RecipeSupplierProductCreate struct {
	ProductID   string  `json:"productId" binding:"required,uuid"`
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	SKU         string  `json:"sku" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	UnitId      int     `json:"unitId" binding:"required"`
	Available   *bool   `json:"available"`
}

type RecipeSuplierProductPriceUpdate struct {
	ProductID string  `json:"productId" binding:"required,uuid"`
	Price     float32 `json:"price" binding:"required"`
}
