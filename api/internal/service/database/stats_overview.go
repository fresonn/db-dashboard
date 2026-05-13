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

	currentStats, err := s.pg.CurrentDBOverviewStats(ctx)
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

	fmt.Printf("current ---- %+v\n", currentStats)

	existingStats, err := s.storage.StatsOverview(databaseID, currentStats)
	if err != nil {

		if errors.Is(err, ErrStatsOverviewNotFound) {
			s.logger.Warn("stats overview not found", "id", databaseID)
		} else {

			s.logger.ErrorContext(ctx, "get existing stats failed", "id", databaseID, "error", err)
			return database.DatabaseStatsOverview{}, err
		}
	}

	err = s.storage.SaveStatsOverview(databaseID, currentStats)
	if err != nil {
		s.logger.ErrorContext(ctx, "save stats failed", "id", databaseID, "error", err)
		return database.DatabaseStatsOverview{}, err
	}

	fmt.Printf("existing ---- %+v\n", existingStats)

	return database.DatabaseStatsOverview{
		ID:   helper.IntToString(currentStats.ID),
		Name: currentStats.Name,
		Size: database.DatabaseSizeStaticStat{
			SizeBytes:  currentStats.Size,
			SizePretty: helper.PrettyByteSize(currentStats.Size),
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
		Tables: database.DatabaseTablesStaticStat{
			Total: currentStats.Tables,
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
		Indexes: database.DatabaseIndexesStaticStat{
			Total: currentStats.Indexes,
			Trend: database.StaticStatTrend{
				Value:     "?",
				Direction: database.StatDirectionUp,
			},
		},
	}, nil
}
