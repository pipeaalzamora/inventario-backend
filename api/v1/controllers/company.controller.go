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
	r.GET(":id/suppliers", c.getCompanySuppliers)

	r.POST("", c.createCompany)
	r.POST(":id", c.updateCompany)
	r.POST(":id/suppliers", c.assignSuppliersToCompany)
	r.DELETE(":id/suppliers/:supplierId", c.unassignSupplierFromCompany)

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
