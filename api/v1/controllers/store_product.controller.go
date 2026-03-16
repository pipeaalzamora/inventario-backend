package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

// StoreProductController maneja los endpoints de productos de tienda.
type StoreProductController struct {
	facade *facades.StoreProductFacade
}

// NewStoreProductController crea una nueva instancia del controlador.
func NewStoreProductController(facade *facades.StoreProductFacade) *StoreProductController {
	return &StoreProductController{
		facade: facade,
	}
}

// RegisterRoutes registra las rutas del controlador.
func (c *StoreProductController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("store-products")

	r.GET(":storeId", c.getAllByStore)
	r.GET(":storeId/suppliers", c.getProductsWithSuppliers)

	r.GET(":storeId/:id", c.getByID)
	r.GET(":storeId/:id/suppliers", c.getSuppliers)

	r.POST(":storeId", c.create)
	r.POST(":storeId/:id", c.update)
}

func (c *StoreProductController) getByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid" errMsg:"El ID del producto de tienda es obligatorio"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	product, err := c.facade.GetStoreProductByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, product)
}

func (c *StoreProductController) getSuppliers(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid" errMsg:"El ID del producto de tienda es obligatorio"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	suppliers, err := c.facade.GetSuppliersForStoreProduct(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, suppliers)
}

func (c *StoreProductController) getProductsWithSuppliers(gctx *gin.Context) {
	type pathParams struct {
		StoreID string `uri:"storeId" binding:"required,uuid" errMsg:"El ID de la tienda es obligatorio"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	products, err := c.facade.GetProductsWithSuppliersByStore(gctx, params.StoreID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, products)
}

func (c *StoreProductController) create(gctx *gin.Context) {
	var input recipe.RecipeCreateStoreProduct
	if err := shared.BindJSON(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	product, err := c.facade.CreateStoreProduct(gctx, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, product)
}

func (c *StoreProductController) update(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid" errMsg:"El ID del producto de tienda es obligatorio"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.RecipeUpdateStoreProduct
	if err := shared.BindJSON(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	product, err := c.facade.UpdateStoreProduct(gctx, params.ID, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, product)
}

func (c *StoreProductController) getAllByStore(gctx *gin.Context) {
	type pathParams struct {
		StoreID string `uri:"storeId" binding:"required,uuid" errMsg:"El ID de la tienda es obligatorio"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	products, err := c.facade.GetProductsByStore(gctx, params.StoreID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, products)
}
