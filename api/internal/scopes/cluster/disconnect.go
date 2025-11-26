package cluster

import (
	"context"
	"errors"
)

func (c *Cluster) Disconnect(ctx context.Context) error {

	c.logger.Info("try to disconnect postgres")

	isConnected := c.pgManager.IsConnected()

	if !isConnected {
		c.logger.Warn("already disconnected from postgres")
		return errors.New("already disconnected from postgres")
	}

	err := c.pgManager.Disconnect()
	if err != nil {
		c.logger.Error("disconnect from postgres", "error", err)
		return err
	}

	c.logger.Info("postgres disconnected")

	return nil
}
