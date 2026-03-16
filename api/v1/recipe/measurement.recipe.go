package recipe

type MeasurementRecipe struct {
	Name             string  `json:"name" binding:"required" errMsg:"El nombre es obligatorio"`
	Description      string  `json:"description" binding:"required" errMsg:"La descripción es obligatoria"`
	Abbreviation     string  `json:"abbreviation" binding:"required" errMsg:"La abreviación es obligatoria"`
	BaseUnitId       int     `json:"baseUnitId" binding:"required" errMsg:"El ID de la unidad base es obligatorio"`
	ConversionFactor float32 `json:"conversionFactor" binding:"required" errMsg:"El factor de conversión es obligatorio"`
}
