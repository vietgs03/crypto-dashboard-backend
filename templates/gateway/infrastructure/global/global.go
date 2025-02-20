package global

import (
	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/database/caching"
	database "crypto-dashboard/pkg/database/postgres"
	"crypto-dashboard/pkg/settings"

	"github.com/gofiber/fiber/v3"
)

type (
	ServerUrls struct {
		User          string `mapstructure:"USER_URL"`
		MarketData    string `mapstructure:"MARKET_DATA_URL"`
		WhaleTracking string `mapstructure:"WHALE_TRACKING_URL"`
		Portfolio     string `mapstructure:"PORTFOLIO_URL"`
		Notification  string `mapstructure:"NOTIFICATION_URL"`
	}
	Config struct {
		AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`
		Server       *settings.ServerSetting
		SQL          *settings.SQLSetting
		Logger       *settings.LoggerSetting
		Cache        *settings.CacheSetting
		JWT          *settings.JWTSetting
	}
)

var (
	Log       *appLogger.Logger
	AppConfig *Config
	App       *fiber.App
	SQLDB     *database.Connection
	Cache     *caching.CacheClient
)
