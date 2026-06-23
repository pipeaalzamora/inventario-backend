package controllers

import (
	"net/http"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyFacade *facades.CompanyFacade
}

func NewCompanyController(companyFacade *facades.CompanyFacade) *CompanyController {
	return &CompanyController{
		companyFacade: companyFacade,
	}
}

func (c *CompanyController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/companies")

	r.GET("", c.getCompanies)
	r.GET(":id", c.getCompanyByID)
	r.GET(":id/branding", c.getCompanyBranding)
	r.GET(":id/branding/assets", c.getCompanyBrandingAssets)
	r.GET(":id/branding/templates", c.getCompanyEmailTemplates)
	r.GET(":id/branding/imports", c.getCompanyImportJobs)
	r.GET(":id/suppliers", c.getCompanySuppliers)

	r.POST("", c.createCompany)
	r.POST(":id", c.updateCompany)
	r.POST(":id/branding/assets", c.uploadCompanyBrandingAsset)
	r.POST(":id/branding/imports", c.createCompanyImportJob)
	r.POST(":id/branding/imports/:importId/execute", c.executeCompanyImportJob)
	r.PUT(":id/branding", c.upsertCompanyBranding)
	r.PUT(":id/branding/templates/:templateKey", c.upsertCompanyEmailTemplate)
	r.POST(":id/suppliers", c.assignSuppliersToCompany)
	r.DELETE(":id/branding/assets/:assetId", c.deleteCompanyBrandingAsset)
	r.DELETE(":id/suppliers/:supplierId", c.unassignSupplierFromCompany)

}

func (c *CompanyController) getCompanyBranding(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	branding, err := c.companyFacade.GetCompanyBranding(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, branding)
}

func (c *CompanyController) getCompanyBrandingAssets(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	assets, err := c.companyFacade.ListCompanyBrandingAssets(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"assets": assets})
}

func (c *CompanyController) uploadCompanyBrandingAsset(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.CompanyBrandingAssetUploadRecipe
	if err := shared.Bind(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	asset, err := c.companyFacade.UploadCompanyBrandingAsset(gctx, params.Id, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, asset)
}

func (c *CompanyController) deleteCompanyBrandingAsset(gctx *gin.Context) {
	type pathParams struct {
		Id      string `uri:"id" binding:"required"`
		AssetID string `uri:"assetId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	if err := c.companyFacade.DeleteCompanyBrandingAsset(gctx, params.Id, params.AssetID); err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusNoContent)
}

func (c *CompanyController) getCompanyEmailTemplates(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	templates, err := c.companyFacade.ListCompanyEmailTemplates(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"templates": templates})
}

func (c *CompanyController) upsertCompanyEmailTemplate(gctx *gin.Context) {
	type pathParams struct {
		Id          string `uri:"id" binding:"required"`
		TemplateKey string `uri:"templateKey" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.CompanyEmailTemplateRecipe
	if err := shared.BindJSON(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	template, err := c.companyFacade.UpsertCompanyEmailTemplate(gctx, params.Id, params.TemplateKey, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, template)
}

func (c *CompanyController) createCompanyImportJob(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.CompanyImportRecipe
	if err := shared.Bind(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	job, err := c.companyFacade.CreateCompanyImportJob(gctx, params.Id, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, job)
}

func (c *CompanyController) getCompanyImportJobs(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	jobs, err := c.companyFacade.ListCompanyImportJobs(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"imports": jobs})
}

func (c *CompanyController) executeCompanyImportJob(gctx *gin.Context) {
	type pathParams struct {
		Id       string `uri:"id" binding:"required"`
		ImportID string `uri:"importId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.CompanyImportExecuteRecipe
	if err := shared.BindJSON(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	job, err := c.companyFacade.ExecuteCompanyImportJob(gctx, params.Id, params.ImportID, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, job)
}

func (c *CompanyController) upsertCompanyBranding(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var input recipe.CompanyBrandingRecipe
	if err := shared.BindJSON(gctx, &input); err != nil {
		gctx.Error(err)
		return
	}

	branding, err := c.companyFacade.UpsertCompanyBranding(gctx, params.Id, &input)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, branding)
}

func (c *CompanyController) getCompanies(gctx *gin.Context) {
	companies, err := c.companyFacade.GetCompanies(gctx)
	if err != nil {
		gctx.Error(err)
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"companies": companies})
}

func (c *CompanyController) getCompanyByID(gctx *gin.Context) {
	type pathParams struct {
		ID string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	company, err := c.companyFacade.GetCompanyByID(gctx, params.ID)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, company)
}

func (c *CompanyController) createCompany(gctx *gin.Context) {
	var recipe recipe.RecipeCreateCompany

	if err := shared.Bind(gctx, &recipe); err != nil {
		gctx.Error(err)
		return
	}

	createdCompany, err := c.companyFacade.CreateCompany(gctx, &recipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusCreated, createdCompany)
}

func (c *CompanyController) updateCompany(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var company recipe.RecipeCreateCompany
	if err := shared.Bind(gctx, &company); err != nil {
		gctx.Error(err)
		return
	}

	updatedCompany, err := c.companyFacade.UpdateCompany(gctx, params.Id, &company)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, updatedCompany)
}

func (c *CompanyController) getCompanySuppliers(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	suppliers, err := c.companyFacade.GetCompanySuppliers(gctx, params.Id)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"suppliers": suppliers})
}

func (c *CompanyController) assignSuppliersToCompany(gctx *gin.Context) {
	type pathParams struct {
		Id string `uri:"id" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	var supplierRecipe recipe.AssignSuppliersToCompanyRecipe
	if err := shared.BindJSON(gctx, &supplierRecipe); err != nil {
		gctx.Error(err)
		return
	}

	suppliers, err := c.companyFacade.AssignSuppliersToCompany(gctx, params.Id, &supplierRecipe)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"suppliers": suppliers})
}

func (c *CompanyController) unassignSupplierFromCompany(gctx *gin.Context) {
	type pathParams struct {
		Id         string `uri:"id" binding:"required"`
		SupplierId string `uri:"supplierId" binding:"required"`
	}

	var params pathParams
	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.Error(err)
		return
	}

	suppliers, err := c.companyFacade.UnassignSupplierFromCompany(gctx, params.Id, params.SupplierId)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"suppliers": suppliers})
}
