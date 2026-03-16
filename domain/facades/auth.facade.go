package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
	"time"
)

type AuthFacade struct {
	appService       *services.ServiceContainer
	externalServices *external.ServiceContainer
	secret           string
}

func NewAuthFacade(
	appService *services.ServiceContainer,
	externalServices *external.ServiceContainer,
	cfg *config.Config,
) *AuthFacade {
	return &AuthFacade{
		appService:       appService,
		externalServices: externalServices,
		secret:           cfg.Secret,
	}
}

func (f *AuthFacade) Login(ctx context.Context, body recipe.LoginRecipe) (*dto.UserWithProfilesDTO, error) {
	// Validate user credentials
	userAccountByEmail, err := f.appService.UserService.GetUserByEmail(ctx, body.UserEmail)
	if err != nil {
		return nil, err
	}

	if userAccountByEmail == nil || !userAccountByEmail.Available {
		return nil, types.ThrowMsg("Usuario o contraseña incorrectos")
	}

	if !shared.CheckPassword(body.Password, userAccountByEmail.UserPassword, f.secret) {
		return nil, types.ThrowMsg("Usuario o contraseña incorrectos")
	}

	powers, err := f.appService.ProfileService.GetPowersByUserID(ctx, userAccountByEmail.ID)
	if err != nil {
		return nil, err
	}

	err = f.appService.AuthService.PersistUserWithPowers(ctx, userAccountByEmail, powers)
	if err != nil {
		return nil, err
	}

	profiles, err := f.appService.ProfileService.GetProfilesByUserID(ctx, userAccountByEmail.ID)
	if err != nil {
		return nil, err
	}

	userDto := &dto.UserWithProfilesDTO{
		ID:           userAccountByEmail.ID,
		UserName:     userAccountByEmail.UserName,
		UserEmail:    userAccountByEmail.UserEmail,
		Description:  userAccountByEmail.Description,
		IsNewAccount: userAccountByEmail.IsNewAccount,
		Available:    userAccountByEmail.Available,
		DeletedAt:    userAccountByEmail.DeletedAt,
		CreatedAt:    userAccountByEmail.CreatedAt,
		UpdatedAt:    userAccountByEmail.UpdatedAt,
		Profiles:     profiles,
	}

	return userDto, nil
}

func (g *AuthFacade) Logout(ctx context.Context) error {
	userId, ok := g.appService.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil
	}
	err := g.appService.AuthService.DeletePersistedUser(ctx, userId)
	if err != nil {
		return err
	}

	// disconnect from SSE if needed
	g.appService.NotificationService.RemoveClient(userId)

	return nil
}

func (f *AuthFacade) RecoverPassword(ctx context.Context, body *recipe.RecoveryPasswordRecipe) (time.Time, error) {
	code, err := f.appService.AuthService.GenerateRecoveryCode(ctx, body.UserEmail)
	if err != nil {
		return time.Time{}, err
	}

	err = f.externalServices.EmailService.SendRecoveryEmail(body.UserEmail, code.Code, code.TTL)
	if err != nil {
		return time.Time{}, err
	}

	return code.TTL, nil
}

func (f *AuthFacade) VerifyRecoveryPassword(ctx context.Context, body *recipe.VerifyCodeRecipe) (time.Time, error) {
	code, err := f.appService.AuthService.GetPersistedRecoveryCode(ctx, body.UserEmail)
	if err != nil {
		return time.Time{}, err
	}

	if code == nil {
		return time.Time{}, types.ThrowMsg("No se ha encontrado un código de recuperación para el correo proporcionado.")
	}

	if code.Code != body.Code {
		return time.Time{}, types.ThrowMsg("El código de recuperación es incorrecto.")
	}

	err = f.appService.AuthService.VerifyRecoveryCode(ctx, body.UserEmail, code)
	if err != nil {
		return time.Time{}, err
	}

	return code.TTL, nil

}

func (f *AuthFacade) ChangePasswordWithRecoveryCode(ctx context.Context, body *recipe.ChangePasswordWithCodeRecipe) error {
	if body.NewPassword != body.ConfirmPassword {
		return types.ThrowMsg("La contraseña nueva y la confirmación no coinciden.")
	}

	code, err := f.appService.AuthService.GetPersistedRecoveryCode(ctx, body.UserEmail)
	if err != nil {
		return err
	}

	if code == nil || !code.Verified {
		return types.ThrowMsg("El código de recuperación no es válido o no ha sido verificado.")
	}

	// Delete the recovery code after use
	err = f.appService.AuthService.DeletePersistedRecoveryCode(ctx, body.UserEmail)
	if err != nil {
		return err
	}

	// Change the user's password using the user service
	user, err := f.appService.UserService.GetUserByEmail(ctx, body.UserEmail)
	if err != nil {
		return err
	}

	if user == nil {
		return types.ThrowMsg("Usuario no encontrado.")
	}

	// encrypt the new password
	encryptedPassword := shared.CreatePassword(body.NewPassword, f.secret)
	body.NewPassword = encryptedPassword

	userRecipe := &recipe.UserRecipe{
		UserName:     user.UserName,
		UserEmail:    user.UserEmail,
		Description:  user.Description,
		IsNewAccount: false,
		Password:     body.NewPassword,
	}

	_, err = f.appService.UserService.UpdateUserNotAuth(ctx, user.ID, userRecipe)
	if err != nil {
		return err
	}

	// If recovery code is valid, proceed to change the password
	return nil
}

func (f *AuthFacade) ChangePassword(ctx context.Context, body *recipe.ChangePasswordRecipe) error {
	userId, ok := f.appService.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil
	}
	if body.NewPassword != body.ConfirmPassword {
		return types.ThrowMsg("La contraseña nueva y la confirmación no coinciden.")
	}

	user, err := f.appService.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return types.ThrowMsg("Usuario no encontrado.")
	}

	if !shared.CheckPassword(body.OldPassword, user.UserPassword, f.secret) {
		return types.ThrowMsg("La contraseña actual es incorrecta.")
	}

	// Encrypt the new password
	encryptedPassword := shared.CreatePassword(body.NewPassword, f.secret)
	user.UserPassword = encryptedPassword
	user.IsNewAccount = false

	// Save the updated user
	userInput := &recipe.UserRecipe{
		UserName:     user.UserName,
		UserEmail:    user.UserEmail,
		Description:  user.Description,
		IsNewAccount: false,
		Password:     user.UserPassword,
	}

	_, err = f.appService.UserService.UpdateUser(ctx, user.ID, userInput)
	if err != nil {
		return err
	}

	return nil
}

func (f *AuthFacade) GetPersistedUser(ctx context.Context) (*models.UserAccountModel, []models.PowerAccountModel, error) {
	userId, ok := f.appService.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, nil, types.ThrowMsg("Usuario no encontrado.")
	}

	user, err := f.appService.AuthService.GetPersistedUser(ctx, userId)
	if err != nil {
		return nil, nil, err
	}

	powers, err := f.appService.AuthService.GetUserPersistedPowers(ctx, userId)
	if err != nil {
		powers = []models.PowerAccountModel{}
	}

	if user == nil || user.ID == "" { // Verifica que el puntero no sea nil y que tenga datos
		dbUser, err := f.appService.UserService.GetUserByID(ctx, userId)
		if err != nil {
			return nil, nil, err
		}

		user = dbUser

		if err := f.appService.AuthService.PersistUser(ctx, user); err != nil {
			return nil, nil, err
		}
	}

	if len(powers) == 0 { // Evita nil y slice vacío
		dbPowers, err := f.appService.ProfileService.GetPowersByUserID(ctx, userId)
		if err != nil {
			return nil, nil, err
		}

		powers = dbPowers

		if err := f.appService.AuthService.PersistPowers(ctx, userId, powers); err != nil {
			return nil, nil, err
		}
	}

	return user, powers, nil
}

func (f *AuthFacade) GetUser(ctx context.Context, userId string) (*models.UserAccountModel, error) {
	user, err := f.appService.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, types.ThrowMsg("Usuario no encontrado.")
	}

	return user, nil
}
