package domain

import "errors"

var (
	ErrRateNotFound        = errors.New("the currency rate for given date was not found")
	ErrEmptyDate           = errors.New("date query parameter can not be empty")
	ErrRateHistoryNotFound = errors.New("the currency rate history for given range was not found")
	ErrSavingCurrencyRate  = errors.New("error saving currency rate by worker")
	ErrParsingDate         = errors.New("invalid date format, must be yyyy-mm-dd")
	ErrFutureDate          = errors.New("date cannot be in the future")
)
