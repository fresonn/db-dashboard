package storage

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
)

const PG_SETTINGS_QUERY = `
SELECT
    name,
    setting,
    unit,
    short_desc
FROM pg_settings
WHERE context = $1 AND name IN (
    'config_file',
    'data_directory',
    'shared_buffers',
    'wal_buffers',
    'max_connections',
    'hba_file',
    'wal_level'
)`

func (s *Storage) PostmasterSettings(ctx context.Context) ([]entities.Setting, error) {

	db, err := s.pgManager.SQLX()
	if err != nil {
		return nil, err
	}

	var dtos []Setting

	err = db.Select(&dtos, PG_SETTINGS_QUERY, "postmaster")
	if err != nil {
		return nil, err
	}

	settings := make([]entities.Setting, 0, len(dtos))

	for _, dto := range dtos {
		settings = append(settings, toSettingEntity(dto))
	}

	return settings, nil
}
