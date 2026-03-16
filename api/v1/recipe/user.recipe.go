package recipe

type UserRecipe struct {
	UserName     string   `json:"userName" binding:"required"`
	UserEmail    string   `json:"userEmail" binding:"required,email"`
	Description  string   `json:"description" binding:"required"`
	IsNewAccount bool     `json:"isNewAccount"`
	Password     string   `json:"password"`
	ProfileIDs   []string `json:"profileIds"`
}
