package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
	"dashboard/api/internal/utils"
	"fmt"
	"log/slog"
)

const (
	ParamConfigFile     = "config_file"
	ParamDataDirectory  = "data_directory"
	ParamSharedBuffers  = "shared_buffers"
	ParamWalBuffers     = "wal_buffers"
	ParamMaxConnections = "max_connections"
	ParamHbaFile        = "hba_file"
	ParamWalLevel       = "wal_level"
)

var postmasterParams = []string{
	ParamConfigFile,
	ParamDataDirectory,
	ParamSharedBuffers,
	ParamWalBuffers,
	ParamMaxConnections,
	ParamHbaFile,
	ParamWalLevel,
}

// More about pg_settings and its context values
// https://www.postgresql.org/docs/current/view-pg-settings.html
func (c *Cluster) PostmasterSettings(ctx context.Context) (entities.PostmasterSettings, error) {

	settings, err := c.storage.PostmasterSettings(ctx, postmasterParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "get cluster settings", "error", err)
		return entities.PostmasterSettings{}, err
	}

	settingsMap := make(map[string]entities.Setting, len(settings))

	for _, s := range settings {
		settingsMap[s.Name] = s
	}

	sharedBuffers, found := parseSizeSetting(ctx, c.logger, settingsMap, ParamSharedBuffers)
	if !found {
		c.logger.WarnContext(ctx, "shared_buffers not found in pg_settings or corrupt")
	}

	walBuffers, found := parseSizeSetting(ctx, c.logger, settingsMap, ParamWalBuffers)
	if !found {
		c.logger.WarnContext(ctx, "wal_buffers not found in pg_settings or corrupt")
	}

	return entities.PostmasterSettings{
		ConfigFile:     settingsMap[ParamConfigFile],
		DataDirectory:  settingsMap[ParamDataDirectory],
		SharedBuffers:  sharedBuffers,
		WalBuffers:     walBuffers,
		MaxConnections: settingsMap[ParamMaxConnections],
		HbaFile:        settingsMap[ParamHbaFile],
		WalLevel:       settingsMap[ParamWalLevel],
	}, nil
}

func parseSizeSetting(ctx context.Context, logger *slog.Logger, store map[string]entities.Setting, key string) (entities.Setting, bool) {
	setting, ok := store[key]
	if !ok {
		return entities.Setting{}, false
	}

	if setting.Unit == "8kB" || setting.Unit == "8KB" {
		n, err := utils.ParseInt64(setting.Value)
		if err != nil {
			logger.WarnContext(ctx, fmt.Sprintf("failed to parse %s", key), "value", setting.Value)
			return entities.Setting{}, false
		}

		setting.Value = utils.PrettyByteSize((8 * utils.KB) * n)
	}

	return setting, true
}
