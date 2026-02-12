package storage

import (
	"context"
	"fmt"
)

const DATABASES_DETAILS_QUERY = `
WITH databases_cte AS (
  SELECT
    datname AS name,
    datdba AS owner_id,
    encoding AS encoding_id,
    datcollate AS collate,
    datctype AS ctype,
    datistemplate AS is_template,
    datallowconn AS allow_connections,
    datconnlimit AS connection_limit,
    pg_database_size(datname) AS size_bytes
  FROM pg_database
),
connections AS (
  SELECT
    datname AS name,
    count(*) AS active_connections
  FROM pg_stat_activity
  WHERE state = 'active'
  GROUP BY datname
)
SELECT
  d.name,
  pg_get_userbyid(d.owner_id) AS owner,
  pg_encoding_to_char(d.encoding_id) AS encoding,
  d.collate,
  d.ctype,
  d.is_template,
  d.allow_connections,
  d.connection_limit,
  d.size_bytes,
  pg_size_pretty(d.size_bytes) AS size_pretty,
  COALESCE(c.active_connections, 0) AS active_connections
FROM databases_cte d
LEFT JOIN connections c ON c.name = d.name
ORDER BY d.size_bytes DESC;
`

func (s *Storage) DatabasesDetails(ctx context.Context) error {

	db, err := s.pgManager.SQLX()
	if err != nil {
		return err
	}

	var dtos []DatabaseDetails

	err = db.Select(&dtos, DATABASES_DETAILS_QUERY)
	if err != nil {
		return err
	}

	for _, v := range dtos {
		fmt.Printf("%+v\n", v)
	}

	return nil

}
