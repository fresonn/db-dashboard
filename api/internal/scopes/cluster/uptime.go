package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
)

func (c *Cluster) Uptime(ctx context.Context) (entities.PostgresUptime, error) {

	c.logger.DebugContext(ctx, "try to get cluster uptime")

	if uptime, ok := c.cache.ClusterUptime(ctx); ok {
		c.logger.DebugContext(ctx, "got cluster uptime", "started_at", uptime.StartedAt)
		return uptime, nil
	}

	uptime, err := c.storage.Uptime()
	if err != nil {
		return entities.PostgresUptime{}, err
	}

	c.cache.SetClusterUptime(ctx, uptime)

	c.logger.DebugContext(ctx, "got cluster uptime", "started_at", uptime.StartedAt)

	return uptime, nil
}
