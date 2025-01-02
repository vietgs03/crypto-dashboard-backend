package server

import (
	"context"
	"crypto-dashboard-backend/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type BaseServer struct {
	App  *fiber.App
	Log  *logger.Logger
	Port int
	Name string
}

type Config struct {
	Port            int
	Name            string
	Environment     string
	LogLevel        string
	ShutdownTimeout time.Duration
}

func NewBaseServer(cfg *Config) (*BaseServer, error) {
	// Initialize logger
	log, err := logger.NewLogger(&logger.LogConfig{
		Level:       cfg.LogLevel,
		ServiceName: cfg.Name,
		Environment: cfg.Environment,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	// Configure Fiber app
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Add base middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return &BaseServer{
		App:  app,
		Log:  log,
		Port: cfg.Port,
		Name: cfg.Name,
	}, nil
}

func (s *BaseServer) Start() error {
	// Graceful shutdown setup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	serverErr := make(chan error, 1)
	go func() {
		addr := fmt.Sprintf(":%d", s.Port)
		if err := s.App.Listen(addr); err != nil {
			serverErr <- err
		}
	}()

	// Wait for interrupt signal or server error
	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-sigChan:
		s.Log.Info("received signal", zap.String("signal", sig.String()))
	}

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := s.App.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("shutdown error: %w", err)
	}

	s.Log.Info("graceful shutdown completed")
	return nil
}
