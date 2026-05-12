package database

import (
	"context"
	"dashboard/api/internal/model/database"
	"errors"
	"fmt"
)

func (s *Service) StatsOverview(ctx context.Context, databaseID int) (database.DatabaseStatsOverview, error) {

	db, err := s.Database(ctx, databaseID)
	if err != nil {
		return database.DatabaseStatsOverview{}, err
	}

	dbName, err := s.storage.DatabaseName(databaseID)
	if err != nil {
		s.logger.Error("get db name", "error", err)

		// temp, for RND
		if err.Error() == "Key not found" {
			s.storage.SetDatabaseName(databaseID, db.Name)
		}
	}

	fmt.Println("DB name:", dbName)

	return database.DatabaseStatsOverview{}, errors.New("service not implemented in")
}
