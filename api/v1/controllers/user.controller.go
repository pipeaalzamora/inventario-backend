package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userFacade *facades.UserFacade
}

func NewUserController(userFacade *facades.UserFacade) *UserController {
	return &UserController{
		userFacade: userFacade,
	}
}

func (c *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users-accounts")

	r.POST("", c.createUser)
	r.PUT(":id", c.updateUser)
	r.GET("", c.getUsers)
	r.GET(":id", c.getUserById)
	r.DELETE(":id", c.deleteUser)
	r.PATCH(":id/activate", c.activateUser)
}

func (c *UserController) createUser(gctx *gin.Context) {
	var recipe recipe.UserRecipe
	if err := gctx.ShouldBindJSON(&recipe); err != nil {
		gctx.Error(err)
		return
	}

	user, err := c.userFacade.CreateUser(gctx, &recipe)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusCreated, user)
}

func (c *UserController) updateUser(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.UserRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	user, err := c.userFacade.UpdateUser(gctx, params.ID, &body)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, user)
}

func (c *UserController) getUsers(gctx *gin.Context) {
	users, err := c.userFacade.GetAllUsers(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *UserController) getUserById(gctx *gin.Context) {
	type queryParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params queryParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	user, err := c.userFacade.GetUserById(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, user)
}

func (c *UserController) deleteUser(gctx *gin.Context) {
	type queryParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params queryParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	if err := c.userFacade.DeactivateUser(gctx, params.ID); err != nil {
		gctx.Error(err)
		return
	}

	userDeleted, err := c.userFacade.GetUserById(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, userDeleted)
}

func (c *UserController) activateUser(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	if err := c.userFacade.ActivateUser(gctx, params.ID); err != nil {
		gctx.Error(err)
		return
	}

	userActivated, err := c.userFacade.GetUserById(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, userActivated)
}
