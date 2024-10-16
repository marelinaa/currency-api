package migrations

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

const migrationsPath = "/currency/migrations"

func RunMigrations(dsn string) (err error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Printf("start migrating database \n")
	return goose.Run("up", db, migrationsPath)
}
