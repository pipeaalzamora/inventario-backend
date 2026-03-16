package recipe

type RecoveryPasswordRecipe struct {
	UserEmail string `json:"userEmail" binding:"required,email"`
}

type ChangePasswordWithCodeRecipe struct {
	UserEmail       string `json:"userEmail" binding:"required,email"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8"`
}

type VerifyCodeRecipe struct {
	UserEmail string `json:"userEmail" binding:"required,email"`
	Code      string `json:"code" binding:"required"`
}

type ChangePasswordRecipe struct {
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	OldPassword     string `json:"oldPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8"`
}

type LoginRecipe struct {
	UserEmail string `json:"userEmail" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}
