package database

import "crypto-dashboard-backend/configs/development"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxConns int
	Timeout  int
}

func DefaultConfig() *Config {
	return &Config{
		Host:     development.GetEnvStr("DB_HOST", "localhost"),
		Port:     development.GetEnvInt("DB_PORT", 5432),
		User:     development.GetEnvStr("DB_USER", "postgres"),
		Password: development.GetEnvStr("DB_PASSWORD", "postgres"),
		DBName:   development.GetEnvStr("DB_NAME", "crypto_dashboard"),
		SSLMode:  development.GetEnvStr("DB_SSL_MODE", "disable"),
		MaxConns: development.GetEnvInt("DB_MAX_CONNS", 10),
		Timeout:  development.GetEnvInt("DB_TIMEOUT", 5),
	}
}
