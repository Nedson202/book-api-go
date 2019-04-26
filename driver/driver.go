package driver

import (
	"log"
	"os"

	"database/sql"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// SetupDatabaseConnection with sql and postgres
func SetupDatabaseConnection() *sql.DB {
	dbURL, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", dbURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
