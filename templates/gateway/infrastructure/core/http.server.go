package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"crypto-dashboard/gw-example/Interfaces/middlewares"
	"crypto-dashboard/gw-example/infrastructure/global"
	"crypto-dashboard/pkg/response"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"

	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/zap"
)

type httpServer struct {
	App  *fiber.App
	Port int
	Name string
}

// TODO this core will not in package
func NewHttpServer() (*httpServer, *response.AppError) {
	// Configure Fiber app
	defer global.Log.Sync()
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Use(helmet.New(helmet.Config{
		CrossOriginEmbedderPolicy: "false",
		ContentSecurityPolicy:     "false",
		CrossOriginOpenerPolicy:   "cross-origin",
		CrossOriginResourcePolicy: "cross-origin",
	}), cors.New(cors.Config{
		AllowOrigins:     strings.Split(global.AppConfig.AllowOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}),
		recover.New(),
		app.Use(requestid.New()),
		fiberzap.New(fiberzap.Config{
			Logger: global.Log.Logger,
		}),
		middlewares.ReqContextHandler,
		middlewares.LoggingInterceptor,
	)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("I'm good!")
	})

	return &httpServer{
		App:  app,
		Port: global.AppConfig.Server.Port,
		Name: global.AppConfig.Server.ServiceName,
	}, nil
}

func (s *httpServer) Start() *response.AppError {
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
		return response.ServerError(fmt.Sprintf("server error: %v", err.Error()))
	case sig := <-sigChan:
		global.Log.Info("received signal", zap.String("signal", sig.String()))
	}

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := s.App.ShutdownWithContext(ctx); err != nil {
		return response.ServerError(fmt.Sprintf("shutdown error: %v", err.Error()))
	}

	global.Log.Info("graceful shutdown completed", nil)
	return nil
}
