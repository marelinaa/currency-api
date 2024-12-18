package handler

import (
	"errors"
	"net/http"

	"github.com/marelinaa/currency-api/currency/internal/domain"
	"github.com/marelinaa/currency-api/currency/internal/service"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	currencyService *service.CurrencyService
}

func NewCurrencyHandler(currencyService *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{currencyService: currencyService}
}

// DefineRoutes defines routes for handling currency-related API endpoints
func (h *CurrencyHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	currency := v1.Group("/currency")
	{
		currency.GET("/date", h.GetCurrencyByDate)
		currency.GET("/history", h.GetCurrencyHistory)
	}

}

// GetCurrencyByDate retrieves the currency rate for a specific date from the currency service
func (h *CurrencyHandler) GetCurrencyByDate(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrEmptyDate.Error()})

		return
	}

	rate, err := h.currencyService.GetCurrencyByDate(c, date)
	if err != nil {
		if errors.Is(err, domain.ErrFutureDate) || errors.Is(err, domain.ErrParsingDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"date": date, "rate": rate})
}

// GetCurrencyHistory retrieves the currency rate history for a specified period from the currency service
func (h *CurrencyHandler) GetCurrencyHistory(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrEmptyDate.Error()})

		return
	}

	history, err := h.currencyService.GetCurrencyHistory(c, startDate, endDate)
	if err != nil {
		if errors.Is(err, domain.ErrFutureDate) || errors.Is(err, domain.ErrParsingDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"history": history})
}
