package controllers

import (
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type SupplierOCController struct {
	facade *facades.SupplierOCFacade
}

func NewSupplierOCController(facade *facades.SupplierOCFacade) *SupplierOCController {
	return &SupplierOCController{
		facade: facade,
	}
}

func (c *SupplierOCController) RegisterRoutes(rg *gin.RouterGroup) {
	// supplierOC := rg.Group("/supplier-oc")

	// supplierOC.GET("", c.GetSupplierOC)
	// supplierOC.POST("", c.UpdateSupplierOC)
}

/*
func (c *SupplierOCController) GetSupplierOC(gctx *gin.Context) {
	type queryParam struct {
		Token string `form:"token" binding:"required"`
	}

	var params queryParam
	if err := gctx.ShouldBindQuery(&params); err != nil {
		gctx.Error(err)
		return
	}

	// Call the facade to get the supplier OC
	supplierOC, err := c.facade.GetSupplierOC(gctx, params.Token)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, supplierOC)

}

func (c *SupplierOCController) UpdateSupplierOC(gctx *gin.Context) {
	type queryParam struct {
		Token string `form:"token" binding:"required"`
	}

	var params queryParam
	if err := gctx.ShouldBindQuery(&params); err != nil {
		gctx.Error(err)
		return
	}

	var recipe recipe.SupplierOCRecipe
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	err := c.facade.UpdateSupplierOC(gctx, params.Token, recipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"message": "supplier OC updated successfully"})

}
*/
