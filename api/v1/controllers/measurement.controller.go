package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type MeasurementController struct {
	measurementFacade *facades.MeasurementFacade
}

func NewMeasurementController(measurementFacade *facades.MeasurementFacade) *MeasurementController {
	return &MeasurementController{
		measurementFacade: measurementFacade,
	}
}

func (c *MeasurementController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/measurements")

	r.GET("", c.getMeasurements)
	r.POST("/create", c.createMeasurement)
	r.GET(":id/related", c.getRelatedUnits)
}

func (c *MeasurementController) getMeasurements(gctx *gin.Context) {
	measurements, err := c.measurementFacade.GetMeasurements()
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, measurements)
}

func (c *MeasurementController) createMeasurement(gctx *gin.Context) {
	var measurementRecipe recipe.MeasurementRecipe
	if err := shared.BindJSON(gctx, &measurementRecipe); err != nil {
		gctx.Error(err)
		return
	}

	measurement, err := c.measurementFacade.CreateMeasurement(gctx, measurementRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, measurement)
}

func (c *MeasurementController) getRelatedUnits(gctx *gin.Context) {
	type pathParams struct {
		ID int `uri:"id" binding:"required"`
	}
	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	units, err := c.measurementFacade.GetRelatedUnits(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, units)
}
