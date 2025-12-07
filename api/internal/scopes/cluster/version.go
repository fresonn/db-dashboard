package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
	"strings"
)

func (c *Cluster) Version(ctx context.Context) (entities.PostgresVersion, error) {

	rawVersion, err := c.storage.Version()
	if err != nil {
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

	return entities.PostgresVersion{
		Version:  version,
		Compiler: compiler,
		BitDepth: bitDepth,
	}, nil
}
