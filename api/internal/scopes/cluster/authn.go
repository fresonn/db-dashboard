package cluster

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/scopes/cluster/entities"
	"errors"
)

func (c *Cluster) Connect(ctx context.Context, authData entities.AuthData) error {

	c.logger.InfoContext(ctx, "try to establish postgres connection")

	err := c.validate.Struct(&authData)
	if err != nil {
		c.logger.ErrorContext(ctx, "validation error", "error", err)
		return err
	}

	err = c.pgManager.UpdateConnection(ctx, config.Connection(authData))
	if err != nil {
		c.logger.ErrorContext(ctx, "update connection error", "error", err)
		return err
	}

	if ok := c.pgManager.IsConnected(); !ok {
		c.logger.WarnContext(ctx, "connection not established after its update")
		return errors.New("connection not established after its update")
	}

	c.logger.InfoContext(ctx, "postgres connection established")

	return nil
}
