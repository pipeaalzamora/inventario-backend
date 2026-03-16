package controllers

import (
	"fmt"
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type RequestController struct {
	requestFacade *facades.RequestFacade
}

func NewRequestController(requestFacade *facades.RequestFacade) *RequestController {
	return &RequestController{
		requestFacade: requestFacade,
	}
}

func (c *RequestController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("request")

	//r.GET(":id", c.getRequestById)
	r.POST("", c.createRequest)
}

func (c *RequestController) createRequest(gctx *gin.Context) {
	var request recipe.RecipeNewRequest
	if err := shared.BindJSON(gctx, &request); err != nil {
		gctx.Error(err)
		return
	}

	newRequest, err := c.requestFacade.CreateRequest(gctx, &request)
	if err != nil {
		gctx.Error(err)
		fmt.Println("Error creating request:", err)
		return
	}

	gctx.JSON(http.StatusCreated, newRequest)
}

// func (c *RequestController) getRequestById(gctx *gin.Context) {
// 	type pathParams struct {
// 		Id string `uri:"id" binding:"required"`
// 	}

// 	var params pathParams
// 	if err := gctx.ShouldBindUri(&params); err != nil {
// 		gctx.Error(err)
// 		return
// 	}

// 	request, err := c.requestFacade.GetRequestById(gctx, params.Id)
// 	if err != nil {
// 		gctx.Error(err)
// 		return
// 	}

// 	gctx.JSON(http.StatusOK, request)
// }
