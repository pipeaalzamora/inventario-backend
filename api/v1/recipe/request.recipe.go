package recipe

type RecipeNewRequest struct {
	CompanyID   string `json:"companyId" binding:"required" errMsg:"El ID de la empresa es obligatorio"`
	StoreID     string `json:"storeId" binding:"required" errMsg:"El ID de la tienda es obligatorio"`
	RequestType string `json:"requestType" binding:"required" errMsg:"El tipo de solicitud es obligatorio"`
}
