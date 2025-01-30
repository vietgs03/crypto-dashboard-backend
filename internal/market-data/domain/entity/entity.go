package entity

import "time"

type Coin struct {
	ID                string    `json:"id"`
	Symbol            string    `json:"symbol"`
	Name              string    `json:"name"`
	Price             float64   `json:"price"`
	MarketCap         float64   `json:"market_cap"`
	Volume24h         float64   `json:"volume_24h"`
	PriceChange       float64   `json:"price_change_24h"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	ATH               float64   `json:"ath"`
	ATHDate           time.Time `json:"ath_date"`
	UpdatedAt         time.Time `json:"updated_at"`
}
