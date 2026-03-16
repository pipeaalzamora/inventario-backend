package dto

type CategoryAccountWithPowersDTO struct {
	ID           string            `json:"id"`
	CategoryName string            `json:"categoryName"`
	Description  string            `json:"description"`
	Ownable      bool              `json:"ownable"`
	Powers       []PowerAccountDTO `json:"powers"`
}

type PowerAccountDTO struct {
	ID          string `json:"id"`
	PowerName   string `json:"powerName"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}
