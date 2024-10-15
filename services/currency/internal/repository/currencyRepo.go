package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/marelinaa/currency-api/services/currency/internal/domain"
)

// PostgresCurrencyRepository Пример реализации с использованием PostgreSQL
type PostgresCurrencyRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *PostgresCurrencyRepository {

	return &PostgresCurrencyRepository{
		db: db,
	}
}

type CurrencyRepository interface {
	Save(data domain.CurrencyData) error
	FindByDate(ctx context.Context, date time.Time) (domain.CurrencyData, error)
	FindInRange(ctx context.Context, startDate, endDate time.Time) ([]domain.CurrencyData, error)
}

func (r *PostgresCurrencyRepository) Save(data domain.CurrencyData) error {
	return nil
}

func (r *PostgresCurrencyRepository) FindByDate(ctx context.Context, date time.Time) (domain.CurrencyData, error) {

	var currency domain.CurrencyData
	querySelect := `SELECT date, rate 
					FROM currency_rate
					WHERE date=$1;`
	row := r.db.QueryRowContext(ctx, querySelect, date)
	err := row.Scan(&currency.Date, &currency.Rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.CurrencyData{}, domain.ErrRateNotFound
		}

		return domain.CurrencyData{}, err
	}

	return currency, nil
}

func (r *PostgresCurrencyRepository) FindInRange(ctx context.Context, startDate, endDate time.Time) ([]domain.CurrencyData, error) {
	var history []domain.CurrencyData

	querySelect := `SELECT date, rate 
					FROM currency_rate
					WHERE date BETWEEN $1 AND $2 
					ORDER BY date;`

	rows, err := r.db.QueryContext(ctx, querySelect, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency domain.CurrencyData
		if err := rows.Scan(&currency.Date, &currency.Rate); err != nil {
			return nil, err
		}
		history = append(history, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(history) == 0 {
		return nil, domain.ErrRateHistoryNotFound
	}

	return history, nil
}
