package database

import (
	"database/sql"

	"github.com/nedson202/book-api-go/config"
)

// CreateBookTable with psql
func CreateBookTable(db *sql.DB) {
	var query = `
		CREATE TABLE IF NOT EXISTS books (
			id serial PRIMARY KEY,
			title text NOT NULL,
			author text NOT NULL,
			year text NOT NULL,
			created_at timestamp with time zone DEFAULT current_timestamp NOT NULL,
			updated_at timestamp with time zone DEFAULT current_timestamp NOT NULL,
			deleted_at timestamp with time zone DEFAULT current_timestamp
		)
	`
	_, err := db.Query(query)
	config.LogFatal(err)

	return
}

// DropBookTable with psql
func DropBookTable(db *sql.DB) {
	var query = `
		DROP TABLE books
	`
	_, err := db.Query(query)
	config.LogFatal(err)

	return
}
