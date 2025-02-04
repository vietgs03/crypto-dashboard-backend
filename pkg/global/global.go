package global

import (
	"crypto-dashboard-backend/pkg/appLogger"
	"crypto-dashboard-backend/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AppServer struct {
	Logger *appLogger.Logger
	Config *settings.Config
	DB     *pgxpool.Pool
	Redis  *redis.Client
}
