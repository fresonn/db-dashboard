package usecase

import (
	"context"
	"dashboard/api/internal/model/cluster"
	"dashboard/api/internal/model/database"
)

type ClusterService interface {
	PostgresStatus(ctx context.Context) cluster.Status
}

type DatabaseService interface {
	Database(ctx context.Context, id int) (database.Database, error)
}
