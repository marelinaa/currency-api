package domain

import "errors"

var (
	ErrEmptyInput         = errors.New("given login or password is empty")
	ErrInvalidCredentials = errors.New("given credentials are invalid")
)
