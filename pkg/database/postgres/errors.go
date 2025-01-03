package database

import "errors"

var (
	ErrConnectionFailed = errors.New("failed to connect to database")
	ErrQueryFailed      = errors.New("failed to execute query")
	ErrNoRows           = errors.New("no rows found")
)
