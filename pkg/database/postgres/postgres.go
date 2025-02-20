package database

import (
	"context"
	"fmt"
	"time"

	"crypto-dashboard/pkg/response"
	"crypto-dashboard/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connect(dsn string, cfg *settings.SQLSetting) (*Connection, *response.AppError) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, response.DatabaseError(fmt.Errorf("%w: %v", ErrConfigFailed, err))
	}

	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MaxConnLifetime = time.Duration(cfg.MaxConnLifetime)
	poolConfig.MaxConnIdleTime = time.Duration(cfg.MaxConnIdleTime)

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, response.DatabaseError(fmt.Errorf("%w: %v", ErrConnectionFailed, err))
	}

	carpool, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MaxConnLifetime)*time.Second)
	defer cancel()

	if err := db.Ping(carpool); err != nil {
		return nil, response.DatabaseError(fmt.Errorf("connection test failed: %w", err))
	}

	return &Connection{
		db:  db,
		cfg: cfg,
	}, nil
}

func (c *Connection) DB() *pgxpool.Pool {
	return c.db
}

func (c *Connection) Close() error {
	c.db.Close()
	return nil
}

func (c *Connection) HealthCheck(ctx context.Context) *response.AppError {
	if err := c.db.Ping(ctx); err != nil {
		return response.DatabaseError(fmt.Errorf("%w: %v", ErrHealthCheckFailed, err))
	}
	return nil
}
