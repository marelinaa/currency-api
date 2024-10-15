package domain

import "time"

type CurrencyData struct {
	Date time.Time `json:"date"`
	Rate float64   `json:"rate"`
}
