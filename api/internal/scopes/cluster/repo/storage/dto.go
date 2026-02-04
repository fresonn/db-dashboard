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

func toSettingEntity(dto Setting) entities.Setting {
	return entities.Setting{
		Name:        dto.Name,
		Setting:     dto.Setting,
		Unit:        dto.Unit.String(),
		Description: dto.ShortDesc,
	}
}
