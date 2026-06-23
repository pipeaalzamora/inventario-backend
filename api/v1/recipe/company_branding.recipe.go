package recipe

import "mime/multipart"

type CompanyBrandingRecipe struct {
	AppName       string  `json:"appName" binding:"required"`
	LogoURL       *string `json:"logoUrl"`
	FaviconURL    *string `json:"faviconUrl"`
	PrimaryColor  string  `json:"primaryColor" binding:"required"`
	AccentColor   string  `json:"accentColor" binding:"required"`
	EmailFromName string  `json:"emailFromName" binding:"required"`
	SupportEmail  *string `json:"supportEmail"`
	CustomDomain  *string `json:"customDomain"`
	Subdomain     *string `json:"subdomain"`
	WelcomeText   *string `json:"welcomeText"`
}

type CompanyBrandingAssetUploadRecipe struct {
	AssetType string                `form:"assetType" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
}

type CompanyEmailTemplateRecipe struct {
	Subject  string  `json:"subject" binding:"required"`
	BodyHTML string  `json:"bodyHtml" binding:"required"`
	BodyText *string `json:"bodyText"`
}

type CompanyImportRecipe struct {
	ImportType string                `form:"importType" binding:"required"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

type CompanyImportExecuteRecipe struct {
	Mapping map[string]string `json:"mapping" binding:"required"`
}
