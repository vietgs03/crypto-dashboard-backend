package main

import (
	"context"
	"crypto-dashboard-backend/configs/development"
	"crypto-dashboard-backend/internal/market-data/application/service"
	"crypto-dashboard-backend/internal/market-data/domain/entity"
	"crypto-dashboard-backend/internal/market-data/infrastructure/coingecko"
	"crypto-dashboard-backend/internal/market-data/infrastructure/persistence/postgres"
	"crypto-dashboard-backend/pkg/core/server"
	database "crypto-dashboard-backend/pkg/database/postgres"
	"crypto-dashboard-backend/pkg/logger"
	"crypto-dashboard-backend/pkg/middleware"
	"crypto-dashboard-backend/pkg/migration"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MarketDataServer struct {
	*server.BaseServer
	db          *database.Connection
	coinService *service.CoinService
}

func NewMarketDataServer(cfg *server.Config) (*MarketDataServer, error) {
	baseServer, err := server.NewBaseServer(cfg)
	if err != nil {
		return nil, err
	}

	return &MarketDataServer{
		BaseServer: baseServer,
	}, nil
}

func (s *MarketDataServer) Initialize() error {
	// Load environment variables
	if err := development.Init(); err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	// Add request logging middleware
	s.App.Use(middleware.RequestLogger(s.Log))

	// Initialize migrate
	migrateConfig := migration.DefaultConfigMigrate()
	migration.RunDBMigration(migrateConfig.MigrationURL, migrateConfig.DBSource)

	// Initialize database
	dbConfig := database.DefaultConfig()
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		return err
	}
	s.db = db

	// Initialize services
	coinRepo := postgres.NewCoinRepository(s.db.DB())
	coinGeckoClient := coingecko.NewClient()
	s.coinService = service.NewCoinService(coinRepo, coinGeckoClient, s.Log)

	// Setup routes
	s.setupRoutes()

	return nil
}

func (s *MarketDataServer) setupRoutes() {
	v1 := s.App.Group("/api/v1")
	coins := v1.Group("/coins")

	// Get all coins
	//coins.Get("/", s.handleGetCoins)

	// Get specific coin by ID
	//coins.Get("/:id", s.handleGetCoin)

	// Manually trigger sync
	coins.Post("/sync", s.handleSyncCoins)

	// Get Bitcoin specific data
	coins.Get("/bitcoin", s.handleGetBitcoin)
}

func (s *MarketDataServer) handleSyncCoins(c *fiber.Ctx) error {
	ctx := c.UserContext()
	logger := s.Log.WithContext(ctx)

	logger.Info("starting coin sync")

	if err := s.coinService.FetchAndStoreCoins(ctx); err != nil {
		logger.Error("failed to sync coins", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to sync coins",
		})
	}

	logger.Info("coin sync completed successfully")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Coins synced successfully",
	})
}

func (s *MarketDataServer) handleGetBitcoin(c *fiber.Ctx) error {
	if s.coinService == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service not initialized",
		})
	}

	ctx := c.UserContext()
	logger := s.Log.WithContext(ctx)

	coins, err := s.coinService.GetCoins(ctx)
	if err != nil {
		logger.Error("failed to get coins", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get Bitcoin data",
		})
	}

	// Find Bitcoin in the coins list
	var bitcoin *entity.Coin
	for _, coin := range coins {
		if coin.ID == "bitcoin" {
			bitcoin = coin
			break
		}
	}

	if bitcoin == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Bitcoin data not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   bitcoin,
	})
}

func (s *MarketDataServer) Cleanup() error {
	return s.db.Close()
}

func main() {
	cfg := &server.Config{
		Port:        development.GetEnvInt("SERVER_PORT", 8080),
		Name:        "market-data-service",
		Environment: development.GetEnvStr("ENVIRONMENT", "development"),
		LogLevel:    development.GetEnvStr("LOG_LEVEL", "info"),
	}
	logConfig := logger.DefaultConfig()
	logConfig.ServiceName = "market-data-service"
	logConfig.Environment = development.GetEnvStr("ENVIRONMENT", "development")
	logConfig.Level = development.GetEnvStr("LOG_LEVEL", "info")
	logConfig.LogDir = development.GetEnvStr("LOG_DIR", "./logs")

	srv, err := NewMarketDataServer(cfg)
	if err != nil {
		fmt.Printf("Failed to create server: %v\n", err)
		os.Exit(1)
	}

	if err := srv.Initialize(); err != nil {
		srv.Log.Error(context.Background(), "failed to initialize server", err)
		os.Exit(1)
	}

	defer srv.Cleanup()

	if err := srv.Start(); err != nil {
		srv.Log.Error(context.Background(), "server error", err)
		os.Exit(1)
	}
}
