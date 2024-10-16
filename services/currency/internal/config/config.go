package config

import (
	"log"
	"os"
	"time"
)

type WorkerConfig struct {
	ApiURL             string
	RunFetchingOnStart bool
	RunTime            time.Time
}

type AppConfig struct {
	DatabaseURL string
	APIPort     string
	Worker      WorkerConfig
}

func Load() AppConfig {
	return AppConfig{
		DatabaseURL: getEnv("DB_URL", "postgres://postgres:p1111@localhost:5432/currency_db?sslmode=disable"),
		APIPort:     getEnv("API_PORT", "8080"),
		Worker: WorkerConfig{
			ApiURL:             getEnv("CURRENCY_API_URL", "https://latest.currency-api.pages.dev/v1/currencies/rub.json"),
			RunFetchingOnStart: getEnvAsBool("WORKER_RUN_ON_START", true),
			RunTime:            parseRunTime(getEnv("WORKER_RUN_TIME", "00:00")),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true"
	}
	return defaultValue
}

func parseRunTime(timeString string) time.Time {
	layout := "15:04"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		log.Fatalf("invalid WORKER_RUN_TIME format. Expected HH:MM, got: %s", timeString)
	}
	return t
}
