package cluster

import (
	"context"
	"dashboard/api/internal/model/cluster"
)

func (s *Service) Uptime(ctx context.Context) (cluster.PostgresUptime, error) {

	if uptime, ok := s.cache.ClusterUptime(ctx); ok {
		s.logger.DebugContext(ctx, "cluster uptime cache hit", "started_at", uptime.StartedAt)
		return uptime, nil
	}

	s.logger.DebugContext(ctx, "cluster uptime cache miss")

	uptime, err := s.pg.Uptime()
	if err != nil {
		s.logger.ErrorContext(ctx, "cluster uptime fetch failed", "error", err)
		return cluster.PostgresUptime{}, err
	}

	s.cache.SetClusterUptime(ctx, uptime)

	s.logger.DebugContext(ctx, "cluster uptime fetched", "started_at", uptime.StartedAt)

	return uptime, nil
}
