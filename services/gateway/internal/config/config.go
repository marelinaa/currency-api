package config

import (
	"os"
)

type AppConfig struct {
	APIPort            string
	AuthServiceURL     string
	CurrencyServiceURL string
}

func Load() AppConfig {
	return AppConfig{
		APIPort:            getEnv("API_PORT", "8088"),
		AuthServiceURL:     getEnv("AUTH_URL", "localhost:8082"),
		CurrencyServiceURL: getEnv("CURRENCY_URL", "localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
