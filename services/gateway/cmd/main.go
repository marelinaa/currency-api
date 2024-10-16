package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/marelinaa/currency-api/gateway/internal/config"
	"github.com/marelinaa/currency-api/gateway/internal/handler"
	"github.com/marelinaa/currency-api/gateway/internal/service"

	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	users := map[string]string{
		"Artyom": "12345",
		"Elya":   "54321",
	}
	gatewayService := service.NewGatewayService(users)

	gatewayHandler := handler.NewGatewayHandler(gatewayService, cfg.AuthServiceURL, cfg.CurrencyServiceURL)

	router := gin.Default()
	gatewayHandler.DefineRoutes(router)

	apiPort := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("Starting server on %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}
