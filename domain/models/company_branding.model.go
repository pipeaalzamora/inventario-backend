package models

import "time"

type ModelCompanyBranding struct {
	CompanyID     string    `json:"companyId"`
	AppName       string    `json:"appName"`
	LogoURL       *string   `json:"logoUrl"`
	FaviconURL    *string   `json:"faviconUrl"`
	PrimaryColor  string    `json:"primaryColor"`
	AccentColor   string    `json:"accentColor"`
	EmailFromName string    `json:"emailFromName"`
	SupportEmail  *string   `json:"supportEmail"`
	CustomDomain  *string   `json:"customDomain"`
	Subdomain     *string   `json:"subdomain"`
	WelcomeText   *string   `json:"welcomeText"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ModelCompanyBrandingAsset struct {
	ID        string    `json:"id"`
	CompanyID string    `json:"companyId"`
	AssetType string    `json:"assetType"`
	FileName  string    `json:"fileName"`
	FileURL   string    `json:"fileUrl"`
	MimeType  string    `json:"mimeType"`
	FileSize  int64     `json:"fileSize"`
	CreatedAt time.Time `json:"createdAt"`
}

type ModelCompanyEmailTemplate struct {
	ID          string    `json:"id"`
	CompanyID   string    `json:"companyId"`
	TemplateKey string    `json:"templateKey"`
	Subject     string    `json:"subject"`
	BodyHTML    string    `json:"bodyHtml"`
	BodyText    *string   `json:"bodyText"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ModelCompanyImportJob struct {
	ID            string              `json:"id"`
	CompanyID     string              `json:"companyId"`
	ImportType    string              `json:"importType"`
	FileName      string              `json:"fileName"`
	FileURL       string              `json:"fileUrl"`
	MimeType      string              `json:"mimeType"`
	Status        string              `json:"status"`
	TotalRows     int                 `json:"totalRows"`
	Headers       []string            `json:"headers"`
	SampleRows    []map[string]string `json:"sampleRows"`
	Rows          []map[string]string `json:"-"`
	Mapping       map[string]string   `json:"mapping,omitempty"`
	ProcessedRows int                 `json:"processedRows"`
	FailedRows    int                 `json:"failedRows"`
	Errors        []string            `json:"errors,omitempty"`
	ExecutedAt    *time.Time          `json:"executedAt"`
	CreatedAt     time.Time           `json:"createdAt"`
}

type ModelProductImportItem struct {
	RowNumber     int
	Name          string
	SKU           string
	Description   string
	Image         *string
	CostEstimated float64
}

type ModelSupplierImportItem struct {
	RowNumber     int
	SupplierName  string
	IDFiscal      string
	Email         string
	CountryID     int
	FiscalName    string
	FiscalAddress string
	FiscalState   string
	FiscalCity    string
	Description   string
	ContactName   string
	ContactPhone  string
}

type ModelInventoryImportItem struct {
	RowNumber      int
	CompanyID      string
	StoreProductID string
	WarehouseID    string
	Quantity       float64
	CostAvg        float64
}
