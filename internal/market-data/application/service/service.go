package service

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/repository"
	"crypto-dashboard-backend/internal/market-data/infrastructure/coingecko"

	"time"
)

type CoinService struct {
	repo   repository.CoinRepository
	client *coingecko.Client
}

func NewCoinService(repo repository.CoinRepository, client *coingecko.Client) *CoinService {
	return &CoinService{
		repo:   repo,
		client: client,
	}
}

func (s *CoinService) FetchAndStoreCoins(ctx context.Context) error {
	coins, err := s.client.GetCoins(ctx)
	if err != nil {
		return err
	}

	// Update timestamp
	now := time.Now()
	for _, coin := range coins {
		coin.UpdatedAt = now
	}

	return s.repo.SaveCoins(ctx, coins)
}
