package cluster

import "context"

func (c *Cluster) Version(ctx context.Context) (string, error) {
	return c.storage.Version()
}
