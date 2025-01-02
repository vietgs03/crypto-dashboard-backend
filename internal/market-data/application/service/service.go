package service

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/repository"
	"crypto-dashboard-backend/internal/market-data/infrastructure/coingecko"
	"crypto-dashboard-backend/pkg/logger"

	"time"

	"go.uber.org/zap"
)

type CoinService struct {
	repo   repository.CoinRepository
	client *coingecko.Client
	log    *logger.Logger
}

func NewCoinService(repo repository.CoinRepository, client *coingecko.Client, log *logger.Logger) *CoinService {
	return &CoinService{
		repo:   repo,
		client: client,
		log:    log,
	}
}

func (s *CoinService) FetchAndStoreCoins(ctx context.Context) error {
	logger := s.log.WithContext(ctx)

	logger.Info("fetching coins from coingecko")

	coins, err := s.client.GetCoins(ctx)
	if err != nil {
		logger.Error("failed to fetch coins",
			zap.Error(err),
		)
		return err
	}

	logger.Info("successfully fetched coins",
		zap.Int("count", len(coins)),
	)

	// Update timestamp
	now := time.Now()
	for _, coin := range coins {
		coin.UpdatedAt = now
	}

	if err := s.repo.SaveCoins(ctx, coins); err != nil {
		logger.Error("failed to save coins",
			zap.Error(err),
		)
		return err
	}

	logger.Info("successfully stored coins",
		zap.Int("count", len(coins)),
	)

	return nil
}
