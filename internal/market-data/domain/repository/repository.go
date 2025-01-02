package repository

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/entity"
)

type CoinRepository interface {
	SaveCoins(ctx context.Context, coint []*entity.Coin) error
	GetCoins(ctx context.Context) ([]*entity.Coin, error)
}
