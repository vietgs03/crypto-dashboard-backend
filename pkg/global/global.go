package global

import (
	"database/sql"

	"crypto-dashboard-backend/pkg/appLogger"
	"crypto-dashboard-backend/pkg/settings"

	"github.com/redis/go-redis/v9"
)

var (
	Config *settings.Config
	Logger *appLogger.Logger
	Rdb    *redis.Client
	Mdbc   *sql.DB
)
