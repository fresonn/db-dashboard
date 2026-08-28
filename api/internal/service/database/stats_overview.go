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

		if errors.Is(err, ErrDatabaseNotFound) {
			s.logger.DebugContext(ctx, "database stats not found", "id", databaseID)
			return database.DatabaseStatsOverview{}, err
		}

		s.logger.ErrorContext(ctx, "fetch stats failed", "id", databaseID, "error", err)
		return database.DatabaseStatsOverview{}, err
	}

	if databaseID != int(currentStats.ID) {
		return database.DatabaseStatsOverview{}, ErrDatabaseNotFound
	}

	sizeTrend, err := s.processOverviewMetric(ctx, databaseID, database.MetricSize, currentStats.Size)
	if err != nil {
		return database.DatabaseStatsOverview{}, err
	}

	sizeTrend.Value = helper.PrettyByteSize(sizeTrend.Diff)

	tablesTrend, err := s.processOverviewMetric(ctx, databaseID, database.MetricTables, currentStats.Tables)
	if err != nil {
		return database.DatabaseStatsOverview{}, err
	}

	indexesTrend, err := s.processOverviewMetric(ctx, databaseID, database.MetricIndexes, currentStats.Indexes)
	if err != nil {
		return database.DatabaseStatsOverview{}, err
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

func (s *Service) processOverviewMetric(ctx context.Context, databaseID int, metric database.Metric, currentValue int64) (database.StaticStatTrend, error) {

	state, err := s.storage.OverviewStatState(databaseID, metric)

	hasState := err == nil

	if err != nil && !errors.Is(err, ErrOverviewStatStateNotFound) {
		s.logger.ErrorContext(ctx, "get metric state failed", "id", databaseID, "metric", metric, "error", err)
		return database.StaticStatTrend{}, err
	}

	hasChanges := !hasState || currentValue != state.Current

	if hasChanges {

		nextState := database.OverviewStatState{
			Current:   currentValue,
			Previous:  state.Current,
			ChangedAt: time.Now(),
		}

		if !hasState {
			nextState.Previous = currentValue
		}

		err = s.storage.SaveOverviewStatState(databaseID, metric, nextState)
		if err != nil {
			s.logger.ErrorContext(ctx, "save metric state failed", "id", databaseID, "metric", metric, "error", err)
			return database.StaticStatTrend{}, err
		}

		state = nextState
	}

	showTrend := hasState && time.Since(state.ChangedAt) < 15*time.Minute

	if !showTrend {
		return database.StaticStatTrend{}, nil
	}

	return calculateTrend(state.Current, state.Previous), nil
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
