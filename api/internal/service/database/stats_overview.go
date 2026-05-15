package database

import (
	"context"
	"dashboard/api/internal/helper"
	"dashboard/api/internal/model/database"
	"errors"
	"time"
)

/*
	Ephemeral trend - shows the trend of the last N minutes:

	if now - LastChanged < TIME:
		show trend
	else:
		hide trend

Save without TTL in storage, because it breaks deterministic behavior
Possible add user-based settings: user → 5m → 30m → 1h → 6h → 24h
*/
func (s *Service) StatsOverview(ctx context.Context, databaseID int) (database.DatabaseStatsOverview, error) {

	currentStats, err := s.pg.CurrentDBOverviewStats(ctx)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			s.logger.DebugContext(ctx, "database stats not found", "id", databaseID)
			return database.DatabaseStatsOverview{}, err
		}

		s.logger.ErrorContext(ctx, "fetch stats failed", "id", databaseID, "error", err)
		return database.DatabaseStatsOverview{}, err
	}

	if databaseID != int(currentStats.ID) {
		return database.DatabaseStatsOverview{}, ErrNotFound
	}

	storedStats, err := s.storage.StatsOverview(databaseID)

	hasStored := err == nil

	if err != nil && !errors.Is(err, ErrStatsOverviewNotFound) {
		s.logger.ErrorContext(ctx, "get stored stats failed", "id", databaseID, "error", err)
		return database.DatabaseStatsOverview{}, err
	}

	hasChanges := !hasStored ||
		currentStats.Size != storedStats.Current.Size ||
		currentStats.Tables != storedStats.Current.Tables ||
		currentStats.Indexes != storedStats.Current.Indexes

	if hasChanges {

		nextStored := database.StoredOverviewStats{
			Current:   currentStats,
			Previous:  storedStats.Current,
			ChangedAt: time.Now(),
		}

		// first save
		if !hasStored {
			nextStored.Previous = currentStats
		}

		err = s.storage.SaveStatsOverview(databaseID, nextStored)
		if err != nil {
			s.logger.ErrorContext(ctx, "save stats failed", "id", databaseID, "error", err)
			return database.DatabaseStatsOverview{}, err
		}

		storedStats = nextStored
	}

	showTrend := hasStored && time.Since(storedStats.ChangedAt) < 5*time.Minute

	sizeTrend := database.StaticStatTrend{}
	tablesTrend := database.StaticStatTrend{}
	indexesTrend := database.StaticStatTrend{}

	if showTrend {

		sizeTrend = calculateTrend(
			storedStats.Current.Size,
			storedStats.Previous.Size,
		)

		sizeTrend.Value = helper.PrettyByteSize(sizeTrend.Diff)

		tablesTrend = calculateTrend(
			storedStats.Current.Tables,
			storedStats.Previous.Tables,
		)

		indexesTrend = calculateTrend(
			storedStats.Current.Indexes,
			storedStats.Previous.Indexes,
		)
	}

	overviewStats := database.DatabaseStatsOverview{
		ID:     helper.IntToString(currentStats.ID),
		DbName: currentStats.Name,

		Size: database.DatabaseSizeStaticStat{
			SizeBytes:  currentStats.Size,
			SizePretty: helper.PrettyByteSize(currentStats.Size),
			Trend:      sizeTrend,
		},

		Tables: database.DatabaseTablesStaticStat{
			Total: currentStats.Tables,
			Trend: tablesTrend,
		},

		Indexes: database.DatabaseIndexesStaticStat{
			Total: currentStats.Indexes,
			Trend: indexesTrend,
		},
	}

	return overviewStats, nil
}

func calculateTrend(current, previous int64) database.StaticStatTrend {

	if current == previous {
		return database.StaticStatTrend{}
	}

	diff := current - previous

	direction := database.StatDirectionUp

	if diff < 0 {
		direction = database.StatDirectionDown
		diff = -diff
	}

	return database.StaticStatTrend{
		Diff:      diff,
		Value:     helper.IntToString(diff),
		Direction: direction,
	}
}
