package cluster

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/model/cluster"

	"errors"
)

func (s *Service) Connect(ctx context.Context, newConn cluster.NewConnection) (cluster.Status, error) {

	if ok := s.pgManager.IsConnected(); ok {
		s.logger.WarnContext(ctx, "connection already established")
		return cluster.Status{}, errors.New("connection already established")
	}

	err := s.validate.Struct(&newConn)
	if err != nil {
		s.logger.ErrorContext(ctx, "connection validation failed", "error", err)
		return cluster.Status{}, err
	}

	conn := config.Connection{
		Host:     newConn.Host,
		Port:     newConn.Port,
		User:     newConn.User,
		Password: newConn.Password,
		SSLMode:  newConn.SSLMode,
	}

	if newConn.Database != nil {
		conn.Database = *newConn.Database
	} else {
		conn.Database = "postgres"
	}

	err = s.pgManager.UpdateConnection(ctx, conn)
	if err != nil {
		s.logger.Error("failed to update postgres connection", "error", err)
		return cluster.Status{}, err
	}

	s.logger.Info("postgres connection established",
		"host", conn.Host,
		"port", conn.Port,
		"user", conn.User,
		"password", "***",
		"database", conn.Database,
		"sslmode", conn.SSLMode,
	)

	return cluster.Status{
		ConnectionStatus: s.pgManager.Status(),
		CurrentUser:      &conn.User,
		CurrentDatabase:  &conn.Database,
	}, nil
}
