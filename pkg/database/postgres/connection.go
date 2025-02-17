package database

import (
	"fmt"

	"crypto-dashboard/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Connection struct {
	db  *pgxpool.Pool
	cfg *settings.SQLSetting
}

func NewConnection(cfg *settings.SQLSetting) (*Connection, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode,
	)
	return connect(dsn, cfg)
}
