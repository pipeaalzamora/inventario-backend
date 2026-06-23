package services

import (
	"context"
	"encoding/csv"
	"html"
	"io"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

var hexColorPattern = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

type CompanyBrandingService struct {
	PowerChecker
	repo   ports.PortCompanyBranding
	bucket ports.PortBucket
}

func NewCompanyBrandingService(repo ports.PortCompanyBranding, bucket ports.PortBucket) *CompanyBrandingService {
	return &CompanyBrandingService{repo: repo, bucket: bucket}
}

func (s *CompanyBrandingService) GetByCompanyID(ctx context.Context, companyID string) (*models.ModelCompanyBranding, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a la configuración de esta empresa")
	}
	return s.repo.GetByCompanyID(companyID)
}

func (s *CompanyBrandingService) Upsert(ctx context.Context, companyID string, input *recipe.CompanyBrandingRecipe) (*models.ModelCompanyBranding, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar la configuración de esta empresa")
	}
	if !hexColorPattern.MatchString(input.PrimaryColor) {
		return nil, types.ThrowRecipe("El color principal debe tener formato hexadecimal #RRGGBB", "primaryColor")
	}
	if !hexColorPattern.MatchString(input.AccentColor) {
		return nil, types.ThrowRecipe("El color secundario debe tener formato hexadecimal #RRGGBB", "accentColor")
	}

	return s.repo.Upsert(&models.ModelCompanyBranding{
		CompanyID:     companyID,
		AppName:       input.AppName,
		LogoURL:       input.LogoURL,
		FaviconURL:    input.FaviconURL,
		PrimaryColor:  input.PrimaryColor,
		AccentColor:   input.AccentColor,
		EmailFromName: input.EmailFromName,
		SupportEmail:  input.SupportEmail,
		CustomDomain:  input.CustomDomain,
		Subdomain:     input.Subdomain,
		WelcomeText:   input.WelcomeText,
	})
}

func (s *CompanyBrandingService) ListAssets(ctx context.Context, companyID string) ([]models.ModelCompanyBrandingAsset, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a los archivos de esta empresa")
	}
	return s.repo.ListAssets(companyID)
}

func (s *CompanyBrandingService) UploadAsset(ctx context.Context, companyID string, input *recipe.CompanyBrandingAssetUploadRecipe) (*models.ModelCompanyBrandingAsset, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar archivos de esta empresa")
	}
	if input.File == nil {
		return nil, types.ThrowRecipe("Debe adjuntar un archivo", "file")
	}

	assetType := strings.ToLower(strings.TrimSpace(input.AssetType))
	if !isAllowedBrandingAssetType(assetType) {
		return nil, types.ThrowRecipe("El tipo de archivo debe ser logo, favicon o image", "assetType")
	}

	mimeType := detectFileMimeType(input.File)
	if !strings.HasPrefix(mimeType, "image/") {
		return nil, types.ThrowRecipe("El archivo debe ser una imagen válida", "file")
	}

	url, err := s.uploadMultipartFile(ctx, input.File)
	if err != nil {
		return nil, err
	}

	asset, err := s.repo.AddAsset(&models.ModelCompanyBrandingAsset{
		CompanyID: companyID,
		AssetType: assetType,
		FileName:  input.File.Filename,
		FileURL:   url,
		MimeType:  mimeType,
		FileSize:  input.File.Size,
	})
	if err != nil {
		return nil, err
	}

	if assetType == "logo" || assetType == "favicon" {
		branding, err := s.repo.GetByCompanyID(companyID)
		if err != nil {
			return nil, err
		}
		if assetType == "logo" {
			branding.LogoURL = &asset.FileURL
		} else {
			branding.FaviconURL = &asset.FileURL
		}
		if _, err := s.repo.Upsert(branding); err != nil {
			return nil, err
		}
	}

	return asset, nil
}

func (s *CompanyBrandingService) DeleteAsset(ctx context.Context, companyID string, assetID string) error {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return types.ThrowPower("No tienes permiso para eliminar archivos de esta empresa")
	}

	asset, err := s.repo.GetAssetByID(assetID)
	if err != nil {
		return err
	}
	if asset.CompanyID != companyID {
		return types.ThrowPower("El archivo no pertenece a esta empresa")
	}

	if s.bucket != nil {
		_ = s.bucket.DeleteFile(ctx, asset.FileURL)
	}

	return s.repo.DeleteAsset(assetID)
}

func (s *CompanyBrandingService) ListEmailTemplates(ctx context.Context, companyID string) ([]models.ModelCompanyEmailTemplate, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a plantillas de esta empresa")
	}

	templates, err := s.repo.ListEmailTemplates(companyID)
	if err != nil {
		return nil, err
	}

	byKey := map[string]models.ModelCompanyEmailTemplate{}
	for _, template := range templates {
		byKey[template.TemplateKey] = template
	}

	defaults := defaultCompanyEmailTemplates(companyID)
	result := make([]models.ModelCompanyEmailTemplate, len(defaults))
	for i, template := range defaults {
		if saved, ok := byKey[template.TemplateKey]; ok {
			result[i] = saved
			continue
		}
		result[i] = template
	}

	return result, nil
}

func (s *CompanyBrandingService) UpsertEmailTemplate(ctx context.Context, companyID string, templateKey string, input *recipe.CompanyEmailTemplateRecipe) (*models.ModelCompanyEmailTemplate, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar plantillas de esta empresa")
	}

	templateKey = strings.ToLower(strings.TrimSpace(templateKey))
	if !isAllowedTemplateKey(templateKey) {
		return nil, types.ThrowRecipe("Plantilla no soportada", "templateKey")
	}
	if strings.TrimSpace(input.Subject) == "" {
		return nil, types.ThrowRecipe("El asunto es obligatorio", "subject")
	}
	if strings.TrimSpace(input.BodyHTML) == "" {
		return nil, types.ThrowRecipe("El contenido HTML es obligatorio", "bodyHtml")
	}

	return s.repo.UpsertEmailTemplate(&models.ModelCompanyEmailTemplate{
		CompanyID:   companyID,
		TemplateKey: templateKey,
		Subject:     input.Subject,
		BodyHTML:    input.BodyHTML,
		BodyText:    input.BodyText,
	})
}

func (s *CompanyBrandingService) ResolveEmailTemplate(ctx context.Context, companyID string, templateKey string, vars map[string]string) (string, string, error) {
	if !isAllowedTemplateKey(templateKey) {
		return "", "", types.ThrowRecipe("Plantilla no soportada", "templateKey")
	}

	branding, err := s.GetByCompanyID(ctx, companyID)
	if err != nil {
		return "", "", err
	}

	templates, err := s.ListEmailTemplates(ctx, companyID)
	if err != nil {
		return "", "", err
	}

	var selected *models.ModelCompanyEmailTemplate
	for i := range templates {
		if templates[i].TemplateKey == templateKey {
			selected = &templates[i]
			break
		}
	}
	if selected == nil {
		defaults := defaultCompanyEmailTemplates(companyID)
		for i := range defaults {
			if defaults[i].TemplateKey == templateKey {
				selected = &defaults[i]
				break
			}
		}
	}
	if selected == nil {
		return "", "", types.ThrowRecipe("Plantilla no soportada", "templateKey")
	}

	renderVars := map[string]string{
		"appName":       branding.AppName,
		"primaryColor":  branding.PrimaryColor,
		"accentColor":   branding.AccentColor,
		"emailFromName": branding.EmailFromName,
	}
	if branding.SupportEmail != nil {
		renderVars["supportEmail"] = *branding.SupportEmail
	}
	for key, value := range vars {
		renderVars[key] = value
	}

	subject := renderTemplateString(selected.Subject, renderVars, false)
	body := renderTemplateString(selected.BodyHTML, renderVars, true)
	return subject, body, nil
}

func (s *CompanyBrandingService) CreateImportJob(ctx context.Context, companyID string, input *recipe.CompanyImportRecipe) (*models.ModelCompanyImportJob, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para importar datos en esta empresa")
	}
	if input.File == nil {
		return nil, types.ThrowRecipe("Debe adjuntar un archivo", "file")
	}

	importType := strings.ToLower(strings.TrimSpace(input.ImportType))
	if !isAllowedImportType(importType) {
		return nil, types.ThrowRecipe("El tipo de importación debe ser products, suppliers o inventory", "importType")
	}

	headers, sampleRows, rows, totalRows, err := parseImportFile(input.File)
	if err != nil {
		return nil, err
	}

	url, err := s.uploadMultipartFile(ctx, input.File)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateImportJob(&models.ModelCompanyImportJob{
		CompanyID:  companyID,
		ImportType: importType,
		FileName:   input.File.Filename,
		FileURL:    url,
		MimeType:   detectFileMimeType(input.File),
		Status:     "previewed",
		TotalRows:  totalRows,
		Headers:    headers,
		SampleRows: sampleRows,
		Rows:       rows,
	})
}

func (s *CompanyBrandingService) ListImportJobs(ctx context.Context, companyID string) ([]models.ModelCompanyImportJob, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a importaciones de esta empresa")
	}
	return s.repo.ListImportJobs(companyID)
}

func (s *CompanyBrandingService) ExecuteImportJob(ctx context.Context, companyID string, importJobID string, input *recipe.CompanyImportExecuteRecipe) (*models.ModelCompanyImportJob, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para ejecutar importaciones en esta empresa")
	}

	job, err := s.repo.GetImportJobByID(companyID, importJobID)
	if err != nil {
		return nil, err
	}
	if job.Status == "executed" || job.Status == "executed_with_errors" {
		return nil, types.ThrowRecipe("Esta importación ya fue ejecutada", "importId")
	}

	mapping := normalizeImportMapping(input.Mapping)
	var processed int
	var rowErrors []string

	switch job.ImportType {
	case "products":
		rows, errors, err := s.buildProductImportRows(job.Rows, mapping)
		if err != nil {
			return nil, err
		}
		processed, rowErrors, err = s.repo.UpsertImportedProducts(rows)
		rowErrors = append(errors, rowErrors...)
		if err != nil {
			return nil, err
		}
	case "suppliers":
		rows, errors, err := s.buildSupplierImportRows(job.Rows, mapping)
		if err != nil {
			return nil, err
		}
		processed, rowErrors, err = s.repo.UpsertImportedSuppliers(companyID, rows)
		rowErrors = append(errors, rowErrors...)
		if err != nil {
			return nil, err
		}
	case "inventory":
		rows, errors, err := s.buildInventoryImportRows(companyID, job.Rows, mapping)
		if err != nil {
			return nil, err
		}
		processed, rowErrors, err = s.repo.UpsertImportedInventory(companyID, rows)
		rowErrors = append(errors, rowErrors...)
		if err != nil {
			return nil, err
		}
	default:
		return nil, types.ThrowRecipe("Tipo de importación no soportado", "importType")
	}

	executedAt := time.Now().UTC()
	job.Mapping = mapping
	job.ProcessedRows = processed
	job.FailedRows = len(rowErrors)
	job.Errors = rowErrors
	job.ExecutedAt = &executedAt
	job.Status = "executed"
	if len(rowErrors) > 0 {
		job.Status = "executed_with_errors"
	}

	return s.repo.UpdateImportJobExecution(job)
}

func (s *CompanyBrandingService) buildProductImportRows(rows []map[string]string, mapping map[string]string) ([]models.ModelProductImportItem, []string, error) {
	if err := requireImportFields(mapping, "name", "sku"); err != nil {
		return nil, nil, err
	}

	result := make([]models.ModelProductImportItem, 0, len(rows))
	rowErrors := []string{}
	for i, row := range rows {
		rowNumber := i + 2
		name := importValue(row, mapping, "name")
		sku := importValue(row, mapping, "sku")
		if name == "" || sku == "" {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": producto requiere nombre y SKU")
			continue
		}

		cost, err := parseImportFloat(importValue(row, mapping, "costEstimated"))
		if err != nil {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": costo estimado inválido")
			continue
		}

		image := optionalString(importValue(row, mapping, "image"))
		result = append(result, models.ModelProductImportItem{
			RowNumber:     rowNumber,
			Name:          name,
			SKU:           sku,
			Description:   importValue(row, mapping, "description"),
			Image:         image,
			CostEstimated: cost,
		})
	}

	return result, rowErrors, nil
}

func (s *CompanyBrandingService) buildSupplierImportRows(rows []map[string]string, mapping map[string]string) ([]models.ModelSupplierImportItem, []string, error) {
	if err := requireImportFields(mapping, "supplierName", "idFiscal", "email"); err != nil {
		return nil, nil, err
	}

	result := make([]models.ModelSupplierImportItem, 0, len(rows))
	rowErrors := []string{}
	for i, row := range rows {
		rowNumber := i + 2
		supplierName := importValue(row, mapping, "supplierName")
		idFiscal := importValue(row, mapping, "idFiscal")
		email := importValue(row, mapping, "email")
		if supplierName == "" || idFiscal == "" || email == "" {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": proveedor requiere nombre, RUT/ID fiscal y email")
			continue
		}

		countryID := 1
		if rawCountryID := importValue(row, mapping, "countryId"); rawCountryID != "" {
			parsed, err := strconv.Atoi(rawCountryID)
			if err != nil {
				rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": país inválido")
				continue
			}
			countryID = parsed
		}

		fiscalName := importValue(row, mapping, "fiscalName")
		if fiscalName == "" {
			fiscalName = supplierName
		}

		contactName := importValue(row, mapping, "contactName")
		if contactName == "" {
			contactName = supplierName
		}

		result = append(result, models.ModelSupplierImportItem{
			RowNumber:     rowNumber,
			SupplierName:  supplierName,
			IDFiscal:      idFiscal,
			Email:         email,
			CountryID:     countryID,
			FiscalName:    fiscalName,
			FiscalAddress: defaultImportValue(importValue(row, mapping, "fiscalAddress"), "Sin dirección"),
			FiscalState:   defaultImportValue(importValue(row, mapping, "fiscalState"), "Sin región"),
			FiscalCity:    defaultImportValue(importValue(row, mapping, "fiscalCity"), "Sin ciudad"),
			Description:   importValue(row, mapping, "description"),
			ContactName:   contactName,
			ContactPhone:  importValue(row, mapping, "contactPhone"),
		})
	}

	return result, rowErrors, nil
}

func (s *CompanyBrandingService) buildInventoryImportRows(companyID string, rows []map[string]string, mapping map[string]string) ([]models.ModelInventoryImportItem, []string, error) {
	if err := requireImportFields(mapping, "storeProductId", "warehouseId", "quantity"); err != nil {
		return nil, nil, err
	}

	result := make([]models.ModelInventoryImportItem, 0, len(rows))
	rowErrors := []string{}
	for i, row := range rows {
		rowNumber := i + 2
		storeProductID := importValue(row, mapping, "storeProductId")
		warehouseID := importValue(row, mapping, "warehouseId")
		quantity, err := parseImportFloat(importValue(row, mapping, "quantity"))
		if err != nil {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": cantidad inválida")
			continue
		}
		costAvg, err := parseImportFloat(importValue(row, mapping, "costAvg"))
		if err != nil {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": costo promedio inválido")
			continue
		}
		if storeProductID == "" || warehouseID == "" {
			rowErrors = append(rowErrors, "Fila "+strconv.Itoa(rowNumber)+": inventario requiere producto de tienda y bodega")
			continue
		}

		result = append(result, models.ModelInventoryImportItem{
			RowNumber:      rowNumber,
			CompanyID:      companyID,
			StoreProductID: storeProductID,
			WarehouseID:    warehouseID,
			Quantity:       quantity,
			CostAvg:        costAvg,
		})
	}

	return result, rowErrors, nil
}

func (s *CompanyBrandingService) uploadMultipartFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	if s.bucket == nil {
		return "", types.ThrowData("El servicio de archivos no está configurado")
	}

	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer openedFile.Close()

	return s.bucket.UploadFile(ctx, openedFile, file.Filename)
}

func renderTemplateString(template string, vars map[string]string, escape bool) string {
	rendered := template
	for key, value := range vars {
		if escape {
			value = html.EscapeString(value)
		}
		rendered = strings.ReplaceAll(rendered, "{{"+key+"}}", value)
	}
	return rendered
}

func normalizeImportMapping(mapping map[string]string) map[string]string {
	normalized := map[string]string{}
	for field, header := range mapping {
		field = strings.TrimSpace(field)
		header = strings.TrimSpace(header)
		if field == "" || header == "" {
			continue
		}
		normalized[field] = header
	}
	return normalized
}

func requireImportFields(mapping map[string]string, fields ...string) error {
	for _, field := range fields {
		if strings.TrimSpace(mapping[field]) == "" {
			return types.ThrowRecipe("Falta mapear el campo "+field, "mapping")
		}
	}
	return nil
}

func importValue(row map[string]string, mapping map[string]string, field string) string {
	header := mapping[field]
	if header == "" {
		return ""
	}
	return strings.TrimSpace(row[header])
}

func defaultImportValue(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func optionalString(value string) *string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return &value
}

func parseImportFloat(value string) (float64, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}
	value = strings.ReplaceAll(value, ",", ".")
	return strconv.ParseFloat(value, 64)
}

func isAllowedBrandingAssetType(assetType string) bool {
	return assetType == "logo" || assetType == "favicon" || assetType == "image"
}

func isAllowedTemplateKey(templateKey string) bool {
	switch templateKey {
	case "welcome", "recovery", "supplier_view":
		return true
	default:
		return false
	}
}

func isAllowedImportType(importType string) bool {
	switch importType {
	case "products", "suppliers", "inventory":
		return true
	default:
		return false
	}
}

func detectFileMimeType(file *multipart.FileHeader) string {
	if mimeType := file.Header.Get("Content-Type"); mimeType != "" {
		return mimeType
	}
	return shared.GetMimeType(filepath.Ext(file.Filename))
}

func defaultCompanyEmailTemplates(companyID string) []models.ModelCompanyEmailTemplate {
	return []models.ModelCompanyEmailTemplate{
		{
			CompanyID:   companyID,
			TemplateKey: "welcome",
			Subject:     "Bienvenido a {{appName}}",
			BodyHTML:    "<p>Hola {{userName}}, bienvenido a {{appName}}.</p>",
		},
		{
			CompanyID:   companyID,
			TemplateKey: "recovery",
			Subject:     "Código de recuperación",
			BodyHTML:    "<p>Tu código de recuperación es <strong>{{code}}</strong>.</p>",
		},
		{
			CompanyID:   companyID,
			TemplateKey: "supplier_view",
			Subject:     "Solicitud de aprobación de orden de compra",
			BodyHTML:    "<p>Revisa la orden de compra en {{url}}. Expira el {{expirationDate}}.</p>",
		},
	}
}

func parseImportFile(file *multipart.FileHeader) ([]string, []map[string]string, []map[string]string, int, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	openedFile, err := file.Open()
	if err != nil {
		return nil, nil, nil, 0, err
	}
	defer openedFile.Close()

	switch ext {
	case ".csv":
		reader := csv.NewReader(openedFile)
		reader.TrimLeadingSpace = true
		rows, err := reader.ReadAll()
		if err != nil {
			return nil, nil, nil, 0, types.ThrowRecipe("No se pudo leer el CSV", "file")
		}
		return rowsToImportData(rows)
	case ".xlsx":
		xlsx, err := excelize.OpenReader(openedFile)
		if err != nil {
			return nil, nil, nil, 0, types.ThrowRecipe("No se pudo leer el Excel", "file")
		}
		defer xlsx.Close()

		sheets := xlsx.GetSheetList()
		if len(sheets) == 0 {
			return nil, nil, nil, 0, types.ThrowRecipe("El Excel no tiene hojas", "file")
		}
		rows, err := xlsx.GetRows(sheets[0])
		if err != nil {
			return nil, nil, nil, 0, types.ThrowRecipe("No se pudo leer la primera hoja del Excel", "file")
		}
		return rowsToImportData(rows)
	case ".xls":
		return nil, nil, nil, 0, types.ThrowRecipe("Use formato .xlsx para archivos Excel", "file")
	default:
		if _, err := io.Copy(io.Discard, openedFile); err != nil {
			return nil, nil, nil, 0, err
		}
		return nil, nil, nil, 0, types.ThrowRecipe("El archivo debe ser CSV o XLSX", "file")
	}
}

func rowsToImportData(rows [][]string) ([]string, []map[string]string, []map[string]string, int, error) {
	if len(rows) < 2 {
		return nil, nil, nil, 0, types.ThrowRecipe("El archivo debe tener encabezados y al menos una fila de datos", "file")
	}

	headers := make([]string, len(rows[0]))
	for i, header := range rows[0] {
		header = strings.TrimSpace(header)
		if header == "" {
			header = "columna_" + strconv.Itoa(i+1)
		}
		headers[i] = header
	}

	totalRows := len(rows) - 1
	limit := totalRows
	if limit > 20 {
		limit = 20
	}

	allRows := make([]map[string]string, 0, totalRows)
	for _, row := range rows[1:] {
		item := map[string]string{}
		for i, header := range headers {
			value := ""
			if i < len(row) {
				value = row[i]
			}
			item[header] = value
		}
		allRows = append(allRows, item)
	}

	sampleRows := allRows
	if len(sampleRows) > limit {
		sampleRows = sampleRows[:limit]
	}

	return headers, sampleRows, allRows, totalRows, nil
}
