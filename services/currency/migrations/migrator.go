package migrations

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const migrationsPath = "."

func RunMigrations(dsn string) (err error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Printf("start migrating database \n")
	return goose.Run("up", db, migrationsPath)
}
