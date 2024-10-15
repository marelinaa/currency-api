package service

import (
	"context"
	"github.com/marelinaa/currency-api/services/currency/internal/domain"
	"github.com/marelinaa/currency-api/services/currency/internal/repository"
)

type CurrencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) SaveCurrencyData(data domain.CurrencyData) error {
	return s.repo.Save(data)
}

func (s *CurrencyService) GetCurrencyByDate(ctx context.Context, dateStr string) (domain.CurrencyData, error) {
	date, err := ValidateAndParseDate(dateStr)
	if err != nil {
		return domain.CurrencyData{}, err
	}

	return s.repo.FindByDate(ctx, date)
}

func (s *CurrencyService) GetCurrencyHistory(ctx context.Context, startDateStr, endDateStr string) ([]domain.CurrencyData, error) {
	startDate, err := ValidateAndParseDate(startDateStr)
	if err != nil {
		return []domain.CurrencyData{}, err
	}

	endDate, err := ValidateAndParseDate(startDateStr)
	if err != nil {
		return []domain.CurrencyData{}, err
	}

	return s.repo.FindInRange(ctx, startDate, endDate)
}
