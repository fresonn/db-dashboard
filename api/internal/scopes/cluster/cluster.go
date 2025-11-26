package cluster

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/postgres"
	"dashboard/api/pkg/logger"
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type Cluster struct {
	config    config.AppConfig
	logger    *slog.Logger
	pgManager *postgres.Manager
	validate  *validator.Validate
	storage   ClusterStorage
}

type ClusterStorage interface {
	Version() (string, error)
}

func New(config config.AppConfig, lgr *slog.Logger, storage ClusterStorage, pgManager *postgres.Manager) *Cluster {

	return &Cluster{
		config:    config,
		logger:    logger.WithScopeLogger(lgr, "cluster"),
		validate:  validator.New(validator.WithRequiredStructEnabled()),
		storage:   storage,
		pgManager: pgManager,
	}
}

func (c *Cluster) PostgresStatus(ctx context.Context) postgres.ConnectionStatus {

	c.logger.DebugContext(ctx, "get postres connection status")

	switch {
	case c.pgManager.IsConnected():
		return postgres.StatusConnected
	case c.pgManager.Connecting():
		return postgres.StatusConnecting
	default:
		return postgres.StatusDisconnected
	}
}
