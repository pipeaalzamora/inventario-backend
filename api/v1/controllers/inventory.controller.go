package controllers

import (
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryFacade *facades.InventoryFacade
}

func NewInventoryController(inventoryFacade *facades.InventoryFacade) *InventoryController {
	return &InventoryController{
		inventoryFacade: inventoryFacade,
	}
}

func (c *InventoryController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("inventory")

	//r.GET(":companyId/:storeId", c.getAllProductsReportByCompanyId)
	r.POST(":companyId/:storeId", c.getAllInventory)
	r.GET(":companyId/:storeId/:warehouseId/:productId", c.getSingleProductStock)
	//r.GET("/:companyId/:productId", c.getProductStockReportDetails)
	//r.GET("/products/:warehouseId", c.getInventoryByWarehouseId)
}

func (c *InventoryController) getAllInventory(gctx *gin.Context) {
	type pathParams struct {
		CompanyId string `uri:"companyId" binding:"required,uuid"`
		StoreId   string `uri:"storeId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var jsonParam struct {
		WarehousesId []string `json:"warehousesId" binding:"omitempty"`
	}
	if err := shared.BindJSON(gctx, &jsonParam); err != nil {
		gctx.Error(err)
		return
	}

	inventory, err := c.inventoryFacade.GetAllInventoryByStoreId(gctx, params.CompanyId, params.StoreId, jsonParam.WarehousesId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, inventory)
}

func (c *InventoryController) getSingleProductStock(gctx *gin.Context) {
	type pathParams struct {
		CompanyId   string `uri:"companyId" binding:"required,uuid"`
		StoreId     string `uri:"storeId" binding:"required,uuid"`
		WarehouseId string `uri:"warehouseId" binding:"required,uuid"`
		ProductId   string `uri:"productId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	productStock, err := c.inventoryFacade.GetSingleProductStockByWarehouse(gctx, params.CompanyId, params.StoreId, params.WarehouseId, params.ProductId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, productStock)
}

/*
func (c *InventoryController) getAllProductsReportByCompanyId(gctx *gin.Context) {

	type pathParams struct {
		CompanyId   string `uri:"companyId" binding:"required,uuid"`
		StoreId     string `uri:"storeId" binding:"required,uuid"`
		WarehouseId string `uri:"warehouseId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		errors := shared.FormatValidationErrors(err, &params)
		shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	result, err := c.inventoryFacade.GetAllProductsByCompanyId(gctx, params.CompanyId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryController) getProductStockReportDetails(gctx *gin.Context) {
	type pathParams struct {
		CompanyId string `uri:"companyId" binding:"required,uuid"`
		ProductId string `uri:"productId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		errors := shared.FormatValidationErrors(err, &params)
		shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	result, err := c.inventoryFacade.GetDetailedProduct(gctx, params.CompanyId, params.ProductId)
	if err != nil {
		shared.SendErrorResponse(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, result)

}
*/

/*
func (c *InventoryController) getInventoryByWarehouseId(gctx *gin.Context) {
	type pathParams struct {
		WarehouseId string `uri:"warehouseId" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		errors := shared.FormatValidationErrors(err, &params)
		shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	result, err := c.inventoryFacade.GetInventoryByWarehouseId(gctx, params.WarehouseId)
	if err != nil {
		shared.SendErrorResponse(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}
*/
