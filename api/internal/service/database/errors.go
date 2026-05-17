package database

import "errors"

var (
	ErrNotFound = errors.New("database not found")
)

var (
	ErrOverviewStatStateNotFound = errors.New("stats overview not found")
)
