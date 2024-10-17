package handler

import (
	"fmt"
	"github.com/marelinaa/currency-api/gateway/internal/domain"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurrencyByDate retrieves currency data for a specific date from the currency service
func (h *GatewayHandler) GetCurrencyByDate(c *gin.Context) {
	date := c.Param("date")

	url := fmt.Sprintf("http://%s/v1/currency/date/%s", h.currencyServiceURL, date)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch currency data"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(body)})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": domain.ErrReadingResponse})
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}

// GetCurrencyHistory retrieves historical currency data within a specified date range from the currency service
func (h *GatewayHandler) GetCurrencyHistory(c *gin.Context) {
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")

	url := fmt.Sprintf("http://%s/v1/currency/history/%s/%s", h.currencyServiceURL, startDate, endDate)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch currency history"})

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(body)})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": domain.ErrReadingResponse})
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}
