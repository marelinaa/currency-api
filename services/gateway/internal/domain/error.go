package domain

import "errors"

var (
	ErrEmptyInput         = errors.New("given login or password is empty")
	ErrInvalidCredentials = errors.New("given credentials are invalid")
	ErrReadingResponse    = errors.New("failed to read response")
	ErrEmptyDate          = errors.New("date query parameter can not be empty")
)
