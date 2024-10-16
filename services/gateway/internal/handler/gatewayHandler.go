package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/marelinaa/currency-api/gateway/internal/domain"
	"github.com/marelinaa/currency-api/gateway/internal/service"

	"github.com/gin-gonic/gin"
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

func (h *GatewayHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/sign-in", h.SignIn)
	currency := v1.Group("/currency", h.Authorize())
	{
		currency.GET("/date/:date", h.GetCurrencyByDate)
		currency.GET("/history/:startDate/:endDate", h.GetCurrencyHistory)
	}
}

func (h *GatewayHandler) SignIn(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error decoding request body"})
		return
	}
	log.Println(req)

	// Проверка пользовательских данных через GatewayService
	err := h.service.SignIn(req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Генерация токена через внешний сервис авторизации
	token, err := h.requestToken(req.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println(token)

	// Возвращаем токен пользователю
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *GatewayHandler) requestToken(login string) (string, error) {
	url := fmt.Sprintf("http://%s/generate?login=%s", h.authServiceURL, login)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest: %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()
	log.Println(resp)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to generate token: %s", string(body))
	}

	tokenResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll: %w", err)
	}

	return string(tokenResp), nil
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response"})
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response"})
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}

func (h *GatewayHandler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		err := h.validateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (h *GatewayHandler) validateToken(token string) error {
	url := fmt.Sprintf("http://%s/validate", h.authServiceURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create validation request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send validation request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token validation failed with status code: %d", resp.StatusCode)
	}

	return nil
}
