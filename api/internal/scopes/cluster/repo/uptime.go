package repo

import "dashboard/api/internal/scopes/cluster/entities"

const PG_RUNTIME_QUERY = "SELECT pg_postmaster_start_time() as cluster_started_at"

func (s *Storage) Uptime() (entities.PostgresUptime, error) {

	db, err := s.pgManager.SQLX()
	if err != nil {
		return entities.PostgresUptime{}, err
	}

	var uptime entities.PostgresUptime

	err = db.Get(&uptime, PG_RUNTIME_QUERY)
	if err != nil {
		return entities.PostgresUptime{}, err
	}

	return uptime, nil
}
