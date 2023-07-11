package server

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://sayat:Sayat2005!@localhost:5432/word_rarity_db?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
