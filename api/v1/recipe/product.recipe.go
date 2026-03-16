package recipe

import (
	"mime/multipart"
	"sofia-backend/domain/models"
)

// RecipeProduct - Bypass directo al modelo ModelProduct
type RecipeProduct = models.ModelProduct

type RecipeProductInput struct {
	Name          string                `form:"name" binding:"required" errMsg:"El nombre del producto es obligatorio"`
	Description   *string               `form:"description" binding:"required" errMsg:"La descripción es obligatoria"`
	CostEstimated float32               `form:"costEstimated"`
	Image         *multipart.FileHeader `form:"image"`
	CategoryIds   []int64               `form:"categoryIds"`
	Codes         string                `form:"codes" binding:"required" errMsg:"Los códigos son obligatorios"`
	IsNewImage    bool                  `form:"isNewImage"`
	CodesList     []RecipeProductCode   `form:"-"`
}

type RecipeProductCode struct {
	ID    int    `json:"id" binding:"required" errMsg:"El ID del código es obligatorio"`
	Value string `json:"value" binding:"required" errMsg:"El valor del código es obligatorio"`
}

type RecipeProductCategory struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	//Available   *bool   `json:"available"`
}
