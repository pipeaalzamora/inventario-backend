package controllers

import (
	"net/http"
	apiservices "sofia-backend/api/v1/api-services"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/config"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authFacade *facades.AuthFacade
	jwtSecret  string
	isDebug    bool
}

func NewAuthController(authFacade *facades.AuthFacade, config *config.Config) *AuthController {
	return &AuthController{authFacade: authFacade, jwtSecret: config.JwtSecret, isDebug: config.Debug}
}

func (ctrl *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/auth")

	r.POST("/login", ctrl.login)
	r.POST("/recovery", ctrl.recoveryPassword)
	r.POST("/verify-code", ctrl.verifyCode)
	r.POST("/change-password/recovery", ctrl.changePasswordWithCode)

	// Rutas protegidas
	protected := r.Group("")
	protected.Use(apiservices.AuthMiddleware(ctrl.authFacade, ctrl.jwtSecret))
	protected.POST("/change-password", ctrl.changePassword)
	protected.GET("/powers", ctrl.getPowers)
	protected.POST("/logout", ctrl.logout)
}

func (ctrl *AuthController) login(gctx *gin.Context) {
	// Implementación del login

	// 1. Validar el cuerpo de la solicitud
	var body recipe.LoginRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// 2. Llamar al facade para autenticar al usuario
	user, err := ctrl.authFacade.Login(gctx, body)
	if err != nil {
		gctx.Error(err)
		return
	}

	// 3. Generar un token JWT
	ttl := 1
	if ctrl.isDebug {
		ttl = 24
	}

	token, err := shared.GenerateToken(user.ID, ctrl.jwtSecret, ttl)
	if err != nil {
		gctx.Error(err)
		return
	}

	// 4. Devolver la respuesta
	gctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (ctrl *AuthController) recoveryPassword(gctx *gin.Context) {
	var body recipe.RecoveryPasswordRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}
	ttl, err := ctrl.authFacade.RecoverPassword(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"ttl": ttl})
}

func (ctrl *AuthController) verifyCode(gctx *gin.Context) {
	var body recipe.VerifyCodeRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	// 2. Llamar al facade para verificar el código
	ttl, err := ctrl.authFacade.VerifyRecoveryPassword(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"ttl": ttl})
}

func (ctrl *AuthController) changePasswordWithCode(gctx *gin.Context) {
	var body recipe.ChangePasswordWithCodeRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	err := ctrl.authFacade.ChangePasswordWithRecoveryCode(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusNoContent)
}

func (ctrl *AuthController) logout(gctx *gin.Context) {
	err := ctrl.authFacade.Logout(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusNoContent) // 204 No Content
}

func (ctrl *AuthController) changePassword(gctx *gin.Context) {
	var body recipe.ChangePasswordRecipe
	if err := gctx.ShouldBindJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	err := ctrl.authFacade.ChangePassword(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusNoContent)
}

func (ctrl *AuthController) getPowers(gctx *gin.Context) {
	_, powers, err := ctrl.authFacade.GetPersistedUser(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"powers": powers})
}
