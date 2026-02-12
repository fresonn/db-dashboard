package storage

import (
	"dashboard/api/internal/postgres"
	"dashboard/api/internal/scopes/cluster/entities"
)

type Setting struct {
	Name      string        `db:"name"`
	Setting   string        `db:"setting"`
	Unit      postgres.Text `db:"unit"`
	ShortDesc string        `db:"short_desc"`
}

type DatabaseDetails struct {
	Name string `db:"name"`
	// If the OID is invalid or the role has been deleted, the function "pg_get_userbyid(owner_id)", will return NULL
	Owner postgres.Text `db:"owner"`
	// Encoding NULL, if the base is damaged, theoretically
	Encoding          postgres.Text `db:"encoding"`
	Collate           string        `db:"collate"`
	Ctype             string        `db:"ctype"`
	IsTemplate        bool          `db:"is_template"`
	AllowConnections  bool          `db:"allow_connections"`
	ConnectionLimit   int           `db:"connection_limit"`
	SizeBytes         int64         `db:"size_bytes"`
	SizePretty        string        `db:"size_pretty"`
	ActiveConnections int           `db:"active_connections"`
}

func toSettingEntity(dto Setting) entities.Setting {
	return entities.Setting{
		Name:        dto.Name,
		Value:       dto.Setting,
		Unit:        dto.Unit.String(),
		Description: dto.ShortDesc,
	}
}
