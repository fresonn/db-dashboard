package database

import "errors"

var (
	ErrNotFound = errors.New("database not found")
)

var (
	ErrStatsOverviewNotFound = errors.New("stats overview not found")
)
