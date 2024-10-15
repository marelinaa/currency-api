package domain

import "errors"

var (
	ErrRateNotFound        = errors.New("the currency rate for given date was not found")
	ErrEmptyDate           = errors.New("date can not be empty")
	ErrRateHistoryNotFound = errors.New("the currency rate history for given range was not found")
)
