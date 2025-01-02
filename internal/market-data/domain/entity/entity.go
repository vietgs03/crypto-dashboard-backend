package entity

import "time"

type Coin struct {
	ID          string    `json:"id"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	MarketCap   float64   `json:"market_cap"`
	Volume24h   float64   `json:"volume_24h"`
	PriceChange float64   `json:"price_change_24h"`
	UpdatedAt   time.Time `json:"updated_at"`
}
