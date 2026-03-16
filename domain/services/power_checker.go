package services

import (
	"context"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/shared"
	"strings"

	"github.com/gin-gonic/gin"
)

type PowerChecker struct{}

func (pc PowerChecker) GetPowersFromContext(ctx context.Context) ([]string, bool) {
	gctx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, false
	}
	v, ok := gctx.Get(shared.UserPowersKeys())
	if !ok {
		return nil, false
	}
	powers, ok := v.([]string)
	if !ok {
		return nil, false
	}
	return powers, true
}

func (pc PowerChecker) GetUserIDFromContext(ctx context.Context) (string, bool) {
	gctx, ok := ctx.(*gin.Context)
	if !ok {
		return "", false
	}

	v, ok := gctx.Get(shared.UserIdKey())
	if !ok {
		return "", false
	}
	userId, ok := v.(string)
	if !ok {
		return "", false
	}
	return userId, true
}

func (pc PowerChecker) GetUserFromContext(ctx context.Context) (*models.UserAccountModel, bool) {
	gctx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, false
	}
	v, ok := gctx.Get(shared.UserKey())
	if !ok {
		return nil, false
	}

	// Asegurarse de que el tipo es correcto
	user, ok := v.(*models.UserAccountModel)
	if !ok {
		return nil, false
	}

	if user == nil {
		return nil, false
	}

	return user, true
}

// Devuelve true si el usuario tiene al menos uno de los permisos solicitados
func (pc PowerChecker) SomePower(ctx context.Context, powers ...string) bool {
	userPowers, ok := pc.GetPowersFromContext(ctx)
	if !ok {
		return false
	}

	for _, required := range powers {
		for _, userPower := range userPowers {
			if cmpPowers(required, userPower) {
				return true
			}
		}
	}
	return false
}

// Devuelve true si el usuario tiene **todos** los permisos requeridos
func (pc PowerChecker) EveryPower(ctx context.Context, powers ...string) bool {
	userPowers, ok := pc.GetPowersFromContext(ctx)
	if !ok {
		fmt.Println("No powers found in context")
		return false
	}

	for _, required := range powers {
		found := false
		for _, userPower := range userPowers {
			if cmpPowers(required, userPower) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (pc PowerChecker) EvalFunc(ctx context.Context, f func(...string) bool) bool {
	cuserPowers, ok := pc.GetPowersFromContext(ctx)
	if !ok {
		fmt.Println("No powers found in context")
		return false
	}
	return f(cuserPowers...)
}

func cmpPowers(a, b string) bool {
	aCmp := strings.ToLower(strings.TrimSpace(a))
	if aCmp == "" {
		return false
	}
	bCmp := strings.ToLower(strings.TrimSpace(b))
	if bCmp == "" {
		return false
	}
	return aCmp == bCmp
}
