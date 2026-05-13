package database

import (
	"context"
	"dashboard/api/internal/helper"
	"dashboard/api/internal/model/database"
	"errors"
	"fmt"
)

// Returns simple stats for current connected database, not only by database ID, connection required
func (s *Service) StatsOverview(ctx context.Context, databaseID int) (database.DatabaseStatsOverview, error) {

	// db, err := s.Database(ctx, databaseID)
	// if err != nil {
	// 	return database.DatabaseStatsOverview{}, err
	// }

	// dbName, err := s.storage.DatabaseName(databaseID)
	// if err != nil {
	// 	s.logger.Error("get db name", "error", err)

	// 	// temp, for RND
	// 	if err.Error() == "Key not found" {
	// 		s.storage.SetDatabaseName(databaseID, db.Name)
	// 	}
	// }

	// fmt.Println("DB name:", dbName)

	pgStats, err := s.pg.CurrentDBOverviewStats(ctx)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			s.logger.Warn("database stats not found", "id", databaseID)
			return database.DatabaseStatsOverview{}, err
		}

		s.logger.ErrorContext(ctx, "fetch stats failed", "id", databaseID, "error", err)
		return database.DatabaseStatsOverview{}, err
	}

	// if pgStats.ID != int64(databaseID) {
	// 	return database.DatabaseStatsOverview{}, fmt.Errorf("to get stats connect to database=%d first, current=%d", databaseID, pgStats.ID)
	// }

	fmt.Printf("%+v\n", pgStats)

	return database.DatabaseStatsOverview{
		ID:   helper.IntToString(pgStats.ID),
		Name: pgStats.Name,
		Size: database.DatabaseSizeStaticStat{
			SizeBytes:  pgStats.Size,
			SizePretty: helper.PrettyByteSize(pgStats.Size),
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
		Tables: database.DatabaseTablesStaticStat{
			Total: pgStats.Tables,
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
		Indexes: database.DatabaseIndexesStaticStat{
			Total: pgStats.Indexes,
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
	}, nil
}
