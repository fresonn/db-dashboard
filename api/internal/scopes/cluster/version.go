package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
	"strings"
)

func (c *Cluster) Version(ctx context.Context) (entities.PostgresVersion, error) {

	c.logger.DebugContext(ctx, "try to get postgres version")

	if pgVersion, ok := c.cache.PgVersion(ctx); ok {
		c.logger.DebugContext(ctx, "got postgres version", "version", pgVersion.Version)
		return pgVersion, nil
	}

	rawVersion, err := c.storage.Version()
	if err != nil {
		c.logger.ErrorContext(ctx, "get postgres version", "error", err)
		return entities.PostgresVersion{}, err
	}

	parts := strings.Split(rawVersion, ",")

	if len(parts) == 1 {
		return entities.PostgresVersion{
			Version: strings.TrimSpace(parts[0]),
		}, nil
	}

	if len(parts) == 2 {
		return entities.PostgresVersion{
			Version:  strings.TrimSpace(parts[0]),
			Compiler: strings.TrimSpace(parts[1]),
		}, nil
	}

	version := strings.TrimSpace(parts[0])
	bitDepth := strings.TrimSpace(parts[len(parts)-1])
	compiler := strings.TrimSpace(strings.Join(parts[1:len(parts)-1], ","))

	pgVersion := entities.PostgresVersion{
		Version:  version,
		Compiler: compiler,
		BitDepth: bitDepth,
	}

	c.cache.SetPgVersion(ctx, pgVersion)

	c.logger.DebugContext(ctx, "got postgres version", "version", pgVersion.Version)

	return pgVersion, nil
}
