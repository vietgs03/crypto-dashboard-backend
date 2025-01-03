package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func connect(dsn string, cfg *Config) (*Connection, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	db.SetMaxOpenConns(cfg.MaxConns)
	db.SetMaxIdleConns(cfg.MaxConns)
	db.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	return &Connection{
		db:  db,
		cfg: cfg,
	}, nil
}

func (c *Connection) DB() *sql.DB {
	return c.db
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func (c *Connection) HealthCheck(ctx context.Context) error {
	return c.db.PingContext(ctx)
}
