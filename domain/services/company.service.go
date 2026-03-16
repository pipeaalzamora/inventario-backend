package services

import (
	"context"
	"mime/multipart"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
	"strings"
)

type CompanyService struct {
	PowerChecker
	repo         ports.PortCompany
	profileRepo  ports.PortProfile
	cacheService ports.PortCache
	userRepo     ports.PortUser
	bucket       ports.PortBucket
}

func NewCompanyService(companyRepo ports.PortCompany,
	profileRepo ports.PortProfile,
	bucket ports.PortBucket,
	cacheService ports.PortCache,
	userRepo ports.PortUser,
) *CompanyService {
	return &CompanyService{
		repo:         companyRepo,
		profileRepo:  profileRepo,
		bucket:       bucket,
		userRepo:     userRepo,
		cacheService: cacheService,
	}
}

func (cs *CompanyService) GetCompanies(ctx context.Context) ([]models.ModelCompany, error) {
	// get ownable powers
	powers, ok := cs.GetPowersFromContext(ctx)
	if !ok {
		return nil, nil
	}

	// if the user has the "company:" prefix power, return all companies
	filteredCompanyID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixCompany) {
			companyId := p[len(PowerPrefixCompany):]
			filteredCompanyID[companyId] = true
		}
	}

	allCompanies, err := cs.repo.GetCompanies()
	if err != nil {
		return nil, err
	}

	result := []models.ModelCompany{}
	for _, company := range allCompanies {
		if filteredCompanyID[company.ID] {
			result = append(result, company)
		}
	}

	return result, nil
}

func (cs *CompanyService) GetCompanyByID(ctx context.Context, id string) (*models.ModelCompany, error) {
	if ok := cs.EveryPower(ctx, PowerPrefixCompany+id); !ok {
		return nil, types.ThrowPower("Insufficient powers to access this company")
	}
	return cs.repo.GetCompanyByID(id)
}

func (cs *CompanyService) GetCompanyByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelCompany, error) {
	return cs.repo.GetCompanyByFiscalIDAndCountry(fiscalID, countryID)
}

func (cs *CompanyService) GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelCompany, error) {
	return cs.repo.GetCompanyByFiscalNameAndCountry(fiscalName, countryID)
}

func (cs *CompanyService) GetCompanySuppliers(ctx context.Context, companyID string) ([]models.CompanySupplierModel, error) {
	if ok := cs.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta compañía")
	}

	if _, err := cs.repo.GetCompanyByID(companyID); err != nil {
		return nil, err
	}

	return cs.repo.GetSuppliersByCompanyID(companyID)
}

func (cs *CompanyService) AssignSuppliersToCompany(ctx context.Context, companyID string, recipe *recipe.AssignSuppliersToCompanyRecipe) ([]models.CompanySupplierModel, error) {
	if ok := cs.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta compañía")
	}

	if len(recipe.SupplierIds) == 0 {
		return nil, types.ThrowRecipe("Debe enviar al menos un proveedor", "supplierIds")
	}

	if _, err := cs.repo.GetCompanyByID(companyID); err != nil {
		return nil, types.ThrowMsg("No se encontró la empresa")
	}

	if err := cs.repo.AssignSuppliersToCompany(companyID, recipe.SupplierIds); err != nil {
		return nil, err
	}

	return cs.repo.GetSuppliersByCompanyID(companyID)
}

func (cs *CompanyService) UnassignSupplierFromCompany(ctx context.Context, companyID, supplierID string) ([]models.CompanySupplierModel, error) {
	if ok := cs.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta compañía")
	}

	if _, err := cs.repo.GetCompanyByID(companyID); err != nil {
		return nil, err
	}

	if err := cs.repo.UnassignSupplierFromCompany(companyID, supplierID); err != nil {
		return nil, err
	}

	return cs.repo.GetSuppliersByCompanyID(companyID)
}

func (cs *CompanyService) CreateCompany(ctx context.Context, createRecipe *recipe.RecipeCreateCompany) (*models.ModelCompany, error) {
	if ok := cs.SomePower(ctx, PowerPrefixCompany+"create"); !ok {
		return nil, types.ThrowMsg("No tienes los permisos para crear empresas")
	}

	if createRecipe.ImageLogo.Size < 1 {
		return nil, types.ThrowRecipe("El logo de la empresa es obligatorio", "imageLogo")
	}

	profiles, err := cs.profileRepo.GetProfilesByPowerID("company:create")
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(createRecipe.IDFiscal, "CL-") {
		createRecipe.IDFiscal = "CL-" + createRecipe.IDFiscal
	}

	rawFiscalID := strings.ReplaceAll(createRecipe.IDFiscal, "CL-", "")
	rawFiscalID = strings.ReplaceAll(rawFiscalID, ".", "")

	companymodel := &models.ModelCompany{
		// CountryID: createRecipe.CountryID,
		CountryID:   1,
		CompanyName: createRecipe.CompanyName,
		Description: createRecipe.Description,
		FiscalData: models.ModelFiscalData{
			IDFiscal:      createRecipe.IDFiscal,
			RawFiscalID:   rawFiscalID,
			FiscalName:    createRecipe.FiscalName,
			FiscalAddress: createRecipe.FiscalAddress,
			FiscalState:   createRecipe.FiscalState,
			FiscalCity:    createRecipe.FiscalCity,
			Email:         createRecipe.ContactEmail,
		},
		ImageLogo: nil,
	}

	createdCompany, err := cs.repo.CreateCompany(companymodel, profiles)
	if err != nil {
		return nil, err
	}

	url, err := cs.addFileToCreatedCompany(ctx, createdCompany.ID, &createRecipe.ImageLogo)
	if err != nil {
		return nil, err
	}

	createdCompany.ImageLogo = url

	var profileIds []string
	for _, profile := range profiles {
		profileIds = append(profileIds, profile.ID)
	}

	users, err := cs.userRepo.GetUsersByProfileIDs(profileIds)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		cs.cacheService.DeleteByKey("POWERS:" + user.ID)
	}

	return createdCompany, nil
}

func (cs *CompanyService) UpdateCompany(ctx context.Context, id string, updateRecipe *recipe.RecipeCreateCompany, ogCompany *models.ModelCompany) (*models.ModelCompany, error) {
	// Validate ownership of the company
	if ok := cs.EveryPower(ctx, PowerPrefixCompany+id); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía")
	}

	if updateRecipe.IsNewLogo && updateRecipe.ImageLogo.Size < 1 {
		return nil, types.ThrowRecipe("El logo de la empresa es obligatorio", "imageLogo")
	}

	updateModel := &models.ModelCompany{
		ID:          id,
		CountryID:   1,
		CompanyName: updateRecipe.CompanyName,
		Description: updateRecipe.Description,
		FiscalData: models.ModelFiscalData{
			ID:            ogCompany.FiscalData.ID,
			IDFiscal:      updateRecipe.IDFiscal,
			FiscalName:    updateRecipe.FiscalName,
			FiscalAddress: updateRecipe.FiscalAddress,
			FiscalState:   updateRecipe.FiscalState,
			FiscalCity:    updateRecipe.FiscalCity,
			Email:         updateRecipe.ContactEmail,
		},
		ImageLogo: ogCompany.ImageLogo,
	}

	updatedCompany, err := cs.repo.UpdateCompany(id, updateModel)
	if err != nil {
		return nil, err
	}

	if updateRecipe.IsNewLogo {
		updatedCompany.ImageLogo, err = cs.updateFileOfCompany(ctx, id, &updateRecipe.ImageLogo, ogCompany.ImageLogo)
		if err != nil {
			return nil, err
		}
	}

	return updatedCompany, nil
}

func (cs *CompanyService) addFileToCreatedCompany(ctx context.Context, companyID string, file *multipart.FileHeader) (*string, error) {
	// Open the file to get its content
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Generate a unique name for the file
	url, err := cs.bucket.UploadFile(ctx, f, file.Filename)
	if err != nil {
		return nil, err
	}

	err = cs.repo.AddLogoToCompany(companyID, url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (cs *CompanyService) updateFileOfCompany(ctx context.Context, companyID string, file *multipart.FileHeader, oldUrl *string) (*string, error) {

	err := cs.bucket.DeleteFile(ctx, *oldUrl)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar el logo de la empresa")
	}

	return cs.addFileToCreatedCompany(ctx, companyID, file)

}
