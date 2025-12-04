package cluster

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/scopes/cluster/entities"
	"errors"
)

func (c *Cluster) Connect(ctx context.Context, authData entities.AuthData) error {

	c.logger.Info("try to establish postgres connection")

	authData.PrettyLog()

	if ok := c.pgManager.IsConnected(); ok {
		c.logger.WarnContext(ctx, "connection already established")
		return errors.New("connection already established")
	}

	err := c.validate.Struct(&authData)
	if err != nil {
		c.logger.ErrorContext(ctx, "connection validation", "error", err)
		return err
	}

	conn := config.Connection{
		Host:     authData.Host,
		Port:     authData.Port,
		User:     authData.User,
		Password: authData.Password,
		SSLMode:  authData.SSLMode,
	}

	if authData.Database != nil {
		conn.Database = *authData.Database
	} else {
		conn.Database = "postgres"
	}

	err = c.pgManager.UpdateConnection(ctx, conn)
	if err != nil {
		c.logger.Error("set connection", "error", err)
		return err
	}

	c.logger.Info("postgres connection established")

	return nil
}
