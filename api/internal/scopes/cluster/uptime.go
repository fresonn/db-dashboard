package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
)

func (c *Cluster) Uptime(ctx context.Context) (entities.PostgresUptime, error) {

	uptime, err := c.storage.Uptime()
	if err != nil {
		return entities.PostgresUptime{}, err
	}

	c.logger.InfoContext(ctx, "get uptime", "started_at", uptime.StartedAt.String())

	return uptime, nil
}
