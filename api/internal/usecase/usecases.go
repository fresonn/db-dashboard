package usecase

import (
	"dashboard/api/internal/infra/postgres"
)

type useCases struct {
	clusterService  ClusterService
	databaseService DatabaseService
	pgManager       *postgres.Manager
}

type Options struct {
	ClusterService  ClusterService
	DatabaseService DatabaseService
	PgManager       *postgres.Manager
}

func New(options Options) *useCases {
	return &useCases{
		clusterService:  options.ClusterService,
		databaseService: options.DatabaseService,
		pgManager:       options.PgManager,
	}
}
