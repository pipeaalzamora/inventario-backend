package recipe

type ProfileRecipe struct {
	ProfileName string   `json:"profileName" binding:"required"`
	Description string   `json:"description" binding:"required"`
	PowerIDs    []string `json:"powerIds"`
}
