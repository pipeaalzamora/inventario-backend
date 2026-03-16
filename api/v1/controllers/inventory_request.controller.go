package controllers

import (
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type InventoryRequestController struct {
	inventoryRequestFacade *facades.InventoryRequestFacade
}

func NewInventoryRequestController(inventoryRequestFacade *facades.InventoryRequestFacade) *InventoryRequestController {
	return &InventoryRequestController{
		inventoryRequestFacade: inventoryRequestFacade,
	}
}

func (c *InventoryRequestController) RegisterRoutes(rg *gin.RouterGroup) {
	// r := rg.Group("/inventory-requests")

	// r.POST("", c.CreateInventoryRequest)
	// r.PUT(":id/approve", c.ApproveInventoryRequest)

	// // Assuming ChangeInventoryRequestStatus is a method to handle status changes
	// r.PATCH(":id/cancel", c.CancelInventoryRequest)
	// r.PATCH(":id/reject", c.RejectInventoryRequest)
	// r.GET(":id", c.GetInventoryRequestByID)
	// r.PUT(":id", c.UpdateInventoryRequest)
	// r.GET("by-store/:storeId", c.GetAllInventoryRequests)
}

/*
func (c *InventoryRequestController) GetAllInventoryRequests(gctx *gin.Context) {
	type pathParams struct {
		StoreId string `uri:"storeId" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var queryParams shared.PageQueryParams
	if err := gctx.ShouldBindQuery(&queryParams); err != nil {
		gctx.Error(err)
		return
	}

	// TODO: Handle pagination defaults
	page := queryParams.Page
	if page <= 0 {
		page = 1
	}
	size := queryParams.Size
	if size <= 0 {
		size = 10
	}

	requests, err := c.inventoryRequestFacade.GetAllInventoryRequests(gctx, params.StoreId, page, size, nil)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, requests)
}
func (c *InventoryRequestController) GetInventoryRequestByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	request, err := c.inventoryRequestFacade.GetInventoryRequestByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, request)
}

func (c *InventoryRequestController) CreateInventoryRequest(gctx *gin.Context) {
	var request *recipe.RecipeInventoryRequest
	if err := gctx.ShouldBindJSON(&request); err != nil {
		gctx.Error(err)
		return
	}

	newRequest, err := c.inventoryRequestFacade.CreateInventoryRequest(gctx, request)
	if err != nil {
		gctx.Error(err)
		return
	}

	newRequestDetail, err := c.inventoryRequestFacade.GetInventoryRequestByID(gctx, newRequest.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, newRequestDetail)
}

func (c *InventoryRequestController) UpdateInventoryRequest(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var request *recipe.RecipeInventoryRequest
	if err := gctx.ShouldBindJSON(&request); err != nil {
		gctx.Error(err)
		return
	}

	updatedRequest, err := c.inventoryRequestFacade.UpdateInventoryRequest(gctx, params.ID, request)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, updatedRequest)
}

func (c *InventoryRequestController) ApproveInventoryRequest(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe *recipe.RecipeInventoryRequest
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	updated, err := c.inventoryRequestFacade.ApprovedAndUpdate(gctx, params.ID, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, updated)
}

func (c *InventoryRequestController) RejectInventoryRequest(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe *recipe.RecipeInventoryRequestStatus
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	updated, err := c.inventoryRequestFacade.ChangeInventoryRequestStatus(gctx, params.ID, "rejected", recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, updated)
}

func (c *InventoryRequestController) CancelInventoryRequest(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe *recipe.RecipeInventoryRequestStatus
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	updated, err := c.inventoryRequestFacade.ChangeInventoryRequestStatus(gctx, params.ID, "cancelled", recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, updated)
}*/
