package database

import "errors"

var (
	ErrDatabaseNotFound = errors.New("database not found")
)

var (
	ErrOverviewStatStateNotFound = errors.New("stats overview not found")
)
