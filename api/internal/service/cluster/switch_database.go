package cluster

import (
	"context"
	"dashboard/api/internal/model/cluster"
	"errors"
	"fmt"
)

// If the connection is successful, switch to the new database and update the status
func (s *Service) SwitchDatabase(ctx context.Context, id int) (cluster.Status, error) {
	status := s.pgManager.Connection()

	fmt.Printf("%+v\n", status)

	return cluster.Status{}, errors.New("not implemented")
}
