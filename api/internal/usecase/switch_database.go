package usecase

import (
	"context"
	"dashboard/api/internal/model/cluster"
)

// If the connection is successful, switch to the new database and update connection
func (u *useCases) SwitchDatabase(ctx context.Context, id int) (cluster.Status, error) {
	conn := u.pgManager.Connection()

	database, err := u.databaseService.Database(ctx, id)
	if err != nil {
		return cluster.Status{}, err
	}

	conn.Database = database.Name

	err = u.pgManager.UpdateConnection(ctx, *conn)
	if err != nil {
		return cluster.Status{}, err
	}

	newStatus := u.clusterService.PostgresStatus(ctx)

	return newStatus, nil
}
