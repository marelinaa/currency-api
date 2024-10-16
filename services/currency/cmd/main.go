package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/marelinaa/currency-api/currency/internal/config"
	"github.com/marelinaa/currency-api/currency/internal/handler"
	"github.com/marelinaa/currency-api/currency/internal/repository"
	"github.com/marelinaa/currency-api/currency/internal/service"
	"github.com/marelinaa/currency-api/currency/migrations"

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

	apiPort := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("Starting server on %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}
