package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/marelinaa/currency-api/gateway/internal/domain"

	"github.com/gin-gonic/gin"
)

// SignIn handles user sign-in requests, validates credentials, and generates a token for authentication
func (h *GatewayHandler) SignIn(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error decoding request body"})
		return
	}
	log.Println(req)

	err := h.service.SignIn(req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.requestToken(req.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println(token)

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// requestToken makes a request to the authentication service to generate a token for a given login
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
