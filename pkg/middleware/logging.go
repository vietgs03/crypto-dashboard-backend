package middleware

import (
	"context"
	"crypto-dashboard-backend/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(log *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cid := uuid.New().String()
		rid := uuid.New().String()

		ctx := context.WithValue(c.Context(), logger.CorrelationIDKey, cid)
		ctx = context.WithValue(ctx, logger.RequestIDKey, rid)

		c.SetUserContext(ctx)
		start := time.Now()

		// log request started
		log.WithContext(ctx).Info("request_started",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("request_id", rid),
			zap.Time("start_time", time.Now()),
		)
		// process request
		err := c.Next()

		// calculate duration

		duration := time.Since(start)

		log.WithContext(ctx).Info("request_completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("request_id", rid),
			zap.Int("status_code", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)
		return err
	}
}
