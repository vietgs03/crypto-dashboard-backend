package database

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type Repository[T any] interface {
	Create(ctx context.Context, entity T) error
	FindByID(ctx context.Context, id string) (T, error)
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, id string) error
}

// Base PostgreSQL implementation
type BasePostgresRepository[T any] struct {
	db  *sql.DB
	log *zap.Logger
}
