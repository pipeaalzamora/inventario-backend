package controllers

import (
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type DeliveryPurchaseNoteController struct {
	deliveryPurchaseNoteFacade *facades.DeliveryPurchaseNoteFacade
}

func NewDeliveryPurchaseNoteController(deliveryPurchaseNoteFacade *facades.DeliveryPurchaseNoteFacade) *DeliveryPurchaseNoteController {
	return &DeliveryPurchaseNoteController{
		deliveryPurchaseNoteFacade: deliveryPurchaseNoteFacade,
	}
}

func (c *DeliveryPurchaseNoteController) RegisterRoutes(rg *gin.RouterGroup) {
	// r := rg.Group("/delivery-purchase-notes")

	// r.GET("/by-store/:storeId", c.getAllDeliveryPurchaseNotes)
	// r.DELETE("/:noteId/file/:fileId", c.removeFileFromDeliveryPurchaseNote)
	// r.POST("/:id/file", c.uploadFileToDeliveryPurchaseNote)
	// r.POST("/:id/complete", c.completeDeliveryPurchaseNote)
	// r.POST("/:id/fix", c.fixDeliveryPurchaseNote)
	// r.GET("/:id", c.getDeliveryPurchaseNoteByID)
	// r.PUT("/:id", c.updateDeliveryPurchaseNote)
	// r.POST("", c.createDeliveryPurchaseNote)
}

/*
func (c *DeliveryPurchaseNoteController) getAllDeliveryPurchaseNotes(gctx *gin.Context) {
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

	notes, err := c.deliveryPurchaseNoteFacade.GetAllDeliveryPurchaseNotes(gctx, params.StoreId, page, size, nil)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, notes)
}

func (c *DeliveryPurchaseNoteController) getDeliveryPurchaseNoteByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	note, err := c.deliveryPurchaseNoteFacade.GetDeliveryPurchaseNoteByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(200, note)
}

func (c *DeliveryPurchaseNoteController) createDeliveryPurchaseNote(gctx *gin.Context) {
	var request *recipe.RecipeDeliveryPurchaseNote
	if err := shared.BindJSON(gctx, &request); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.CreateDeliveryPurchaseNote(gctx, request)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, dto)
}

func (c *DeliveryPurchaseNoteController) updateDeliveryPurchaseNote(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var request *recipe.RecipeDeliveryPurchaseNote
	if err := gctx.ShouldBindJSON(&request); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.UpdateDeliveryPurchaseNote(gctx, params.ID, request)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, dto)
}

func (c *DeliveryPurchaseNoteController) uploadFileToDeliveryPurchaseNote(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var form recipe.UploadForm
	if err := gctx.ShouldBind(&form); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.UploadFileToDeliveryPurchaseNote(gctx, params.ID, &form)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, dto)
}

func (c *DeliveryPurchaseNoteController) removeFileFromDeliveryPurchaseNote(gctx *gin.Context) {
	type pathParams struct {
		NoteID string `uri:"noteId" binding:"required,uuid"`
		FileID string `uri:"fileId" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.RemoveFileFromDeliveryPurchaseNote(gctx, params.NoteID, params.FileID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, dto)
}

func (c *DeliveryPurchaseNoteController) completeDeliveryPurchaseNote(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	type invoices struct {
		InvoiceFolio string `json:"invoiceFolio"`
		InvoiceGuide string `json:"invoiceGuide"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var invoiceData invoices
	if err := gctx.ShouldBindJSON(&invoiceData); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.ConfirmDeliveryPurchaseNote(gctx, params.ID, invoiceData.InvoiceFolio, invoiceData.InvoiceGuide)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, dto)
}

func (c *DeliveryPurchaseNoteController) fixDeliveryPurchaseNote(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required,uuid"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	dto, err := c.deliveryPurchaseNoteFacade.GeneratePurchaseCorrectionNote(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, dto)
}
*/
