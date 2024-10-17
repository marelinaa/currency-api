package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marelinaa/currency-api/gateway/internal/service"
)

type GatewayHandler struct {
	service            *service.GatewayService
	authServiceURL     string
	currencyServiceURL string
}

func NewGatewayHandler(srv *service.GatewayService, authServiceURL, currencyServiceURL string) *GatewayHandler {
	return &GatewayHandler{
		service:            srv,
		authServiceURL:     authServiceURL,
		currencyServiceURL: currencyServiceURL,
	}
}

// DefineRoutes defines routes for handling different API endpoints
func (h *GatewayHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/sign-in", h.SignIn)

	currency := v1.Group("/currency", h.Authorize())
	{
		currency.GET("/date/:date", h.GetCurrencyByDate)
		currency.GET("/history/:startDate/:endDate", h.GetCurrencyHistory)
	}
}
