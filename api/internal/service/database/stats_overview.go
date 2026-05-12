package database

import (
	"context"
	"dashboard/api/internal/model/database"
	"errors"
)

func (s *Service) StatsOverview(ctx context.Context, databaseID int) (database.DatabaseStatsOverview, error) {

	return database.DatabaseStatsOverview{}, errors.New("service not implemented in")
}
