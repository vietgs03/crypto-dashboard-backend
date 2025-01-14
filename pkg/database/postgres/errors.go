package database

import "errors"

var (
	ErrConnectionFailed  = errors.New("failed to connect to database")
	ErrConfigFailed      = errors.New("failed to connect to config")
	ErrHealthCheckFailed = errors.New("health check failed")
	ErrQueryFailed       = errors.New("failed to execute query")
	ErrNoRows            = errors.New("no rows found")
)
