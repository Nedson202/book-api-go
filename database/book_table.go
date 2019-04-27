package database

import (
	"database/sql"

	"github.com/nedson202/book-api-go/config"
)

// CreateUserTable with psql
func CreateUserTable(db *sql.DB) {
	var query = `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			username text NOT NULL,
			email text NOT NULL,
			password text NOT NULL,
			role text NOT NULL,
			created_at timestamp with time zone DEFAULT current_timestamp NOT NULL,
			updated_at timestamp with time zone DEFAULT current_timestamp NOT NULL,
			deleted_at timestamp with time zone DEFAULT current_timestamp
		)
	`
	_, err := db.Query(query)
	config.LogFatal(err)

	return
}

// DropUserTable with psql
func DropUserTable(db *sql.DB) {
	var query = `
		DROP TABLE users
	`
	_, err := db.Query(query)
	config.LogFatal(err)

	return
}
