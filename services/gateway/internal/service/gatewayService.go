package service

import (
	"github.com/marelinaa/currency-api/gateway/internal/domain"
)

type GatewayService struct {
	users map[string]string
}

func NewGatewayService(users map[string]string) *GatewayService {
	gatewayService := &GatewayService{
		users: users,
	}

	return gatewayService
}

func (s *GatewayService) SignIn(userSignIn domain.User) error {
	if userSignIn.Login == "" || userSignIn.Password == "" {
		return domain.ErrEmptyInput
	}

	password, ok := s.users[userSignIn.Login]
	if !ok || password != userSignIn.Password {
		return domain.ErrInvalidCredentials
	}

	return nil
}
