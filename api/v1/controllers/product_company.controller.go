package controllers

import (
	"net/http"
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

// ProductCompanyController is deprecated - use ProductPerStoreController instead
// This controller now routes to the new store-based product functionality
type ProductCompanyController struct {
	productFacade *facades.ProductFacade
}

func NewProductCompanyController(productFacade *facades.ProductFacade) *ProductCompanyController {
	return &ProductCompanyController{productFacade: productFacade}
}

func (ctrl *ProductCompanyController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/products-company")

	// Deprecated: use /products-store/by-store/:storeId instead
	r.GET("/by-store/:storeId", ctrl.getStoreProductsByStoreId)
}

func (ctrl *ProductCompanyController) getStoreProductsByStoreId(gctx *gin.Context) {
	type pathParams struct {
		StoreId string `uri:"storeId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	products, err := ctrl.productFacade.GetStoreProductsByStoreId(gctx, params.StoreId)
	if err != nil {
		gctx.Error(err)
		return
	}
	// paginated
	gctx.JSON(http.StatusOK, products)
}
