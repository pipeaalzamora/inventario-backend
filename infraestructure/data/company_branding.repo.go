package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"time"

	"github.com/jmoiron/sqlx"
)

type CompanyBrandingRepo struct {
	db *sqlx.DB
}

func NewCompanyBrandingRepo(db *sqlx.DB) ports.PortCompanyBranding {
	return &CompanyBrandingRepo{db: db}
}

func (r *CompanyBrandingRepo) GetByCompanyID(companyID string) (*models.ModelCompanyBranding, error) {
	var entity entities.EntityCompanyBranding
	err := r.db.Get(&entity, `
		SELECT
			cb.company_id,
			cb.app_name,
			cb.logo_url,
			cb.favicon_url,
			cb.primary_color,
			cb.accent_color,
			cb.email_from_name,
			cb.support_email,
			cb.custom_domain,
			cb.subdomain,
			cb.welcome_text,
			cb.updated_at
		FROM company_branding cb
		WHERE cb.company_id = $1
	`, companyID)
	if err == nil {
		return r.toModel(&entity), nil
	}
	if err != sql.ErrNoRows {
		return nil, types.ThrowData("Error al obtener la configuración de marca")
	}

	var company struct {
		CompanyName string  `db:"company_name"`
		ImageLogo   *string `db:"image_logo"`
	}
	err = r.db.Get(&company, `
		SELECT company_name, image_logo
		FROM company
		WHERE id = $1
	`, companyID)
	if err != nil {
		return nil, types.ThrowData("No se encontró la empresa")
	}

	return &models.ModelCompanyBranding{
		CompanyID:     companyID,
		AppName:       company.CompanyName,
		LogoURL:       company.ImageLogo,
		FaviconURL:    nil,
		PrimaryColor:  "#2563eb",
		AccentColor:   "#16a34a",
		EmailFromName: company.CompanyName,
		UpdatedAt:     time.Now().UTC(),
	}, nil
}

func (r *CompanyBrandingRepo) Upsert(branding *models.ModelCompanyBranding) (*models.ModelCompanyBranding, error) {
	var entity entities.EntityCompanyBranding
	err := r.db.QueryRowx(`
		INSERT INTO company_branding (
			company_id,
			app_name,
			logo_url,
			favicon_url,
			primary_color,
			accent_color,
			email_from_name,
			support_email,
			custom_domain,
			subdomain,
			welcome_text,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW())
		ON CONFLICT (company_id) DO UPDATE SET
			app_name = EXCLUDED.app_name,
			logo_url = EXCLUDED.logo_url,
			favicon_url = EXCLUDED.favicon_url,
			primary_color = EXCLUDED.primary_color,
			accent_color = EXCLUDED.accent_color,
			email_from_name = EXCLUDED.email_from_name,
			support_email = EXCLUDED.support_email,
			custom_domain = EXCLUDED.custom_domain,
			subdomain = EXCLUDED.subdomain,
			welcome_text = EXCLUDED.welcome_text,
			updated_at = NOW()
		RETURNING
			company_id,
			app_name,
			logo_url,
			favicon_url,
			primary_color,
			accent_color,
			email_from_name,
			support_email,
			custom_domain,
			subdomain,
			welcome_text,
			updated_at
	`,
		branding.CompanyID,
		branding.AppName,
		branding.LogoURL,
		branding.FaviconURL,
		branding.PrimaryColor,
		branding.AccentColor,
		branding.EmailFromName,
		branding.SupportEmail,
		branding.CustomDomain,
		branding.Subdomain,
		branding.WelcomeText,
	).StructScan(&entity)
	if err != nil {
		return nil, types.ThrowData("Error al guardar la configuración de marca")
	}

	return r.toModel(&entity), nil
}

func (r *CompanyBrandingRepo) AddAsset(asset *models.ModelCompanyBrandingAsset) (*models.ModelCompanyBrandingAsset, error) {
	var entity entities.EntityCompanyBrandingAsset
	err := r.db.QueryRowx(`
		INSERT INTO company_branding_assets (
			company_id,
			asset_type,
			file_name,
			file_url,
			mime_type,
			file_size
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, company_id, asset_type, file_name, file_url, mime_type, file_size, created_at
	`,
		asset.CompanyID,
		asset.AssetType,
		asset.FileName,
		asset.FileURL,
		asset.MimeType,
		asset.FileSize,
	).StructScan(&entity)
	if err != nil {
		return nil, types.ThrowData("Error al guardar el archivo de marca")
	}

	return r.assetToModel(&entity), nil
}

func (r *CompanyBrandingRepo) GetAssetByID(assetID string) (*models.ModelCompanyBrandingAsset, error) {
	var entity entities.EntityCompanyBrandingAsset
	err := r.db.Get(&entity, `
		SELECT id, company_id, asset_type, file_name, file_url, mime_type, file_size, created_at
		FROM company_branding_assets
		WHERE id = $1
	`, assetID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.ThrowData("Archivo no encontrado")
		}
		return nil, types.ThrowData("Error al obtener el archivo de marca")
	}
	return r.assetToModel(&entity), nil
}

func (r *CompanyBrandingRepo) ListAssets(companyID string) ([]models.ModelCompanyBrandingAsset, error) {
	var entitiesList []entities.EntityCompanyBrandingAsset
	err := r.db.Select(&entitiesList, `
		SELECT id, company_id, asset_type, file_name, file_url, mime_type, file_size, created_at
		FROM company_branding_assets
		WHERE company_id = $1
		ORDER BY created_at DESC
	`, companyID)
	if err != nil {
		return nil, types.ThrowData("Error al listar archivos de marca")
	}

	result := make([]models.ModelCompanyBrandingAsset, len(entitiesList))
	for i, entity := range entitiesList {
		result[i] = *r.assetToModel(&entity)
	}
	return result, nil
}

func (r *CompanyBrandingRepo) DeleteAsset(assetID string) error {
	_, err := r.db.Exec(`DELETE FROM company_branding_assets WHERE id = $1`, assetID)
	if err != nil {
		return types.ThrowData("Error al eliminar el archivo de marca")
	}
	return nil
}

func (r *CompanyBrandingRepo) ListEmailTemplates(companyID string) ([]models.ModelCompanyEmailTemplate, error) {
	var entitiesList []entities.EntityCompanyEmailTemplate
	err := r.db.Select(&entitiesList, `
		SELECT id, company_id, template_key, subject, body_html, body_text, updated_at
		FROM company_email_templates
		WHERE company_id = $1
		ORDER BY template_key ASC
	`, companyID)
	if err != nil {
		return nil, types.ThrowData("Error al listar plantillas de correo")
	}

	result := make([]models.ModelCompanyEmailTemplate, len(entitiesList))
	for i, entity := range entitiesList {
		result[i] = *r.templateToModel(&entity)
	}
	return result, nil
}

func (r *CompanyBrandingRepo) UpsertEmailTemplate(template *models.ModelCompanyEmailTemplate) (*models.ModelCompanyEmailTemplate, error) {
	var entity entities.EntityCompanyEmailTemplate
	err := r.db.QueryRowx(`
		INSERT INTO company_email_templates (
			company_id,
			template_key,
			subject,
			body_html,
			body_text,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, NOW())
		ON CONFLICT (company_id, template_key) DO UPDATE SET
			subject = EXCLUDED.subject,
			body_html = EXCLUDED.body_html,
			body_text = EXCLUDED.body_text,
			updated_at = NOW()
		RETURNING id, company_id, template_key, subject, body_html, body_text, updated_at
	`,
		template.CompanyID,
		template.TemplateKey,
		template.Subject,
		template.BodyHTML,
		template.BodyText,
	).StructScan(&entity)
	if err != nil {
		return nil, types.ThrowData("Error al guardar la plantilla de correo")
	}

	return r.templateToModel(&entity), nil
}

func (r *CompanyBrandingRepo) CreateImportJob(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	headers, err := json.Marshal(job.Headers)
	if err != nil {
		return nil, err
	}
	sampleRows, err := json.Marshal(job.SampleRows)
	if err != nil {
		return nil, err
	}
	rows, err := json.Marshal(job.Rows)
	if err != nil {
		return nil, err
	}

	var entity entities.EntityCompanyImportJob
	err = r.db.QueryRowx(`
		INSERT INTO company_import_jobs (
			company_id,
			import_type,
			file_name,
			file_url,
			mime_type,
			status,
			total_rows,
			headers,
			sample_rows,
			rows
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb, $9::jsonb, $10::jsonb)
		RETURNING id, company_id, import_type, file_name, file_url, mime_type, status, total_rows, headers, sample_rows, rows, mapping, processed_rows, failed_rows, errors, executed_at, created_at
	`,
		job.CompanyID,
		job.ImportType,
		job.FileName,
		job.FileURL,
		job.MimeType,
		job.Status,
		job.TotalRows,
		string(headers),
		string(sampleRows),
		string(rows),
	).StructScan(&entity)
	if err != nil {
		return nil, types.ThrowData("Error al registrar la importación")
	}

	return r.importJobToModel(&entity)
}

func (r *CompanyBrandingRepo) ListImportJobs(companyID string) ([]models.ModelCompanyImportJob, error) {
	var entitiesList []entities.EntityCompanyImportJob
	err := r.db.Select(&entitiesList, `
		SELECT id, company_id, import_type, file_name, file_url, mime_type, status, total_rows, headers, sample_rows, rows, mapping, processed_rows, failed_rows, errors, executed_at, created_at
		FROM company_import_jobs
		WHERE company_id = $1
		ORDER BY created_at DESC
	`, companyID)
	if err != nil {
		return nil, types.ThrowData("Error al listar importaciones")
	}

	result := make([]models.ModelCompanyImportJob, len(entitiesList))
	for i, entity := range entitiesList {
		model, err := r.importJobToModel(&entity)
		if err != nil {
			return nil, err
		}
		result[i] = *model
	}
	return result, nil
}

func (r *CompanyBrandingRepo) GetImportJobByID(companyID string, importJobID string) (*models.ModelCompanyImportJob, error) {
	var entity entities.EntityCompanyImportJob
	err := r.db.Get(&entity, `
		SELECT id, company_id, import_type, file_name, file_url, mime_type, status, total_rows, headers, sample_rows, rows, mapping, processed_rows, failed_rows, errors, executed_at, created_at
		FROM company_import_jobs
		WHERE company_id = $1 AND id = $2
	`, companyID, importJobID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.ThrowData("Importación no encontrada")
		}
		return nil, types.ThrowData("Error al obtener la importación")
	}

	return r.importJobToModel(&entity)
}

func (r *CompanyBrandingRepo) UpdateImportJobExecution(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	mapping, err := json.Marshal(job.Mapping)
	if err != nil {
		return nil, err
	}
	errorsPayload, err := json.Marshal(job.Errors)
	if err != nil {
		return nil, err
	}

	var entity entities.EntityCompanyImportJob
	err = r.db.QueryRowx(`
		UPDATE company_import_jobs
		SET status = $1,
			mapping = $2::jsonb,
			processed_rows = $3,
			failed_rows = $4,
			errors = $5::jsonb,
			executed_at = $6
		WHERE id = $7 AND company_id = $8
		RETURNING id, company_id, import_type, file_name, file_url, mime_type, status, total_rows, headers, sample_rows, rows, mapping, processed_rows, failed_rows, errors, executed_at, created_at
	`,
		job.Status,
		string(mapping),
		job.ProcessedRows,
		job.FailedRows,
		string(errorsPayload),
		job.ExecutedAt,
		job.ID,
		job.CompanyID,
	).StructScan(&entity)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar la importación")
	}

	return r.importJobToModel(&entity)
}

func (r *CompanyBrandingRepo) UpsertImportedProducts(rows []models.ModelProductImportItem) (int, []string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, nil, types.ThrowData("Error al iniciar la importación de productos")
	}
	defer tx.Rollback()

	processed := 0
	rowErrors := []string{}
	for _, row := range rows {
		_, err := tx.Exec(`
			INSERT INTO product (product_name, description, image, cost_estimated, sku)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (sku) DO UPDATE SET
				product_name = EXCLUDED.product_name,
				description = EXCLUDED.description,
				image = EXCLUDED.image,
				cost_estimated = EXCLUDED.cost_estimated,
				updated_at = NOW()
		`, row.Name, row.Description, row.Image, row.CostEstimated, row.SKU)
		if err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo guardar producto SKU %s", row.RowNumber, row.SKU))
			continue
		}
		processed++
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, types.ThrowData("Error al confirmar la importación de productos")
	}

	return processed, rowErrors, nil
}

func (r *CompanyBrandingRepo) UpsertImportedSuppliers(companyID string, rows []models.ModelSupplierImportItem) (int, []string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, nil, types.ThrowData("Error al iniciar la importación de proveedores")
	}
	defer tx.Rollback()

	processed := 0
	rowErrors := []string{}
	for _, row := range rows {
		var fiscalDataID string
		err := tx.QueryRowx(`
			INSERT INTO fiscal_data (id_fiscal, raw_fiscal_id, fiscal_name, fiscal_address, fiscal_state, fiscal_city, email)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (id_fiscal) DO UPDATE SET
				raw_fiscal_id = EXCLUDED.raw_fiscal_id,
				fiscal_name = EXCLUDED.fiscal_name,
				fiscal_address = EXCLUDED.fiscal_address,
				fiscal_state = EXCLUDED.fiscal_state,
				fiscal_city = EXCLUDED.fiscal_city,
				email = EXCLUDED.email
			RETURNING id
		`,
			row.IDFiscal,
			row.IDFiscal,
			row.FiscalName,
			row.FiscalAddress,
			row.FiscalState,
			row.FiscalCity,
			row.Email,
		).Scan(&fiscalDataID)
		if err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudieron guardar datos fiscales", row.RowNumber))
			continue
		}

		var supplierID string
		err = tx.QueryRowx(`
			SELECT id
			FROM supplier
			WHERE fiscal_data_id = $1 AND country_id = $2
			LIMIT 1
		`, fiscalDataID, row.CountryID).Scan(&supplierID)
		if err == sql.ErrNoRows {
			err = tx.QueryRowx(`
				INSERT INTO supplier (fiscal_data_id, country_id, supplier_name, description, available)
				VALUES ($1, $2, $3, $4, TRUE)
				RETURNING id
			`, fiscalDataID, row.CountryID, row.SupplierName, row.Description).Scan(&supplierID)
		} else if err == nil {
			_, err = tx.Exec(`
				UPDATE supplier
				SET supplier_name = $1, description = $2, updated_at = NOW()
				WHERE id = $3
			`, row.SupplierName, row.Description, supplierID)
		}
		if err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo guardar proveedor", row.RowNumber))
			continue
		}

		if _, err := tx.Exec(`
			INSERT INTO supplier_per_company (supplier_id, company_id)
			VALUES ($1, $2)
			ON CONFLICT (supplier_id, company_id) DO NOTHING
		`, supplierID, companyID); err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo asignar proveedor a empresa", row.RowNumber))
			continue
		}

		var contactExists bool
		if err := tx.QueryRowx(`
			SELECT EXISTS (
				SELECT 1
				FROM supplier_contact
				WHERE supplier_id = $1 AND LOWER(email) = LOWER($2)
			)
		`, supplierID, row.Email).Scan(&contactExists); err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo validar contacto", row.RowNumber))
			continue
		}
		if !contactExists {
			if _, err := tx.Exec(`
				INSERT INTO supplier_contact (supplier_id, contact_name, description, email, phone)
				VALUES ($1, $2, $3, $4, $5)
			`, supplierID, row.ContactName, "", row.Email, row.ContactPhone); err != nil {
				rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo crear contacto", row.RowNumber))
				continue
			}
		}

		processed++
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, types.ThrowData("Error al confirmar la importación de proveedores")
	}

	return processed, rowErrors, nil
}

func (r *CompanyBrandingRepo) UpsertImportedInventory(companyID string, rows []models.ModelInventoryImportItem) (int, []string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, nil, types.ThrowData("Error al iniciar la importación de inventario")
	}
	defer tx.Rollback()

	processed := 0
	rowErrors := []string{}
	for _, row := range rows {
		var belongs bool
		err := tx.QueryRowx(`
			SELECT EXISTS (
				SELECT 1
				FROM product_per_store pps
				INNER JOIN store s ON s.id = pps.store_id
				INNER JOIN warehouse w ON w.store_id = s.id
				WHERE pps.id = $1
				  AND w.id = $2
				  AND s.company_id = $3
			)
		`, row.StoreProductID, row.WarehouseID, companyID).Scan(&belongs)
		if err != nil || !belongs {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: producto/bodega no pertenecen a la empresa", row.RowNumber))
			continue
		}

		result, err := tx.Exec(`
			UPDATE warehouse_per_product
			SET in_stock = $1, cost_avg = $2
			WHERE store_product_id = $3 AND warehouse_id = $4
		`, row.Quantity, row.CostAvg, row.StoreProductID, row.WarehouseID)
		if err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo actualizar inventario", row.RowNumber))
			continue
		}

		affected, err := result.RowsAffected()
		if err != nil {
			rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo confirmar inventario", row.RowNumber))
			continue
		}
		if affected == 0 {
			if _, err := tx.Exec(`
				INSERT INTO warehouse_per_product (store_product_id, warehouse_id, direction, in_stock, cost_avg)
				VALUES ($1, $2, 'IMPORT', $3, $4)
			`, row.StoreProductID, row.WarehouseID, row.Quantity, row.CostAvg); err != nil {
				rowErrors = append(rowErrors, fmt.Sprintf("Fila %d: no se pudo crear inventario", row.RowNumber))
				continue
			}
		}

		processed++
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, types.ThrowData("Error al confirmar la importación de inventario")
	}

	return processed, rowErrors, nil
}

func (r *CompanyBrandingRepo) toModel(entity *entities.EntityCompanyBranding) *models.ModelCompanyBranding {
	return &models.ModelCompanyBranding{
		CompanyID:     entity.CompanyID,
		AppName:       entity.AppName,
		LogoURL:       entity.LogoURL,
		FaviconURL:    entity.FaviconURL,
		PrimaryColor:  entity.PrimaryColor,
		AccentColor:   entity.AccentColor,
		EmailFromName: entity.EmailFromName,
		SupportEmail:  entity.SupportEmail,
		CustomDomain:  entity.CustomDomain,
		Subdomain:     entity.Subdomain,
		WelcomeText:   entity.WelcomeText,
		UpdatedAt:     entity.UpdatedAt,
	}
}

func (r *CompanyBrandingRepo) assetToModel(entity *entities.EntityCompanyBrandingAsset) *models.ModelCompanyBrandingAsset {
	return &models.ModelCompanyBrandingAsset{
		ID:        entity.ID,
		CompanyID: entity.CompanyID,
		AssetType: entity.AssetType,
		FileName:  entity.FileName,
		FileURL:   entity.FileURL,
		MimeType:  entity.MimeType,
		FileSize:  entity.FileSize,
		CreatedAt: entity.CreatedAt,
	}
}

func (r *CompanyBrandingRepo) templateToModel(entity *entities.EntityCompanyEmailTemplate) *models.ModelCompanyEmailTemplate {
	return &models.ModelCompanyEmailTemplate{
		ID:          entity.ID,
		CompanyID:   entity.CompanyID,
		TemplateKey: entity.TemplateKey,
		Subject:     entity.Subject,
		BodyHTML:    entity.BodyHTML,
		BodyText:    entity.BodyText,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (r *CompanyBrandingRepo) importJobToModel(entity *entities.EntityCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	var headers []string
	if len(entity.Headers) > 0 {
		if err := json.Unmarshal(entity.Headers, &headers); err != nil {
			return nil, err
		}
	}

	var sampleRows []map[string]string
	if len(entity.SampleRows) > 0 {
		if err := json.Unmarshal(entity.SampleRows, &sampleRows); err != nil {
			return nil, err
		}
	}

	var rows []map[string]string
	if len(entity.Rows) > 0 {
		if err := json.Unmarshal(entity.Rows, &rows); err != nil {
			return nil, err
		}
	}

	var mapping map[string]string
	if len(entity.Mapping) > 0 {
		if err := json.Unmarshal(entity.Mapping, &mapping); err != nil {
			return nil, err
		}
	}

	var rowErrors []string
	if len(entity.Errors) > 0 {
		if err := json.Unmarshal(entity.Errors, &rowErrors); err != nil {
			return nil, err
		}
	}

	return &models.ModelCompanyImportJob{
		ID:            entity.ID,
		CompanyID:     entity.CompanyID,
		ImportType:    entity.ImportType,
		FileName:      entity.FileName,
		FileURL:       entity.FileURL,
		MimeType:      entity.MimeType,
		Status:        entity.Status,
		TotalRows:     entity.TotalRows,
		Headers:       headers,
		SampleRows:    sampleRows,
		Rows:          rows,
		Mapping:       mapping,
		ProcessedRows: entity.ProcessedRows,
		FailedRows:    entity.FailedRows,
		Errors:        rowErrors,
		ExecutedAt:    entity.ExecutedAt,
		CreatedAt:     entity.CreatedAt,
	}, nil
}
