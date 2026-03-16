package data

import (
	"context"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strings"

	"github.com/jmoiron/sqlx"
)

type CompanyRepo struct {
	db        *sqlx.DB
	powerRepo ports.PortPower
}

func NewCompanyRepo(db *sqlx.DB, powerRepo ports.PortPower) ports.PortCompany {
	return &CompanyRepo{
		db:        db,
		powerRepo: powerRepo,
	}
}

func (r *CompanyRepo) GetCompanies() ([]models.ModelCompany, error) {
	var companies []*entities.EntityCompany
	query := `
		SELECT 
			id,
			country_id,
			fiscal_data_id,
			company_name,
			description,
			image_logo,
			created_at,
			updated_at
		FROM company
		ORDER BY created_at DESC
	`

	err := r.db.SelectContext(context.TODO(), &companies, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener las empresas")
	}

	models := make([]models.ModelCompany, len(companies))
	for i, c := range companies {
		fiscalData, err := r.getFiscalData(c.FiscalDataID)
		if err != nil {
			return nil, err
		}
		models[i] = *r.toModel(c, fiscalData)
	}

	return models, nil
}

func (r *CompanyRepo) GetCompanyByID(id string) (*models.ModelCompany, error) {

	queryCompany := `
		SELECT 
			id,
			country_id,
			fiscal_data_id,
			company_name,
			description,
			image_logo,
			created_at,
			updated_at
		FROM company
		WHERE id = $1
	`

	var company entities.EntityCompany
	err := r.db.GetContext(context.TODO(), &company, queryCompany, id)
	if err != nil {
		return nil, types.ThrowData("No se encontró la empresa con el id: " + id)
	}

	fiscalData, err := r.getFiscalData(company.FiscalDataID)
	if err != nil {
		return nil, err
	}

	return r.toModel(&company, fiscalData), nil
}

func (r *CompanyRepo) GetCompanyByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelCompany, error) {

	query := `
		SELECT
			company.id,
			company.country_id,
			company.fiscal_data_id,
			company.company_name,
			company.description,
			company.image_logo,
			company.created_at,
			company.updated_at
		FROM company
		JOIN fiscal_data fd ON company.fiscal_data_id = fd.id
		WHERE fd.id_fiscal = $1 AND company.country_id = $2
	`

	var company entities.EntityCompany
	if err := r.db.Get(&company, query, fiscalID, countryID); err != nil {
		return nil, types.ThrowData("Error al obtener la empresa por ID fiscal")
	}

	fiscalData, err := r.getFiscalData(company.FiscalDataID)
	if err != nil {
		return nil, err
	}

	return r.toModel(&company, fiscalData), nil

}

func (r *CompanyRepo) GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelCompany, error) {
	fiscalName = strings.TrimSpace(fiscalName)
	query := `
		SELECT 
			company.*,
			fd.*
		FROM company
		JOIN fiscal_data fd ON company.fiscal_data_id = fd.id
		WHERE LOWER(fd.fiscal_name) = LOWER($1) AND company.country_id = $2
	`

	var company entities.EntityCompany
	if err := r.db.Get(&company, query, fiscalName, countryID); err != nil {
		return nil, types.ThrowData("Error al obtener la empresa por nombre fiscal y país")
	}

	fiscalData, err := r.getFiscalData(company.FiscalDataID)
	if err != nil {
		return nil, err
	}

	return r.toModel(&company, fiscalData), nil

}

func (r *CompanyRepo) GetSuppliersByCompanyID(companyID string) ([]models.CompanySupplierModel, error) {
	var supplierEntities []entities.EntityCompanySupplierWithFiscalData
	query := `
		SELECT
			s.id AS supplier_id,
			s.supplier_name,
			s.description,
			s.available,
			s.country_id,
			fd.id_fiscal,
			fd.raw_fiscal_id,
			fd.fiscal_name,
			fd.fiscal_address,
			fd.email
		FROM supplier_per_company spc
		INNER JOIN supplier s ON s.id = spc.supplier_id
		INNER JOIN fiscal_data fd ON fd.id = s.fiscal_data_id
		WHERE spc.company_id = $1
		ORDER BY s.supplier_name
	`

	if err := r.db.SelectContext(context.TODO(), &supplierEntities, query, companyID); err != nil {
		return nil, types.ThrowData("Error al obtener los proveedores de la empresa")
	}

	modelsList := make([]models.CompanySupplierModel, len(supplierEntities))
	for i, sup := range supplierEntities {
		modelsList[i] = models.CompanySupplierModel{
			SupplierID:    sup.SupplierID,
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

	return modelsList, nil
}

func (r *CompanyRepo) AssignSuppliersToCompany(companyID string, supplierIDs []string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return types.ThrowData("Error al iniciar la asignación de proveedores a la empresa")
	}
	defer tx.Rollback()

	insertQuery := `
		INSERT INTO supplier_per_company (supplier_id, company_id)
		VALUES ($1, $2)
		ON CONFLICT (supplier_id, company_id) DO NOTHING
	`

	for _, supplierID := range supplierIDs {
		if _, err := tx.Exec(insertQuery, supplierID, companyID); err != nil {
			return types.ThrowData("Error al asignar proveedores a la empresa")
		}
	}

	if err := tx.Commit(); err != nil {
		return types.ThrowData("Error al confirmar la asignación de proveedores a la empresa")
	}

	return nil
}

func (r *CompanyRepo) UnassignSupplierFromCompany(companyID, supplierID string) error {
	query := `DELETE FROM supplier_per_company WHERE company_id = $1 AND supplier_id = $2`
	if _, err := r.db.Exec(query, companyID, supplierID); err != nil {
		return types.ThrowData("Error al desasignar el proveedor de la empresa")
	}
	return nil
}

func (r *CompanyRepo) CreateCompany(company *models.ModelCompany, profiles []models.ProfileAccountModel) (*models.ModelCompany, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	queryFiscalData := `
		INSERT INTO fiscal_data 
			(id_fiscal, raw_fiscal_id, fiscal_name, fiscal_address, fiscal_state, fiscal_city, email)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING *
	`

	var createdFiscalData entities.EntityFiscalData
	if err = tx.QueryRowx(queryFiscalData,
		company.FiscalData.IDFiscal,
		company.FiscalData.RawFiscalID,
		company.FiscalData.FiscalName,
		company.FiscalData.FiscalAddress,
		company.FiscalData.FiscalState,
		company.FiscalData.FiscalCity,
		company.FiscalData.Email,
	).StructScan(&createdFiscalData); err != nil {
		return nil, types.ThrowData("Error al crear los datos fiscales de la empresa")
	}

	queryCompany := `
	        INSERT INTO company
				(country_id, fiscal_data_id, company_name, description, image_logo)
	        VALUES
				($1, $2, $3, $4, $5)
	        RETURNING *
	    `

	var createdCompany entities.EntityCompany
	if err = tx.QueryRowx(queryCompany,
		company.CountryID,
		createdFiscalData.ID,
		company.CompanyName,
		company.Description,
		company.ImageLogo,
	).StructScan(&createdCompany); err != nil {
		return nil, types.ThrowData("Error al crear la empresa")
	}

	ownPowerModel := &models.PowerAccountModel{
		PowerName:   fmt.Sprintf("company:%s", createdCompany.ID),
		DisplayName: fmt.Sprintf("Propiedad Empresa %s", createdCompany.CompanyName),
		Description: fmt.Sprintf("Propiedad que otorga permisos exclusivos sobre la empresa %s", createdCompany.CompanyName),
		CategoryID:  "1fdbaca4-e12b-4378-a670-fb15bb509e93", // Harcodeadoooooo por si alguien se enoja en el futuro (deberia salir de una variable de entorno)
	}

	queryPower := `
		INSERT INTO power_accounts
			(power_name, power_display, description, power_account_category_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
	err = tx.QueryRowx(queryPower,
		ownPowerModel.PowerName,
		ownPowerModel.DisplayName,
		ownPowerModel.Description,
		ownPowerModel.CategoryID,
	).Scan(&ownPowerModel.ID)
	if err != nil {
		return nil, types.ThrowData("Error al crear el poder propio para la empresa")
	}

	profileIds := []string{}
	for _, profile := range profiles {
		profileIds = append(profileIds, profile.ID)
	}

	err = r.powerRepo.AddOwnPowerToProfileTx(tx, profileIds, ownPowerModel)
	if err != nil {
		fmt.Println("Error al agregar el poder propio al perfil:", err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModel(&createdCompany, &createdFiscalData), nil
}

func (r *CompanyRepo) UpdateCompany(id string, company *models.ModelCompany) (*models.ModelCompany, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}

	queryFiscalData := `
		UPDATE fiscal_data
		SET 
			fiscal_name = $1,
			fiscal_address = $2,
			fiscal_state = $3,
			fiscal_city = $4,
			email = $5
		WHERE id = $6
		RETURNING *
	`
	var fiscalData entities.EntityFiscalData
	if err := tx.QueryRow(queryFiscalData,
		company.FiscalData.FiscalName,
		company.FiscalData.FiscalAddress,
		company.FiscalData.FiscalState,
		company.FiscalData.FiscalCity,
		company.FiscalData.Email,
		company.FiscalData.ID,
	).Scan(
		&fiscalData.ID,
		&fiscalData.IDFiscal,
		&fiscalData.RawFiscalID,
		&fiscalData.FiscalName,
		&fiscalData.FiscalAddress,
		&fiscalData.FiscalState,
		&fiscalData.FiscalCity,
		&fiscalData.Email,
	); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al actualizar los datos fiscales")
	}

	queryCompany := `
		UPDATE company
		SET 
			country_id = $1,
			company_name = $2,
			description = $3
		WHERE id = $4
		RETURNING *
	`
	var updatedCompany entities.EntityCompany
	if err := tx.QueryRow(queryCompany,
		company.CountryID,
		company.CompanyName,
		company.Description,
		id,
	).Scan(
		&updatedCompany.ID,
		&updatedCompany.CountryID,
		&updatedCompany.FiscalDataID,
		&updatedCompany.CompanyName,
		&updatedCompany.Description,
		&updatedCompany.ImageLogo,
		&updatedCompany.CreatedAt,
		&updatedCompany.UpdatedAt,
	); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al actualizar la empresa")
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModel(&updatedCompany, &fiscalData), nil
}

func (r *CompanyRepo) AddLogoToCompany(companyID string, url string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	query := `
		UPDATE company
		SET image_logo = $1
		WHERE id = $2
	`

	_, err = tx.Exec(query, url, companyID)
	if err != nil {
		return types.ThrowData("Error al agregar el logo a la empresa")
	}

	if err = tx.Commit(); err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}

	return nil
}

func (r *CompanyRepo) RemoveLogoFromCompany(fileID string) error {
	panic("unimplemented")
}

// ///////// Private methods ///////////

func (r *CompanyRepo) getFiscalData(id string) (*entities.EntityFiscalData, error) {
	var fiscalData entities.EntityFiscalData
	if err := r.db.Get(&fiscalData, `
		SELECT *
		FROM fiscal_data
		WHERE id = $1`,
		id,
	); err != nil {
		return nil, types.ThrowData("Error al obtener los datos fiscales")
	}

	return &fiscalData, nil

}

func (r *CompanyRepo) toModel(c *entities.EntityCompany, f *entities.EntityFiscalData) *models.ModelCompany {
	return &models.ModelCompany{
		ID:        c.ID,
		CountryID: c.CountryID,
		FiscalData: models.ModelFiscalData{
			ID:            f.ID,
			IDFiscal:      f.IDFiscal,
			RawFiscalID:   f.RawFiscalID,
			FiscalName:    f.FiscalName,
			FiscalAddress: f.FiscalAddress,
			FiscalState:   f.FiscalState,
			FiscalCity:    f.FiscalCity,
			Email:         f.Email,
		},
		CompanyName: c.CompanyName,
		Description: c.Description,
		ImageLogo:   c.ImageLogo,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
