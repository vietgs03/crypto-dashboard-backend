package global

import (
	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Server settings.ServerSetting `mapstructure:"server"`
	Sql    settings.SQLSetting    `mapstructure:"sql"`
	Logger settings.LoggerSetting `mapstructure:"logger"`
	Redis  settings.CacheSetting  `mapstructure:"redis"`
	JWT    settings.JWTSetting    `mapstructure:"jwt"`
}

type AppServer struct {
	Logger *appLogger.Logger
	Config *Config
	SqlDB  *pgxpool.Pool
	Redis  *redis.Client
}
