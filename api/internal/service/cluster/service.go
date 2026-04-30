package cluster

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/model/cluster"
	"dashboard/api/internal/postgres"
	"dashboard/api/pkg/logger"
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Connect(ctx context.Context, authData cluster.AuthData) (cluster.Status, error)
	PostgresStatus(ctx context.Context) cluster.Status
	Uptime(ctx context.Context) (cluster.PostgresUptime, error)
	Version(ctx context.Context) (cluster.PostgresVersion, error)
	PostmasterSettings(ctx context.Context) (cluster.PostmasterSettings, error)
	Disconnect(ctx context.Context) error
}

type service struct {
	config    config.AppConfig
	logger    *slog.Logger
	pgManager *postgres.Manager
	validate  *validator.Validate
	storage   Storage
	cache     Cache
}

type Options struct {
	Config          config.AppConfig
	Logger          *slog.Logger
	PostgresManager *postgres.Manager
	Storage         Storage
	Cache           Cache
}

func New(options Options) *service {

	return &service{
		config:    options.Config,
		logger:    logger.WithScopeLogger(options.Logger, "cluster"),
		validate:  validator.New(validator.WithRequiredStructEnabled()),
		storage:   options.Storage,
		pgManager: options.PostgresManager,
		cache:     options.Cache,
	}
}
