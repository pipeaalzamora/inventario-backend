package services

import (
	"context"
	"fmt"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
)

type ProfileService struct {
	PowerChecker
	profileRepo  ports.PortProfile
	powerRepo    ports.PortPower
	cacheService ports.PortCache
}

func NewProfileService(profileRepo ports.PortProfile, powerRepo ports.PortPower, cacheService ports.PortCache) *ProfileService {
	return &ProfileService{
		profileRepo:  profileRepo,
		powerRepo:    powerRepo,
		cacheService: cacheService,
	}
}

func (s *ProfileService) GetProfiles(ctx context.Context) ([]models.ProfileAccountModel, error) {
	profiles, err := s.profileRepo.GetProfiles()
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (s *ProfileService) GetPowersByUserID(ctx context.Context, userId string) ([]models.PowerAccountModel, error) {
	powers, err := s.powerRepo.GetPowersByUserID(userId)
	if err != nil {
		return nil, err
	}

	return powers, nil
}

func (s *ProfileService) GetProfilesByUserID(ctx context.Context, userId string) ([]models.ProfileAccountModel, error) {
	profiles, err := s.profileRepo.GetProfilesByUserID(userId)

	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (s *ProfileService) GetPowerCategories(ctx context.Context) ([]models.PowerAccountCategoryModel, error) {
	categories, err := s.powerRepo.GetPowerCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *ProfileService) GetPowerCategoryByID(ctx context.Context, categoryId string) ([]models.PowerAccountModel, error) {
	powers, err := s.powerRepo.GetPowerAccountsByCategory(categoryId)
	if err != nil {
		return nil, err
	}

	return powers, nil
}

func (s *ProfileService) CreateProfile(ctx context.Context, input *recipe.ProfileRecipe) (*models.ProfileAccountModel, error) {
	if ok := s.EveryPower(ctx, PowerProfileCreate); !ok {
		return nil, shared.PowerError{
			Message: fmt.Sprintf("User does not have %s power", PowerProfileCreate),
		}
	}
	inputPower := &models.ProfileAccountModel{
		ProfileName: input.ProfileName,
		Description: input.Description,
	}

	profile, err := s.profileRepo.CreateProfileWithPowers(inputPower, input.PowerIDs)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profileId string, body *recipe.ProfileRecipe) (*models.ProfileAccountModel, error) {
	if ok := s.EveryPower(ctx, PowerProfileUpdate); !ok {
		return nil, shared.PowerError{
			Message: fmt.Sprintf("User does not have %s power", PowerProfileUpdate),
		}
	}
	inputProfile := &models.ProfileAccountModel{
		ProfileName: body.ProfileName,
		Description: body.Description,
	}

	profile, err := s.profileRepo.UpdateProfileWithPowers(profileId, inputProfile, body.PowerIDs)
	if err != nil {
		return nil, err
	}

	// delete old profiles in redis
	if err := s.cacheService.DeleteByPrefix("POWERS:"); err != nil {
		return nil, err
	}

	return profile, nil
}

// func (s *ProfileService) CreatePowerCategory(ctx context.Context, input *models.InputPowerAccountCategoryModel) (*models.PowerAccountCategoryModel, error) {
// 	entityCategory := &models.InputPowerAccountCategoryModel{
// 		CategoryName: input.CategoryName,
// 		Description:  input.Description,
// 		Ownable:      input.Ownable,
// 	}
// 	category, err := s.powerRepo.CreatePowerCategory(entityCategory)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return category, nil
// }

// func (s *ProfileService) CreatePowerAccount(ctx context.Context, input *models.InputPowerModel) (*models.PowerAccountModel, error) {
// 	entityPower := &models.InputPowerModel{
// 		PowerName:   input.PowerName,
// 		Description: input.Description,
// 		CategoryID:  input.CategoryID,
// 	}
// 	power, err := s.powerRepo.CreatePowerAccount(entityPower)
// 	if err != nil {
// 		return nil, types.ThrowData("Error al crear el permiso")
// 	}
// 	return power, nil
// }

func (s *ProfileService) GetProfileByID(ctx context.Context, profileId string) (*models.ProfileAccountModel, error) {
	profile, err := s.profileRepo.GetProfileByID(profileId)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, profileId string) error {
	if ok := s.EveryPower(ctx, PowerProfileDelete); !ok {
		return shared.PowerError{Message: fmt.Sprintf("User does not have %s power", PowerProfileDelete)}
	}
	err := s.profileRepo.DeleteProfile(profileId)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) GetPowers(ctx context.Context) ([]models.PowerAccountModel, error) {
	powers, err := s.powerRepo.GetPowers()
	if err != nil {
		return nil, err
	}

	return powers, nil
}

func (s *ProfileService) GetCategories(ctx context.Context) ([]models.PowerAccountCategoryModel, error) {
	categories, err := s.powerRepo.GetPowerCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *ProfileService) GetPowersByProfileID(ctx context.Context, profileId string) ([]models.PowerAccountModel, error) {
	powers, err := s.powerRepo.GetPowersByProfile(profileId)
	if err != nil {
		return nil, err
	}
	return powers, nil
}

func (s *ProfileService) GetAllProfilesWithUserId(ctx context.Context) ([]models.ProfileAccountUserAccountModel, error) {
	profiles, err := s.profileRepo.GetAllProfilesWithUserId()
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (s *ProfileService) GetAllPowersWithProfileId(ctx context.Context) ([]models.PowerAccountProfileModel, error) {
	powers, err := s.powerRepo.GetAllPowersWithProfileId()
	if err != nil {
		return nil, err
	}

	return powers, nil
}

func (s *ProfileService) GetProfilesByPowerID(ctx context.Context, powerID string) ([]models.ProfileAccountModel, error) {
	return s.profileRepo.GetProfilesByPowerID(powerID)
}
