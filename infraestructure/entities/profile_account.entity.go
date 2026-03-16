package entities

type ProfileAccount struct {
	ID          string `db:"id"`
	ProfileName string `db:"profile_name"`
	Description string `db:"description"`
}

type ProfileAccountUserAccount struct {
	ID          string `db:"id"`
	ProfileName string `db:"profile_name"`
	Description string `db:"description"`
	UserID      string `db:"user_id"`
}
