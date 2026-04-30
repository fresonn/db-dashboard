package roles

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/postgres"
	"dashboard/api/pkg/logger"
	"log/slog"

	"dashboard/api/internal/model/role"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Roles(ctx context.Context) ([]role.RoleView, error)
}

type service struct {
	config    config.AppConfig
	logger    *slog.Logger
	pgManager *postgres.Manager
	validate  *validator.Validate
	storage   Storage
}

type Options struct {
	Config          config.AppConfig
	Logger          *slog.Logger
	PostgresManager *postgres.Manager
	Storage         Storage
}

func New(options Options) *service {

	return &service{
		config:    options.Config,
		logger:    logger.WithScopeLogger(options.Logger, "cluster"),
		validate:  validator.New(validator.WithRequiredStructEnabled()),
		storage:   options.Storage,
		pgManager: options.PostgresManager,
	}
}
