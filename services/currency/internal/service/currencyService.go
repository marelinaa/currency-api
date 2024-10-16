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

func (s *CurrencyService) SaveCurrencyData(data domain.CurrencyResponse) error {
	date, err := ValidateDate(data.Date)
	if err != nil {
		return err
	}

	var currency domain.CurrencyData

	currency = domain.CurrencyData{
		Date: date,
		Rate: data.Rub.Eur,
	}

	return s.repo.Save(currency)
}

func (s *CurrencyService) GetCurrencyByDate(ctx context.Context, dateStr string) (domain.CurrencyData, error) {
	date, err := ValidateDate(dateStr)
	if err != nil {
		return domain.CurrencyData{}, err
	}

	return s.repo.FindByDate(ctx, date)
}

func (s *CurrencyService) GetCurrencyHistory(ctx context.Context, startDateStr, endDateStr string) ([]domain.CurrencyData, error) {
	startDate, err := ValidateDate(startDateStr)
	if err != nil {
		return []domain.CurrencyData{}, err
	}

	endDate, err := ValidateDate(endDateStr)
	if err != nil {
		return []domain.CurrencyData{}, err
	}

	return s.repo.FindInRange(ctx, startDate, endDate)
}
