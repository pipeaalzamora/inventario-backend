package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"

	"github.com/gin-gonic/gin"
)

type ProfileAccountController struct {
	profileFacade *facades.ProfileFacade
}

func NewProfileAccountController(profileFacade *facades.ProfileFacade) *ProfileAccountController {
	return &ProfileAccountController{profileFacade: profileFacade}
}

func (ctrl *ProfileAccountController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/profile-accounts")

	r.GET("", ctrl.getProfiles)
	r.GET("/powers", ctrl.getPowers)
	r.GET("/:id", ctrl.getProfileByID)
	r.POST("", ctrl.createProfile)
	r.PUT("/:id", ctrl.updateProfile)
	r.DELETE("/:id", ctrl.deleteProfile)
}

func (ctrl *ProfileAccountController) getProfiles(gctx *gin.Context) {
	profiles, err := ctrl.profileFacade.GetProfiles(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

func (ctrl *ProfileAccountController) getPowers(gctx *gin.Context) {
	powers, err := ctrl.profileFacade.GetPowers(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"categories": powers})
}

func (ctrl *ProfileAccountController) getProfileByID(gctx *gin.Context) {
	type queryParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params queryParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}
	profile, err := ctrl.profileFacade.GetProfileByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, profile)
}

func (ctrl *ProfileAccountController) createProfile(gctx *gin.Context) {
	var body recipe.ProfileRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	profile, err := ctrl.profileFacade.CreateProfile(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusCreated, profile)
}

func (ctrl *ProfileAccountController) updateProfile(gctx *gin.Context) {
	type queryParams struct {
		ID string `uri:"id" binding:"required"`
	}
	var params queryParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var body recipe.ProfileRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	profile, err := ctrl.profileFacade.UpdateProfile(gctx, params.ID, &body)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, profile)
}

func (ctrl *ProfileAccountController) deleteProfile(gctx *gin.Context) {
	type queryParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params queryParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	err := ctrl.profileFacade.DeleteProfile(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.Status(http.StatusNoContent) // 204 No Content
}
