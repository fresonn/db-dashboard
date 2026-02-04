package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
)

// More about pg_settings context values
// https://www.postgresql.org/docs/current/view-pg-settings.html

func (c *Cluster) PostmasterSettings(ctx context.Context) (entities.PostmasterSettings, error) {

	settings, err := c.storage.PostmasterSettings(ctx)
	if err != nil {
		c.logger.ErrorContext(ctx, "get cluster settings", "error", err)
		return entities.PostmasterSettings{}, err
	}

	return entities.PostmasterSettings{
		Settings: settings,
	}, nil
}
