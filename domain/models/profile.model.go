package models

type ProfileAccountModel struct {
	ID          string `json:"id"`
	ProfileName string `json:"profileName"`
	Description string `json:"description"`
}

type ProfileAccountUserAccountModel struct {
	ID          string `json:"id"`
	ProfileName string `json:"profileName"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

type PowerAccountProfileModel struct {
	ID          string `json:"id"`
	PowerName   string `json:"powerName"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	CategoryID  string `json:"categoryId"`
	ProfileID   string `json:"profileId"`
}

type PowerAccountCategoryModel struct {
	ID           string `json:"id"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
	Ownable      bool   `json:"ownable"`
}

type PowerAccountModel struct {
	ID          string `json:"id"`
	PowerName   string `json:"powerName"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	CategoryID  string `json:"categoryId"`
}
