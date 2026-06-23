package services

import (
	"net/http/httptest"
	"sofia-backend/domain/models"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

func serviceTestContext(powers ...string) *gin.Context {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	user := &models.UserAccountModel{
		ID:        "user-1",
		UserName:  "Usuario Test",
		UserEmail: "test@example.com",
	}

	ctx.Set(shared.UserIdKey(), user.ID)
	ctx.Set(shared.UserKey(), user)
	ctx.Set(shared.UserPowersKeys(), powers)

	return ctx
}
