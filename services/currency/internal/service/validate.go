package service

import (
	"fmt"
	"time"
)

// ValidateAndParseDate проверяет формат даты и преобразует строку в time.Time
func ValidateAndParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"

	// Пытаемся распарсить строку в time.Time
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err // Возвращаем ошибку, если не удалось распарсить
	}

	// Дополнительно, можно проверить, не является ли дата в будущем
	if date.After(time.Now()) {
		return time.Time{}, fmt.Errorf("date cannot be in the future")
	}

	return date, nil
}
