package redis

import "crypto-dashboard-backend/configs/development"

type ConfigRedis struct {
	Host     string
	Port     string
	Password string
	Database int
	PoolSize int
}

func DefaultConfigRedis() *ConfigRedis {
	return &ConfigRedis{
		Password: development.GetEnvStr("DBR_PASSWORD", ""),
		Database: development.GetEnvInt("DBR_NAME", 0),
		Host:     development.GetEnvStr("DBR_HOST", ""),
		Port:     development.GetEnvStr("DBR_PORT", ""),
		PoolSize: development.GetEnvInt("DBR_POOL_SIZE", 5),
	}
}
