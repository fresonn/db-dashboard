package postgres

import (
	"context"
	"dashboard/api/internal/model/database"
	service "dashboard/api/internal/service/database"
	"database/sql"
	"errors"
)

/*
Postgres doesn't have cross-database access to pg_class,
so we can only retrieve statistics by connecting to the database in advance

Get metadata from system catalog by pg_class.relkind (relkind char)
https://www.postgresql.org/docs/current/catalog-pg-class.html

	r = ordinary table
	i = index
	p = partitioned table

Here is c.relkind = 'i' without partitioned indexes
*/

const CURRENT_DB_OVERVIEW_STATS_QUERY = `
SELECT
    current_database() AS name,
    (SELECT oid FROM pg_database WHERE datname = current_database()) AS oid,
    pg_database_size(current_database())::bigint AS size_current,
    (
        SELECT COUNT(*)
        FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        WHERE c.relkind IN ('r', 'p')
        AND n.nspname NOT IN ('pg_catalog', 'information_schema')
    ) AS tables_current,
    (
        SELECT COUNT(*)
        FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        WHERE c.relkind = 'i'
        AND n.nspname NOT IN ('pg_catalog', 'information_schema')
    ) AS indexes_current;
`

func (s *Storage) CurrentDBOverviewStats(ctx context.Context) (database.PostgresDbOverviewStats, error) {

	db, err := s.pgManager.SQLX()
	if err != nil {
		return database.PostgresDbOverviewStats{}, err
	}

	var dto overviewStats

	err = db.Get(&dto, CURRENT_DB_OVERVIEW_STATS_QUERY)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.PostgresDbOverviewStats{}, service.ErrDatabaseNotFound
		}

		return database.PostgresDbOverviewStats{}, err
	}

	return toOverviewStats(dto), nil
}
