package recipe

import "mime/multipart"

type RecipeCreateCompany struct {
	CompanyName   string               `form:"companyName" binding:"required" errMsg:"El nombre de la empresa es obligatorio"`
	Description   *string              `form:"description"`
	FiscalName    string               `form:"fiscalName" binding:"required" errMsg:"El nombre fiscal es obligatorio"`
	IDFiscal      string               `form:"idFiscal" binding:"required" errMsg:"El ID fiscal es obligatorio"`
	FiscalAddress string               `form:"fiscalAddress" binding:"required" errMsg:"La dirección fiscal es obligatoria"`
	FiscalState   string               `form:"fiscalState" binding:"required" errMsg:"El estado fiscal es obligatorio"`
	FiscalCity    string               `form:"fiscalCity" binding:"required" errMsg:"La ciudad fiscal es obligatoria"`
	ContactEmail  string               `form:"email" binding:"required,email" errMsg:"El correo electrónico es obligatorio y debe ser válido"`
	ImageLogo     multipart.FileHeader `form:"imageLogo" binding:"required" errMsg:"El logo de la empresa es obligatorio"`
	IsNewLogo     bool                 `form:"isNewLogo"`
}
