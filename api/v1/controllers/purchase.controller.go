package controllers

import (
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	purchaseFacade *facades.PurchaseFacade
}

func NewPurchaseController(purchaseFacade *facades.PurchaseFacade) *PurchaseController {
	return &PurchaseController{purchaseFacade: purchaseFacade}
}

func (pc *PurchaseController) RegisterRoutes(rg *gin.RouterGroup) {
	// r := rg.Group("/purchases")

	// r.GET("by-store/:storeId", pc.getPurchases)
	// r.GET("by-request/:requestId", pc.getPurchasesByInventoryRequestID)
	// r.POST("retry/:purchaseId", pc.retryWithOtherSupplier)
	// r.POST("approve/:purchaseId", pc.approvePurchase)
	// r.POST("cancel/:purchaseId", pc.cancelPurchase)
	// r.GET(":id", pc.getPurchaseByID)
	// r.POST("", pc.createPurchase)
}

/*
func (pc *PurchaseController) getPurchases(gctx *gin.Context) {
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

	requests, err := pc.purchaseFacade.GetAllPurchase(gctx, params.StoreId, page, size, nil)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, requests)
}

func (pc *PurchaseController) getPurchaseByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	request, err := pc.purchaseFacade.GetPurchaseByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, request)
}

func (pc *PurchaseController) getPurchasesByInventoryRequestID(gctx *gin.Context) {
	type pathParams struct {
		RequestId string `uri:"requestId" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	purchases, err := pc.purchaseFacade.GetPurchasesByInventoryRequestID(gctx, params.RequestId)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"purchases": purchases})
}

func (pc *PurchaseController) createPurchase(gctx *gin.Context) {
	var purchase *recipe.RecipePurchase
	if err := shared.BindJSON(gctx, &purchase); err != nil {
		gctx.Error(err)
		return
	}

	createdPurchase, err := pc.purchaseFacade.CreatePurchaseOrder(gctx, purchase)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusCreated, createdPurchase)
}

func (pc *PurchaseController) retryWithOtherSupplier(gctx *gin.Context) {
	type pathParams struct {
		PurchaseId string `uri:"purchaseId" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	purchase, err := pc.purchaseFacade.RetryWithOtherSupplier(gctx, params.PurchaseId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, purchase)
}

func (pc *PurchaseController) approvePurchase(gctx *gin.Context) {
	type pathParams struct {
		PurchaseId string `uri:"purchaseId" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	purchase, err := pc.purchaseFacade.ApprovePurchase(gctx, params.PurchaseId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, purchase)
}

func (pc *PurchaseController) cancelPurchase(gctx *gin.Context) {
	type pathParams struct {
		PurchaseId string `uri:"purchaseId" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var cancelData recipe.RecipeCancelPurchase
	if err := gctx.ShouldBindJSON(&cancelData); err != nil {
		gctx.Error(err)
		return
	}

	purchase, err := pc.purchaseFacade.CancelPurchase(gctx, params.PurchaseId, cancelData.Observation)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, purchase)
}
*/
