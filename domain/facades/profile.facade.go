package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/types"
)

type ProfileFacade struct {
	appService *services.ServiceContainer
}

func NewProfileFacade(appService *services.ServiceContainer) *ProfileFacade {
	return &ProfileFacade{
		appService: appService,
	}
}

func (f *ProfileFacade) GetProfiles(ctx context.Context) ([]dto.ProfileAccountWithPowersDTO, error) {
	profiles, err := f.appService.ProfileService.GetProfiles(ctx)
	if err != nil {
		return nil, err
	}

	power, err := f.appService.ProfileService.GetAllPowersWithProfileId(ctx)
	if err != nil {
		return nil, err
	}

	// performance improvement: map powers by profile ID
	powerMap := make(map[string][]models.PowerAccountModel)
	for _, p := range power {
		powerModel := models.PowerAccountModel{
			ID:          p.ID,
			PowerName:   p.PowerName,
			DisplayName: p.DisplayName,
			Description: p.Description,
			CategoryID:  p.CategoryID,
		}
		powerMap[p.ProfileID] = append(powerMap[p.ProfileID], powerModel)
	}

	result := make([]dto.ProfileAccountWithPowersDTO, len(profiles))
	for i, profile := range profiles {
		if powerMap[profile.ID] == nil {
			powerMap[profile.ID] = []models.PowerAccountModel{}
		}
		result[i] = dto.ProfileAccountWithPowersDTO{
			ID:          profile.ID,
			ProfileName: profile.ProfileName,
			Description: profile.Description,
			Powers:      powerMap[profile.ID],
		}
	}

	return result, nil
}

func (f *ProfileFacade) GetPowers(ctx context.Context) ([]dto.CategoryAccountWithPowersDTO, error) {
	powers, err := f.appService.ProfileService.GetPowers(ctx)
	if err != nil {
		return nil, err
	}

	categories, err := f.appService.ProfileService.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	// create a map grouping powers by category ID
	powerMap := make(map[string][]dto.PowerAccountDTO)
	for _, power := range powers {
		powerMap[power.CategoryID] = append(powerMap[power.CategoryID], dto.PowerAccountDTO{
			ID:          power.ID,
			PowerName:   power.PowerName,
			DisplayName: power.DisplayName,
			Description: power.Description,
		})
	}

	result := make([]dto.CategoryAccountWithPowersDTO, len(categories))
	for i, category := range categories {
		result[i] = dto.CategoryAccountWithPowersDTO{
			ID:           category.ID,
			CategoryName: category.CategoryName,
			Description:  category.Description,
			Ownable:      category.Ownable,
			Powers:       powerMap[category.ID],
		}
	}

	return result, nil
}

func (f *ProfileFacade) GetProfileByID(ctx context.Context, id string) (*dto.ProfileAccountWithPowersDTO, error) {
	// Retrive profile by ID with associated powers
	profile, err := f.appService.ProfileService.GetProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Change to get powers from profile Id
	powers, err := f.appService.ProfileService.GetPowersByProfileID(ctx, id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, types.ThrowMsg("Perfil no encontrado.")
	}

	dto := &dto.ProfileAccountWithPowersDTO{
		ID:          profile.ID,
		ProfileName: profile.ProfileName,
		Description: profile.Description,
		Powers:      powers,
	}

	return dto, nil
}

func (f *ProfileFacade) CreateProfile(ctx context.Context, body *recipe.ProfileRecipe) (*models.ProfileAccountModel, error) {
	profile, err := f.appService.ProfileService.CreateProfile(ctx, body)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		return nil, types.ThrowMsg("No se pudo crear el perfil.")
	}

	return profile, nil
}

func (f *ProfileFacade) UpdateProfile(ctx context.Context, id string, body *recipe.ProfileRecipe) (*models.ProfileAccountModel, error) {
	profile, err := f.appService.ProfileService.UpdateProfile(ctx, id, body)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		return nil, types.ThrowMsg("No se pudo actualizar el perfil.")
	}

	return profile, nil
}

func (f *ProfileFacade) DeleteProfile(ctx context.Context, id string) error {
	err := f.appService.ProfileService.DeleteProfile(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
