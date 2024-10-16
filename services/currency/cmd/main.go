package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/marelinaa/currency-api/services/currency/internal/config"
	"github.com/marelinaa/currency-api/services/currency/internal/handler"
	"github.com/marelinaa/currency-api/services/currency/internal/repository"
	"github.com/marelinaa/currency-api/services/currency/internal/service"
	"github.com/marelinaa/currency-api/services/currency/migrations"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	err := migrations.RunMigrations(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	repo := repository.NewRepository(db)

	currencyService := service.NewCurrencyService(repo)
	currencyHandler := handler.NewCurrencyHandler(currencyService)

	router := gin.Default()
	currencyHandler.DefineRoutes(router)

	worker := service.NewWorker(currencyService, cfg.Worker)
	worker.Start()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
