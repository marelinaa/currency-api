package domain

type CurrencyData struct {
	Date string  `json:"date"`
	Rate float64 `json:"eur"`
}

type CurrencyResponse struct {
	Date string `json:"date"`
	Rub  struct {
		Eur float64 `json:"eur"`
	} `json:"rub"`
}
