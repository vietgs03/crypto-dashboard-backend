package service

import (
	"context"
	"crypto-dashboard-backend/pkg/logger"
	"time"

	"go.uber.org/zap"
)

type CoinSyncService struct {
	coinService  *CoinService
	log          *logger.Logger
	syncInterval time.Duration
	stopChan     chan struct{}
}

func NewCoinSyncService(coinService *CoinService, log *logger.Logger) *CoinSyncService {
	return &CoinSyncService{
		coinService:  coinService,
		log:          log,
		syncInterval: 1 * time.Minute,
		stopChan:     make(chan struct{}),
	}
}

func (s *CoinSyncService) Start(ctx context.Context) {
	ticket := time.NewTicker(s.syncInterval)
	defer ticket.Stop()

	// sync lần đầu ngay khi khởi động.
	logger := s.log.WithContext(ctx)
	if err := s.coinService.FetchAndStoreCoins(ctx); err != nil {
		logger.Error("failed to sync coins", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			logger.Info("context done, stopping coin sync service")
		case <-s.stopChan:
			logger.Info("received stop signal, stopping coin sync service")
			return
		case <-ticket.C:
			if err := s.coinService.FetchAndStoreCoins(ctx); err != nil {
				logger.Error("failed to sync coins", zap.Error(err))
				continue
			}
		}
	}
}

func (s *CoinSyncService) Stop() {
	close(s.stopChan)
}
