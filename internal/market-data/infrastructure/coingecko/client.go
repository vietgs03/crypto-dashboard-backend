package coingecko

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/entity"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://api.coingecko.com/api/v3",
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		apiKey: os.Getenv("COINGECKO_API_KEY"),
	}
}

type CoinGeckoResponse struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	CurrentPrice      float64 `json:"current_price"`
	MarketCap         float64 `json:"market_cap"`
	TotalVolume       float64 `json:"total_volume"`
	PriceChangeDay    float64 `json:"price_change_percentage_24h"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	ATH               float64 `json:"ath"`
	ATHDate           string  `json:"ath_date"`
}

func (c *Client) GetCoins(ctx context.Context) ([]*entity.Coin, error) {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Set("x-cg-pro-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geckoCoins []CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geckoCoins); err != nil {
		return nil, err
	}

	coins := make([]*entity.Coin, len(geckoCoins))
	for i, gc := range geckoCoins {
		athDate, _ := time.Parse(time.RFC3339, gc.ATHDate)
		coins[i] = &entity.Coin{
			ID:                gc.ID,
			Symbol:            gc.Symbol,
			Name:              gc.Name,
			Price:             gc.CurrentPrice,
			MarketCap:         gc.MarketCap,
			Volume24h:         gc.TotalVolume,
			PriceChange:       gc.PriceChangeDay,
			CirculatingSupply: gc.CirculatingSupply,
			TotalSupply:       gc.TotalSupply,
			ATH:               gc.ATH,
			ATHDate:           athDate,
		}
	}

	return coins, nil
}
