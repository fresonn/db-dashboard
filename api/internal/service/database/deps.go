package database

import (
	"context"
	"dashboard/api/internal/model/database"
)

type PostgresRepo interface {
	Database(ctx context.Context, id int) (database.Database, error)
	DatabasesDetails(ctx context.Context, filters database.DatabasesFilter) ([]database.DatabaseDetails, error)
	CurrentDBOverviewStats(ctx context.Context) (database.PostgresDbOverviewStats, error)
}

type Cache interface {
	Database(ctx context.Context, id int) (database.Database, bool)
	SetDatabase(ctx context.Context, id int, db database.Database)
}

type Storage interface {
	StatsOverview(databaseID int) (database.StoredOverviewStats, error)
	SaveStatsOverview(databaseID int, stats database.StoredOverviewStats) error
}
