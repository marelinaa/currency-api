package service

import (
	"fmt"
	"time"
)

// ValidateDate проверяет формат даты
func ValidateDate(dateStr string) (string, error) {
	layout := "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return "", err
	}

	if date.After(time.Now()) {
		return "", fmt.Errorf("date cannot be in the future")
	}

	return dateStr, nil
}
