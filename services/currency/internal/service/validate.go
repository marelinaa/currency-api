package service

import (
	"fmt"
	"time"
)

// ValidateDate
func ValidateDate(dateStr string) (string, error) {
	layout := "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return "", fmt.Errorf("error while parsing date: %v", err)
	}

	if date.After(time.Now()) {
		return "", fmt.Errorf("date cannot be in the future")
	}

	return dateStr, nil
}

func ValidatePeriod(startDate, endDate string) error {
	layout := "2006-01-02"
	startDateTime, err := time.Parse(layout, startDate)
	if err != nil {
		return fmt.Errorf("error while parsing date: %v", err)
	}

	endDateTime, err := time.Parse(layout, endDate)
	if err != nil {
		return fmt.Errorf("error while parsing date: %v", err)
	}

	if startDateTime.After(endDateTime) {
		return fmt.Errorf("startDate can not be earlier, than endDate")
	}

	return nil
}
