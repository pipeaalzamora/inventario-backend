package entities

type PowerAccount struct {
	ID          string `db:"id"`
	PowerName   string `db:"power_name"`
	DisplayName string `db:"power_display"`
	Description string `db:"description"`
	CategoryID  string `db:"power_account_category_id"` // FK a categoría
}

type PowerAccountCategory struct {
	ID           string `db:"id"`
	CategoryName string `db:"category_name"`
	Description  string `db:"description"`
	Ownable      bool   `db:"ownable"`
}

type PowerAccountProfile struct {
	ID          string `db:"id"`
	PowerName   string `db:"power_name"`
	DisplayName string `db:"power_display"`
	Description string `db:"description"`
	CategoryID  string `db:"power_account_category_id"` // FK a categoría
	ProfileID   string `db:"profile_id"`
}
