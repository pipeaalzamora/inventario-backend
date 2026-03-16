package controllers

import (
	"encoding/json"
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"
	"sofia-backend/types"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productFacade *facades.ProductFacade
}

func NewProductController(productFacade *facades.ProductFacade) *ProductController {
	return &ProductController{productFacade: productFacade}
}

func (ctrl *ProductController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("products")

	r.GET("codes", ctrl.getAllCodes)

	categoryGroup := r.Group("categories")
	categoryGroup.GET("", ctrl.getAllCategories)
	categoryGroup.POST("", ctrl.createCategory)

	r.GET("p/:id", ctrl.getById)
	r.GET("ps/:companyId/:id", ctrl.getWithSuppliersById)
	r.GET(":storeId/request-restriction", ctrl.getProductsRequestRestriction)

	// Rutas generales después
	r.GET("", ctrl.getAllProducts)
	r.POST("", ctrl.createProduct)
	r.POST(":id", ctrl.updateProduct)
	//r.DELETE(":id", ctrl.deleteProduct)

	// Rutas específicas primero

}

func (ctrl *ProductController) getById(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	product, err := ctrl.productFacade.GetProductById(params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, product)
}

func (ctrl *ProductController) getWithSuppliersById(gctx *gin.Context) {
	type pathParams struct {
		ID        string `uri:"id" binding:"required,uuid"`
		CompanyID string `uri:"companyId" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	product, err := ctrl.productFacade.GetProductWithSuppliersById(params.CompanyID, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, product)
}

func (ctrl *ProductController) getAllCodes(gctx *gin.Context) {
	products, err := ctrl.productFacade.GetAllCodes()
	if err != nil {
		gctx.Error(err)
		return
	}
	// paginated
	gctx.JSON(200, products)
}

func (ctrl *ProductController) getAllCategories(gctx *gin.Context) {
	categories, err := ctrl.productFacade.GetAllCategories()
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, categories)
}

func (ctrl *ProductController) createCategory(gctx *gin.Context) {
	var categoryRecipe recipe.RecipeProductCategory

	if err := shared.BindJSON(gctx, &categoryRecipe); err != nil {
		gctx.Error(err)
		return
	}

	category, err := ctrl.productFacade.CreateCategory(gctx, &categoryRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, category)
}

func (ctrl *ProductController) getAllProducts(gctx *gin.Context) {
	products, err := ctrl.productFacade.GetAllProducts()
	if err != nil {
		gctx.Error(err)
		return
	}
	// paginated
	gctx.JSON(200, products)
}

func (ctrl *ProductController) getProductsRequestRestriction(gctx *gin.Context) {
	type pathParams struct {
		StoreId string `uri:"storeId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	products, err := ctrl.productFacade.GetProductsRequestRestriction(gctx, params.StoreId)
	if err != nil {
		gctx.Error(err)
		return
	}
	// paginated
	gctx.JSON(200, products)
}

func (ctrl *ProductController) createProduct(gctx *gin.Context) {

	var productInput recipe.RecipeProductInput
	if err := shared.Bind(gctx, &productInput); err != nil {
		gctx.Error(err)
		return
	}

	var codesRecipe []recipe.RecipeProductCode
	err := json.Unmarshal([]byte(productInput.Codes), &codesRecipe)
	if err != nil {
		gctx.Error(types.ThrowRecipe("El formato de los códigos no es válido", "codes"))
		return
	}
	productInput.CodesList = codesRecipe

	product, err := ctrl.productFacade.CreateProduct(gctx, &productInput)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, product)
}

func (ctrl *ProductController) updateProduct(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var productInput recipe.RecipeProductInput
	if err := shared.Bind(gctx, &productInput); err != nil {
		gctx.Error(err)
		return
	}

	var codesRecipe []recipe.RecipeProductCode
	err := json.Unmarshal([]byte(productInput.Codes), &codesRecipe)
	if err != nil {
		gctx.Error(types.ThrowRecipe("El formato de los códigos no es válido", "codes"))
		return
	}
	productInput.CodesList = codesRecipe

	product, err := ctrl.productFacade.UpdateProduct(gctx, params.ID, &productInput)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, product)
}

/*
func (ctrl *ProductController) deleteProduct(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		errors := shared.FormatValidationErrors(err, &params)
		shared.SendValidationErrorResponse(gctx, errors)
		return
	}

	err := ctrl.productFacade.DeleteProduct(gctx.Request.Context(), params.ID)
	if err != nil {
		shared.SendErrorResponse(gctx, err)
		return
	}

	gctx.Status(http.StatusNoContent)
}
*/
