package database

import (
	"context"
	"dashboard/api/internal/model/database"
)

func (s *Service) Database(ctx context.Context, id int) (database.Database, error) {

	if db, ok := s.cache.Database(ctx, id); ok {
		s.logger.DebugContext(ctx, "database cache hit", "id", id, "value", db)
		return db, nil
	}

	s.logger.DebugContext(ctx, "database cache miss", "id", id)

	db, err := s.pg.Database(ctx, id)
	if err != nil {
		s.logger.ErrorContext(ctx, "database fetch failed", "id", id, "error", err)
		return database.Database{}, err
	}

	s.cache.SetDatabase(ctx, id, db)

	s.logger.DebugContext(ctx, "database fetched", "id", id)

	return db, nil
}
