package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
)

func (c *Cluster) DatabasesDetailed(ctx context.Context, filter entities.DatabasesFilter) ([]entities.DatabaseDetails, error) {

	databases, err := c.storage.DatabasesDetails(ctx, filter)
	if err != nil {
		c.logger.ErrorContext(ctx, "available databases", "error", err)
		return nil, err
	}

	return databases, nil
}
