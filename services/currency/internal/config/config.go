package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type WorkerConfig struct {
	ApiURL             string
	RunFetchingOnStart bool
	RunTime            time.Time
}

// AppConfig представляет общую конфигурацию для вашего приложения
type AppConfig struct {
	DatabaseURL string
	Worker      WorkerConfig
}

// Load загружает конфигурацию из переменных среды или по умолчанию
func Load() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found, using default values.")
	}

	return AppConfig{
		DatabaseURL: getEnv("DB_URL", "postgres://postgres:p1111@localhost:5432/currency_db"),
		Worker: WorkerConfig{
			ApiURL:             getEnv("CURRENCY_API_URL", "https://latest.currency-api.pages.dev/v1/currencies/rub.json"),
			RunFetchingOnStart: getEnvAsBool("WORKER_RUN_ON_START", true),
			RunTime:            parseRunTime(getEnv("WORKER_RUN_TIME", "00:00")),
		},
	}
}

// getEnv считывает переменные среды с указанием значения по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsBool возвращает значение переменной среды как bool
func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true"
	}
	return defaultValue
}

// parseRunTime преобразует строку в объект time.Time (только время)
func parseRunTime(timeString string) time.Time {
	layout := "15:04"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		log.Fatalf("invalid WORKER_RUN_TIME format. Expected HH:MM, got: %s", timeString)
	}
	return t
}
