package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strings"
)

type CompanyFacade struct {
	appServices *services.ServiceContainer
}

func NewCompanyFacade(appServices *services.ServiceContainer) *CompanyFacade {
	return &CompanyFacade{
		appServices: appServices,
	}
}

func (f *CompanyFacade) GetCompanies(ctx context.Context) ([]dto.SimpleCompanyDTO, error) {
	companies, err := f.appServices.CompanyService.GetCompanies(ctx)
	if err != nil {
		return nil, err
	}

	return f.toSimpleDtoList(companies), nil
}

func (f *CompanyFacade) GetCompaniesPaginated(ctx context.Context) (shared.PaginationResponse[dto.SimpleCompanyDTO], error) {
	companies, err := f.appServices.CompanyService.GetCompanies(ctx)
	if err != nil {
		return shared.PaginationResponse[dto.SimpleCompanyDTO]{}, err
	}
	return shared.NewPagination(f.toSimpleDtoList(companies), len(companies), 1, len(companies)), nil
}

func (f *CompanyFacade) GetCompanyByID(ctx context.Context, id string) (*dto.DetailedCompanyDTO, error) {
	company, err := f.appServices.CompanyService.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return f.toDetailedDto(company), nil

}

func (f *CompanyFacade) CreateCompany(ctx context.Context, createRecipe *recipe.RecipeCreateCompany) (*models.ModelCompany, error) {

	if _, err := f.appServices.CompanyService.GetCompanyByFiscalIDAndCountry(createRecipe.IDFiscal, 1); err == nil {
		return nil, types.ThrowRecipe("Ya existe una empresa con ese rut", "idFiscal")
	}

	if _, err := f.appServices.CompanyService.GetCompanyByFiscalNameAndCountry(createRecipe.FiscalName, 1); err == nil {
		return nil, types.ThrowRecipe("Ya existe una empresa con ese nombre fiscal", "fiscalName")
	}

	// Remove the "CL-" prefix if it exists
	rawIDFiscal := strings.TrimPrefix(createRecipe.IDFiscal, "CL-")

	if !shared.IsValidRUT(rawIDFiscal) {
		return nil, types.ThrowRecipe("El RUT es inválido", "idFiscal")
	}

	return f.appServices.CompanyService.CreateCompany(ctx, createRecipe)
}

func (f *CompanyFacade) UpdateCompany(ctx context.Context, id string, recipeUpdate *recipe.RecipeCreateCompany) (*models.ModelCompany, error) {

	ogCompany, err := f.appServices.CompanyService.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if ogCompany.FiscalData.IDFiscal != recipeUpdate.IDFiscal {
		return nil, types.ThrowRecipe("No se puede cambiar el RUT de una empresa existente", "idFiscal")
	}

	return f.appServices.CompanyService.UpdateCompany(ctx, id, recipeUpdate, ogCompany)
}

func (f *CompanyFacade) GetCompanySuppliers(ctx context.Context, companyID string) ([]dto.CompanySupplierDTO, error) {
	suppliers, err := f.appServices.CompanyService.GetCompanySuppliers(ctx, companyID)
	if err != nil {
		return nil, err
	}

	return f.toSupplierDtoList(suppliers), nil
}

func (f *CompanyFacade) AssignSuppliersToCompany(ctx context.Context, companyID string, supplierRecipe *recipe.AssignSuppliersToCompanyRecipe) ([]dto.CompanySupplierDTO, error) {
	suppliers, err := f.appServices.CompanyService.AssignSuppliersToCompany(ctx, companyID, supplierRecipe)
	if err != nil {
		return nil, err
	}

	return f.toSupplierDtoList(suppliers), nil
}

func (f *CompanyFacade) UnassignSupplierFromCompany(ctx context.Context, companyID, supplierID string) ([]dto.CompanySupplierDTO, error) {
	suppliers, err := f.appServices.CompanyService.UnassignSupplierFromCompany(ctx, companyID, supplierID)
	if err != nil {
		return nil, err
	}

	return f.toSupplierDtoList(suppliers), nil
}

// ///////// Private methods ///////////

func (f *CompanyFacade) toSimpleDto(model *models.ModelCompany) *dto.SimpleCompanyDTO {
	var logo *string = nil
	if model.ImageLogo != nil {
		logo = model.ImageLogo
	}

	return &dto.SimpleCompanyDTO{
		ID:          model.ID,
		CompanyName: model.CompanyName,
		Description: model.Description,
		IDFiscal:    model.FiscalData.IDFiscal,
		CountryID:   model.CountryID,
		ImageLogo:   logo,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func (f *CompanyFacade) toSimpleDtoList(models []models.ModelCompany) []dto.SimpleCompanyDTO {
	result := []dto.SimpleCompanyDTO{}
	for _, model := range models {
		result = append(result, *f.toSimpleDto(&model))
	}
	return result
}

func (f *CompanyFacade) toDetailedDto(model *models.ModelCompany) *dto.DetailedCompanyDTO {
	var logo *string = nil
	if model.ImageLogo != nil {
		logo = model.ImageLogo
	}

	return &dto.DetailedCompanyDTO{
		ID:          model.ID,
		CountryID:   model.CountryID,
		CompanyName: model.CompanyName,
		Description: model.Description,
		ImageLogo:   logo,
		FiscalData: dto.CompanyFiscalDataDto{
			ID:            model.FiscalData.ID,
			IDFiscal:      model.FiscalData.IDFiscal,
			FiscalName:    model.FiscalData.FiscalName,
			FiscalAddress: model.FiscalData.FiscalAddress,
			FiscalState:   model.FiscalData.FiscalState,
			FiscalCity:    model.FiscalData.FiscalCity,
			Email:         model.FiscalData.Email,
		},
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (f *CompanyFacade) toSupplierDtoList(models []models.CompanySupplierModel) []dto.CompanySupplierDTO {
	result := make([]dto.CompanySupplierDTO, len(models))
	for i, sup := range models {
		result[i] = dto.CompanySupplierDTO{
			ID:            sup.SupplierID,
			SupplierName:  sup.SupplierName,
			Description:   sup.Description,
			Available:     sup.Available,
			CountryID:     sup.CountryID,
			IDFiscal:      sup.IDFiscal,
			RawFiscalID:   sup.RawFiscalID,
			FiscalName:    sup.FiscalName,
			FiscalAddress: sup.FiscalAddress,
			Email:         sup.Email,
		}
	}
	return result
}
