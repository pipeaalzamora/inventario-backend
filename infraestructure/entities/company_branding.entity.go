package entities

import "time"

type EntityCompanyBranding struct {
	CompanyID     string    `db:"company_id"`
	AppName       string    `db:"app_name"`
	LogoURL       *string   `db:"logo_url"`
	FaviconURL    *string   `db:"favicon_url"`
	PrimaryColor  string    `db:"primary_color"`
	AccentColor   string    `db:"accent_color"`
	EmailFromName string    `db:"email_from_name"`
	SupportEmail  *string   `db:"support_email"`
	CustomDomain  *string   `db:"custom_domain"`
	Subdomain     *string   `db:"subdomain"`
	WelcomeText   *string   `db:"welcome_text"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type EntityCompanyBrandingAsset struct {
	ID        string    `db:"id"`
	CompanyID string    `db:"company_id"`
	AssetType string    `db:"asset_type"`
	FileName  string    `db:"file_name"`
	FileURL   string    `db:"file_url"`
	MimeType  string    `db:"mime_type"`
	FileSize  int64     `db:"file_size"`
	CreatedAt time.Time `db:"created_at"`
}

type EntityCompanyEmailTemplate struct {
	ID          string    `db:"id"`
	CompanyID   string    `db:"company_id"`
	TemplateKey string    `db:"template_key"`
	Subject     string    `db:"subject"`
	BodyHTML    string    `db:"body_html"`
	BodyText    *string   `db:"body_text"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type EntityCompanyImportJob struct {
	ID            string     `db:"id"`
	CompanyID     string     `db:"company_id"`
	ImportType    string     `db:"import_type"`
	FileName      string     `db:"file_name"`
	FileURL       string     `db:"file_url"`
	MimeType      string     `db:"mime_type"`
	Status        string     `db:"status"`
	TotalRows     int        `db:"total_rows"`
	Headers       []byte     `db:"headers"`
	SampleRows    []byte     `db:"sample_rows"`
	Rows          []byte     `db:"rows"`
	Mapping       []byte     `db:"mapping"`
	ProcessedRows int        `db:"processed_rows"`
	FailedRows    int        `db:"failed_rows"`
	Errors        []byte     `db:"errors"`
	ExecutedAt    *time.Time `db:"executed_at"`
	CreatedAt     time.Time  `db:"created_at"`
}
