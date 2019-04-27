package database

import (
	"database/sql"
)

// MigrateDatabaseTables
func MigrateDatabaseTables(db *sql.DB) {
	CreateBookTable(db)
	CreateUserTable(db)
}