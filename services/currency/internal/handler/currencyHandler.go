package handler

import (
	"net/http"

	"github.com/marelinaa/currency-api/services/currency/internal/domain"
	"github.com/marelinaa/currency-api/services/currency/internal/service"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	currencyService *service.CurrencyService
}

func NewCurrencyHandler(currencyService *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{currencyService: currencyService}
}

func (h *CurrencyHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	// Protected currency routes (require authorization)
	currency := v1.Group("/currency") // todo: currency := v1.Group("/currency", middleware.Authorize())
	{
		currency.GET("/date/:date", h.GetCurrencyByDate)                   // todo: убрать date
		currency.GET("/history/:startDate/:endDate", h.GetCurrencyHistory) // todo: убрать history
	}

}

func (h *CurrencyHandler) GetCurrencyByDate(c *gin.Context) {
	date := c.Param("date")

	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrEmptyDate})

		return
	}

	rate, err := h.currencyService.GetCurrencyByDate(c, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"date": date, "rate": rate})
}

// GetCurrencyHistory retrieves the currency rate history for a specified period
func (h *CurrencyHandler) GetCurrencyHistory(c *gin.Context) {
	startDate := c.Query("start")
	endDate := c.Query("end")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrEmptyDate})

		return
	}

	history, err := h.currencyService.GetCurrencyHistory(c, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"history": history})
}
