package database

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Connection struct {
	db  *pgxpool.Pool
	cfg *Config
}

func NewConnection(cfg *Config) (*Connection, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	return connect(dsn, cfg)
}
