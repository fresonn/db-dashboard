package cache

import (
	"context"
	"dashboard/api/internal/config"
	"dashboard/api/internal/scopes/cluster/entities"
	"dashboard/api/pkg/inmemory"
	"log/slog"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type Cache struct {
	logger *slog.Logger
	cache  *inmemory.Store[any]
}

func New(config *config.AppConfig, logger *slog.Logger) *Cache {

	cache := inmemory.New[any](5 * time.Minute)

	return &Cache{
		logger: logger,
		cache:  cache,
	}
}

func (c *Cache) DeleteAll(ctx context.Context) {
	c.logger.DebugContext(ctx, "clear cluster cache")
	c.cache.DeleteAll()
}

func (c *Cache) SetPgVersion(ctx context.Context, version entities.PostgresVersion) {
	c.logger.DebugContext(ctx, "cache set", "key", pgVersionKey, "value", version)
	c.cache.Set(pgVersionKey, version, ttlcache.NoTTL)
}

func (c *Cache) PgVersion(ctx context.Context) (entities.PostgresVersion, bool) {

	item := c.cache.Get(pgVersionKey)
	if item == nil {
		c.logger.DebugContext(ctx, "cache miss", "key", pgVersionKey)
		return entities.PostgresVersion{}, false
	}

	c.logger.DebugContext(ctx, "cache hit", "key", item.Key(), "expires_at", item.ExpiresAt())

	val, ok := item.Value().(entities.PostgresVersion)
	if !ok {
		c.logger.ErrorContext(ctx, "cache type cast failed", "key", item.Key())
		return entities.PostgresVersion{}, false
	}

	return val, true
}

func (c *Cache) SetClusterUptime(ctx context.Context, version entities.PostgresUptime) {
	c.logger.DebugContext(ctx, "cache set", "key", clusterUptimeKey, "value", version)
	c.cache.Set(clusterUptimeKey, version, ttlcache.NoTTL)
}

func (c *Cache) ClusterUptime(ctx context.Context) (entities.PostgresUptime, bool) {

	item := c.cache.Get(clusterUptimeKey)
	if item == nil {
		c.logger.DebugContext(ctx, "cache miss", "key", clusterUptimeKey)
		return entities.PostgresUptime{}, false
	}

	c.logger.DebugContext(ctx, "cache hit", "key", item.Key(), "expires_at", item.ExpiresAt())

	val, ok := item.Value().(entities.PostgresUptime)
	if !ok {
		c.logger.ErrorContext(ctx, "cache type cast failed", "key", item.Key())
		return entities.PostgresUptime{}, false
	}

	return val, true
}
