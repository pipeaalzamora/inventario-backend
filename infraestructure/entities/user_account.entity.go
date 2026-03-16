package entities

import (
	"time"
)

type UserAccount struct {
	ID           string     `db:"id"`
	UserName     string     `db:"user_name"`
	UserEmail    string     `db:"user_email"`
	Description  string     `db:"description"`
	UserPassword string     `db:"user_password"`
	Available    bool       `db:"available"`
	IsNewAccount bool       `db:"is_new_account"`
	DeletedAt    *time.Time `db:"deleted_at"` // Nullable field for soft delete
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
}

type InputUserAccount struct {
	UserName     string   `json:"user_name"`
	UserEmail    string   `json:"user_email"`
	Description  string   `json:"description"`
	UserPassword string   `json:"user_password"`
	IsNewAccount bool     `json:"is_new_account"`
	ProfileIDs   []string `json:"profilesId"`
}
