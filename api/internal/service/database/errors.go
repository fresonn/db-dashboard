package database

import "errors"

var (
	ErrDatabaseNotFound          = errors.New("database not found")
	ErrFailedGetDatabase         = errors.New("failed to get database")
	ErrOverviewStatStateNotFound = errors.New("stats overview not found")
)
