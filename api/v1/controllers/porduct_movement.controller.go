package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type ProductMovementController struct {
	facade *facades.ProductMovementFacade
}

func NewProductMovementController(facade *facades.ProductMovementFacade) *ProductMovementController {
	return &ProductMovementController{
		facade: facade,
	}
}

func (c *ProductMovementController) RegisterRoutes(rg *gin.RouterGroup) {
	routeGroup := rg.Group("product-movements")
	routeGroup.GET("", c.getAllProductMovements)
	routeGroup.GET(":id", c.getProductMovementByID)

	routeGroup.GET("company/:companyId", c.getProductMovementsByCompanyID)
	routeGroup.GET("store/:storeId", c.getProductMovementsByStoreID)
	routeGroup.GET("store/:storeId/product/:storeProductId", c.getProductMovementsByStoreProductId)
	routeGroup.GET("warehouse/:warehouseId", c.getProductMovementsByWarehouse)

	routeGroup.POST("warehouses", c.getProductMovementsByWharehouseIDs)
	routeGroup.POST("new-output", c.createNewOutputMovement)
	routeGroup.POST("new-input", c.createNewInputMovement)
	routeGroup.POST("transfer", c.createTransferMovement)

}

func (c *ProductMovementController) getAllProductMovements(gctx *gin.Context) {
	result, err := c.facade.GetAllProductMovements()
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) getProductMovementByID(gctx *gin.Context) {
	type Params struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params Params
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.GetProductMovementByID(params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) getProductMovementsByCompanyID(gctx *gin.Context) {
	type Params struct {
		CompanyID string `uri:"companyId" binding:"required,uuid"`
	}
}

func (c *ProductMovementController) getProductMovementsByStoreID(gctx *gin.Context) {
	type Params struct {
		StoreID string `uri:"storeId" binding:"required,uuid"`
	}

	var params Params
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.GetAllProductMovementsByStoreID(gctx, params.StoreID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) getProductMovementsByStoreProductId(gctx *gin.Context) {
	type Params struct {
		StoreID        string `uri:"storeId" binding:"required,uuid"`
		StoreProductID string `uri:"storeProductId" binding:"required,uuid"`
	}

	var params Params
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.GetAllProductMovementsByStoreProductID(gctx, params.StoreID, params.StoreProductID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) getProductMovementsByWharehouseIDs(gctx *gin.Context) {
	var recipe recipe.RecipeProductMovementByWarehouseIDs
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.GetProductMovementsByWharehouseIDs(recipe.WarehouseIDs)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) getProductMovementsByWarehouse(gctx *gin.Context) {
	type Params struct {
		WarehouseID string `uri:"warehouseId" binding:"required,uuid"`
	}

	var params Params
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.GetAllProductMovementsByDateRange(params.WarehouseID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *ProductMovementController) createNewOutputMovement(gctx *gin.Context) {
	var recipe recipe.RecipeNewMovement
	if err := shared.BindJSON(gctx, &recipe); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.CreateNewOutputMovement(gctx, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, result)

}

func (c *ProductMovementController) createNewInputMovement(gctx *gin.Context) {
	var recipe recipe.RecipeNewMovement
	if err := shared.BindJSON(gctx, &recipe); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.CreateNewInputMovement(gctx, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, result)
}

func (c *ProductMovementController) createTransferMovement(gctx *gin.Context) {
	var transferRecipe recipe.RecipeTransferMovement
	if err := shared.BindJSON(gctx, &transferRecipe); err != nil {
		gctx.Error(err)
		return
	}

	result, err := c.facade.CreateTransferMovement(gctx, transferRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, result)
}
