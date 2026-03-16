package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	storeFacade *facades.StoreFacade
}

func NewStoreController(storeFacade *facades.StoreFacade) *StoreController {
	return &StoreController{
		storeFacade: storeFacade,
	}
}

// Registro de routes
func (sc *StoreController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("stores")

	r.GET("by-company/:companyId", sc.getStoresRequestRestriction)
	r.GET(":id", sc.getStore)
	r.POST("create", sc.createStore)
	r.POST(":id", sc.updateStore)
	// @deprecated - Use POST /companies/:id/suppliers
	//r.POST(":id/suppliers", sc.updateStoreSuppliers)
}

func (sc *StoreController) getStore(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	store, err := sc.storeFacade.GetStoreByID(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	if store == nil {
		gctx.JSON(http.StatusNotFound, gin.H{"message": "Tienda no encontrada"})
		return
	}

	gctx.JSON(http.StatusOK, store)
}

func (sc *StoreController) getStoresRequestRestriction(gctx *gin.Context) {
	type pathParams struct {
		CompanyId string `uri:"companyId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	stores, err := sc.storeFacade.GetStoresByCompanyID(gctx, params.CompanyId)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"stores": stores})
}

func (sc *StoreController) createStore(gctx *gin.Context) {
	var recipe recipe.StoreRecipe
	if err := shared.BindJSON(gctx, &recipe); err != nil {
		gctx.Error(err)
		return
	}

	store, err := sc.storeFacade.CreateStore(gctx, &recipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, store)

}

func (sc *StoreController) updateStore(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var updateRecipe recipe.UpdateStoreRecipe
	if err := shared.BindJSON(gctx, &updateRecipe); err != nil {
		gctx.Error(err)
		return
	}

	store, err := sc.storeFacade.UpdateStore(gctx, params.Id, &updateRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, store)
}

// func (sc *StoreController) updateStoreSuppliers(gctx *gin.Context) {
// 	gctx.JSON(http.StatusGone, gin.H{"error": "Endpoint deprecado. Use POST /companies/:id/suppliers"})
// }
