package database

import (
	"fmt"
	"sofia-backend/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(s *config.Config) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Postgres.Host, s.Postgres.Port, s.Postgres.DbUser, s.Postgres.DbPass, s.Postgres.DbName)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
