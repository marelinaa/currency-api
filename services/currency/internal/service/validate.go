package service

import (
	"fmt"
	"github.com/marelinaa/currency-api/currency/internal/domain"
	"log"
	"time"
)

// ValidateDate validates and parses a date string in the format "yyyy-mm-dd" and ensures it is not in the future
func ValidateDate(dateStr string) (string, error) {
	layout := "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Printf("error while parsing date: %v", err)

		return "", domain.ErrParsingDate
	}

	if date.After(time.Now()) {
		return "", fmt.Errorf("date cannot be in the future")
	}

	return dateStr, nil
}

// ValidatePeriod validates the start and end dates in the format "yyyy-mm-dd" for a period and ensures the start date is not later than the end date
func ValidatePeriod(startDate, endDate string) error {
	layout := "2006-01-02"
	startDateTime, err := time.Parse(layout, startDate)
	if err != nil {
		log.Printf("error while parsing date: %v", err)
		return domain.ErrParsingDate
	}

	endDateTime, err := time.Parse(layout, endDate)
	if err != nil {
		log.Printf("error while parsing date: %v", err)
		return domain.ErrParsingDate
	}

	if startDateTime.After(endDateTime) {
		return fmt.Errorf("end date can not be earlier, than start date")
	}

	return nil
}
