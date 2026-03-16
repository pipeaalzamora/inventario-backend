package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"
	"sofia-backend/types"

	"github.com/gin-gonic/gin"
)

type WarehouseController struct {
	facade *facades.WarehouseFacade
}

func NewWarehouseController(facade *facades.WarehouseFacade) *WarehouseController {
	return &WarehouseController{
		facade: facade,
	}
}

func (c *WarehouseController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("warehouses")

	r.GET("by-store/:storeId", c.getWarehousesByStoreId)

	r.POST("by-store/:storeId", c.CreateWarehouse)

	r.PATCH("by-store/:storeId/:warehouseId", c.UpdateWarehouse)

	//r.GET("")

	//r.DELETE("/:id", c.DeleteWarehouse)
}

func (c *WarehouseController) getWarehousesByStoreId(gctx *gin.Context) {
	type pathParams struct {
		StoreId string `uri:"storeId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	warehouses, err := c.facade.GetWarehousesByStoreId(gctx, params.StoreId)
	if err != nil {
		gctx.Error(err)
		//shared.SendErrorResponse(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, warehouses)
}

func (c *WarehouseController) CreateWarehouse(gctx *gin.Context) {
	type pathParams struct {
		StoreId string `uri:"storeId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe recipe.WarehouseRecipe
	if err := shared.Bind(gctx, &recipe); err != nil {
		gctx.Error(err)
		//errors := shared.FormatValidationErrors(err, &recipe)
		//shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	if recipe.StoreId != params.StoreId {
		gctx.Error(types.ThrowMsg("El id de la tienda no es el correcto"))
		return
	}

	warehouse, err := c.facade.CreateWarehouse(gctx, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusCreated, warehouse)
}

func (c *WarehouseController) UpdateWarehouse(gctx *gin.Context) {
	type pathParams struct {
		StoreId     string `uri:"storeId" binding:"required"`
		WarehouseId string `uri:"warehouseId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe recipe.WarehouseRecipe
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	if recipe.StoreId != params.StoreId {
		gctx.Error(types.ThrowMsg("El id de la tienda no es el correcto"))
		return
	}

	warehouse, err := c.facade.UpdateWarehouse(gctx, params.WarehouseId, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, warehouse)
}

/*
func (c *WarehouseController) DeleteWarehouse(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		errors := shared.FormatValidationErrors(err, &params)
		shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	if err := c.facade.DeleteWarehouse(gctx, params.Id); err != nil {
		shared.SendErrorResponse(gctx, err)
		return
	}
	gctx.Status(http.StatusNoContent)
}*/
