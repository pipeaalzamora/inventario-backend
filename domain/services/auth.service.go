package services

import (
	"context"
	"fmt"
	"sofia-backend/config"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strings"
	"time"
)

type AuthService struct {
	PowerChecker
	cacheService ports.PortCache
	isDebug      bool
}

func NewAuthService(cacheService ports.PortCache, cfg *config.Config) *AuthService {
	return &AuthService{
		cacheService: cacheService,
		isDebug:      cfg.Debug,
	}
}

func (s *AuthService) GetPersistedUser(ctx context.Context, userId string) (*models.UserAccountModel, error) {
	user := models.UserAccountModel{}
	err := s.cacheService.GetWithKey(fmt.Sprintf("USERS:%s", userId), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) GetUserPersistedPowers(ctx context.Context, userId string) ([]models.PowerAccountModel, error) {
	powers := make([]models.PowerAccountModel, 0)
	err := s.cacheService.GetWithKey(fmt.Sprintf("POWERS:%s", userId), &powers)
	if err != nil {
		return nil, err
	}

	return powers, nil
}

func (s *AuthService) GetPersistedUserWithPowers(ctx context.Context, userId string) (*models.UserAccountModel, []models.PowerAccountModel, error) {
	user := models.UserAccountModel{}
	err := s.cacheService.GetWithKey(fmt.Sprintf("USERS:%s", userId), &user)
	if err != nil {
		return nil, nil, err
	}

	powers := []models.PowerAccountModel{}
	err = s.cacheService.GetWithKey(fmt.Sprintf("POWERS:%s", userId), &powers)
	if err != nil {
		return nil, nil, err
	}

	return &user, powers, nil
}

func (s *AuthService) PersistUserWithPowers(ctx context.Context, user *models.UserAccountModel, powers []models.PowerAccountModel) error {
	if user == nil {
		return types.ThrowMsg("Usuario no encontrado")
	}
	err := s.cacheService.AddKeyValue(fmt.Sprintf("USERS:%s", user.ID), *user, time.Hour*24)
	if err != nil {
		return err
	}

	err = s.cacheService.AddKeyValue(fmt.Sprintf("POWERS:%s", user.ID), powers, time.Hour*24)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) PersistUser(ctx context.Context, user *models.UserAccountModel) error {
	if user == nil {
		return types.ThrowMsg("Usuario no encontrado")
	}
	err := s.cacheService.AddKeyValue(fmt.Sprintf("USERS:%s", user.ID), *user, time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) PersistPowers(ctx context.Context, userId string, powers []models.PowerAccountModel) error {
	if powers == nil {
		return types.ThrowMsg("Permisos no encontrados")
	}

	err := s.cacheService.AddKeyValue(fmt.Sprintf("POWERS:%s", userId), powers, time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) DeletePersistedUser(ctx context.Context, userId string) error {
	err := s.cacheService.DeleteByKey(fmt.Sprintf("USERS:%s", userId))
	if err != nil {
		return err
	}

	err = s.cacheService.DeleteByKey(fmt.Sprintf("POWERS:%s", userId))
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetPersistedRecoveryCode(ctx context.Context, userEmail string) (*models.RecoveryCode, error) {
	code := &models.RecoveryCode{}
	err := s.cacheService.GetWithKey(fmt.Sprintf("RECOVERYCODE:%s", userEmail), code)
	if err != nil {
		return nil, err
	}
	return code, nil
}

func (s *AuthService) VerifyRecoveryCode(ctx context.Context, userEmail string, code *models.RecoveryCode) error {
	if code == nil {
		return types.ThrowMsg("Código de recuperación no encontrado")
	}
	code.Verified = true
	err := s.cacheService.AddKeyValue(fmt.Sprintf("RECOVERYCODE:%s", userEmail), code, s.defaultRecoveryCodeTTL())
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateRecoveryCode(ctx context.Context, userEmail string) (*models.RecoveryCode, error) {
	codeGenerated, err := shared.CreateRandomString(10)
	if err != nil {
		return nil, err
	}

	if s.isDebug {
		codeGenerated = "DEBUGCODE"
	}

	code := &models.RecoveryCode{
		Code:     strings.ToUpper(codeGenerated),
		TTL:      time.Now().Add(s.defaultRecoveryCodeTTL()),
		Verified: false,
	}

	err = s.cacheService.AddKeyValue(fmt.Sprintf("RECOVERYCODE:%s", userEmail), code, s.defaultRecoveryCodeTTL())
	if err != nil {
		return nil, types.ThrowData("Error al persistir el código de recuperación")
	}
	return code, nil
}

func (s *AuthService) DeletePersistedRecoveryCode(ctx context.Context, userEmail string) error {
	err := s.cacheService.DeleteByKey(fmt.Sprintf("RECOVERYCODE:%s", userEmail))
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) defaultRecoveryCodeTTL() time.Duration {
	return 5 * time.Minute
}
