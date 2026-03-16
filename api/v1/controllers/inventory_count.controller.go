package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type InventoryCountController struct {
	facade facades.InventoryCountFacade
}

func NewInventoryCountController(facade facades.InventoryCountFacade) *InventoryCountController {
	return &InventoryCountController{facade: facade}
}

func (c *InventoryCountController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/inventory-counts")

	r.GET("", c.getAllInventoryCountsByUser)
	r.GET("/:id", c.getInventoryCountById)
	r.POST("/company/:companyId", c.createInventoryCount)
	r.PUT("/:id", c.updateInventoryCount)
	r.POST("/cancel/:id", c.cancelInventoryCount)
	r.POST("/start/:id", c.startInventoryCount)
	r.POST("/draft/:id", c.draftInventoryCount)
	r.POST("/commit/:id", c.commitInventoryCount)
	r.POST("/new-date/:id", c.newDateForInventoryCount)
	r.POST("/new-assigned/:id", c.newAsignedForInventoryCount)
	r.POST("/reject/:id", c.rejectInventoryCount)
	r.POST("/:id/incidence", c.saveIncidence)

}

func (c *InventoryCountController) getAllInventoryCountsByUser(gctx *gin.Context) {
	//TODO: validar prmisos para poder ver todos los reportes de inventario

	userId := gctx.Keys[shared.UserIdKey()].(string)

	response, err := c.facade.GetAllByUserId(gctx, userId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *InventoryCountController) getInventoryCountById(gctx *gin.Context) {
	type param struct {
		Id string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}
	result, err := c.facade.GetById(gctx, p.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryCountController) createInventoryCount(gctx *gin.Context) {
	type params struct {
		CompanyId string `uri:"companyId" binding:"required,uuid"`
	}

	// TODO: validar permisos para crear conteos de inventario

	var p params
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeCreateInventoryCount
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	result, err := c.facade.Create(gctx, p.CompanyId, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)

}

func (c *InventoryCountController) updateInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeCreateInventoryCount
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	result, err := c.facade.Update(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryCountController) newAsignedForInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeChangeInventoryCountAssigned
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	result, err := c.facade.NewAssigned(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryCountController) newDateForInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeChangeInventoryCountDate
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	result, err := c.facade.NewDate(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryCountController) cancelInventoryCount(gctx *gin.Context) {
	// TODO: validar permisos

	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	err := c.facade.Cancel(gctx, p.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, true)
}

func (c *InventoryCountController) startInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	err := c.facade.StartInventoryCount(gctx, p.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusNoContent, nil)
}

func (c *InventoryCountController) commitInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeInventoryCount
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	err := c.facade.Commit(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, true)
}

func (c *InventoryCountController) draftInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeInventoryCount
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	err := c.facade.SaveDraft(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusNoContent, nil)
}

func (c *InventoryCountController) rejectInventoryCount(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	result, err := c.facade.RejectInventoryCount(gctx, p.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, result)
}

func (c *InventoryCountController) saveIncidence(gctx *gin.Context) {
	type param struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var p param
	if err := gctx.ShouldBindUri(&p); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.RecipeInventoryCountIncidence
	if err := shared.BindJSON(gctx, &body); err != nil {
		gctx.Error(err)
		return
	}

	// impl
	err := c.facade.SaveIncidence(gctx, p.ID, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Incidencia guardada correctamente",
	})
}
