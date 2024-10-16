package migrations

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const migrationsPath = "services/currency/migrations"

func RunMigrations(dsn string) (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Printf("start migrating database \n")
	return goose.Run("up", db, migrationsPath)
}
