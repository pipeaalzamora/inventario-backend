package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type SupplierController struct {
	supplierFacade *facades.SupplierFacade
}

func NewSupplierController(
	supplierFacade *facades.SupplierFacade,
) *SupplierController {
	return &SupplierController{
		supplierFacade: supplierFacade,
	}
}

func (c *SupplierController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("suppliers")

	r.GET("products/:supplierId/by-store/:storeId", c.getCompanyProductsBySupplierId)
	r.GET(":id", c.getSupplierByID)
	r.GET("", c.getSuppliers)

	r.POST(":id/enable-disable", c.enableDisableSupplier)
	r.POST(":id", c.updateSupplier)
	r.POST("", c.createSupplier)

	r.POST("products/:id/add", c.addSupplierProduct)
	r.POST("products/:id/update", c.updateSupplierProduct)
	r.POST("products/:id/prices", c.updateSupplierProductPrices)

	r.DELETE("products/:id/:productId", c.deleteSupplierProduct)
}

func (c *SupplierController) getSuppliers(gctx *gin.Context) {
	var queryParams shared.PageQueryParams
	if err := gctx.ShouldBindQuery(&queryParams); err != nil {
		gctx.Error(err)
		return
	}

	/* // TODO: Handle pagination defaults
	page := queryParams.Page
	if page <= 0 {
		page = 1
	}
	size := queryParams.Size
	if size <= 0 {
		size = 10
	} */

	suppliers, err := c.supplierFacade.GetAllSuppliers()
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, suppliers)
}

func (c *SupplierController) getSupplierByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	supplier, err := c.supplierFacade.GetSupplierById(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, supplier)
}

func (c *SupplierController) createSupplier(gctx *gin.Context) {
	var createRecipe recipe.RecipeCreateSupplier
	if err := shared.BindJSON(gctx, &createRecipe); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.CreateSupplier(gctx, createRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *SupplierController) updateSupplier(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var updateRecipe recipe.RecipeCreateSupplier
	if err := shared.BindJSON(gctx, &updateRecipe); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.UpdateSupplier(gctx, params.ID, &updateRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *SupplierController) enableDisableSupplier(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body struct {
		Available *bool `json:"available" binding:"required" errMsg:"El campo 'available' es obligatorio"`
	}

	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	err := c.supplierFacade.EnableDisableSupplier(params.ID, body.Available)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusOK)
}

func (c *SupplierController) getCompanyProductsBySupplierId(gctx *gin.Context) {
	type pathParams struct {
		SupplierID string `uri:"supplierId" binding:"required,uuid"`
		StoreID    string `uri:"storeId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	products, err := c.supplierFacade.GetCompanyProductsBySupplierId(params.StoreID, params.SupplierID)
	if err != nil {
		gctx.Error(err)
		return
	}
	// paginated
	gctx.JSON(http.StatusOK, products)
}

func (c *SupplierController) addSupplierProduct(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeSupplierProductCreate
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.AddProductToSupplier(gctx, params.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *SupplierController) updateSupplierProductPrices(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body []recipe.RecipeSuplierProductPriceUpdate
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.UpdateSupplierProductPrices(gctx, params.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *SupplierController) updateSupplierProduct(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeSupplierProductCreate
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.UpdateSupplierProduct(gctx, params.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *SupplierController) deleteSupplierProduct(gctx *gin.Context) {
	type pathParams struct {
		ID        string `uri:"id" binding:"required,uuid"`
		ProductID string `uri:"productId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	response, err := c.supplierFacade.DeleteSupplierProduct(gctx, params.ID, params.ProductID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}
