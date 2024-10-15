package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/marelinaa/currency-api/services/currency/migrations"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/marelinaa/currency-api/services/currency/internal/config"
	"github.com/marelinaa/currency-api/services/currency/internal/handler"
	"github.com/marelinaa/currency-api/services/currency/internal/repository"
	"github.com/marelinaa/currency-api/services/currency/internal/service"
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

	// Регистрация маршрутов
	router := gin.Default()
	currencyHandler.DefineRoutes(router)

	// Инициализация и запуск воркера
	worker := service.NewWorker(currencyService, cfg.Worker)
	worker.Start()

	// Запуск HTTP-сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
