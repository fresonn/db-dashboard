package config

import (
	"errors"
	"log"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
)

const (
	LiveRuntime = "live"
	DevRuntime  = "dev"
	TestRuntime = "test"
)

type AppConfig struct {
	Env              EnvironmentVars
	PersistentConfig *PersistentConfigV1
}

type EnvironmentVars struct {
	IsDev      bool
	Port       int    `env:"SERVER_PORT"`
	Runtime    string `env:"RUNTIME"`
	PgHost     string `env:"PG_HOST"`
	PgPort     int    `env:"PG_PORT"`
	PgUser     string `env:"PG_USER"`
	PgPassword string `env:"PG_PASSWORD"`
	PgDatabase string `env:"PG_DATABASE"`
	PgSSLMode  string `env:"PG_SSLMODE"`
}

func New() AppConfig {

	var envCfg EnvironmentVars

	err := env.Parse(&envCfg)
	if err != nil {
		log.Fatalln("environment variables parse error", err)
	}

	envCfg.IsDev = envCfg.Runtime == DevRuntime

	persistentCfg, err := configFromEnvs(envCfg)
	if err != nil {
		slog.Error("parse environment variables ", "error", err)
		os.Exit(1)
	}

	if persistentCfg != nil {
		return AppConfig{
			Env:              envCfg,
			PersistentConfig: persistentCfg,
		}
	}

	persistentCfg, err = LoadPersistentConfig()
	if err != nil {

		slog.Warn("volume data for connection not found or damaged")
		_ = RemovePersistentConfig()

		return AppConfig{
			Env:              envCfg,
			PersistentConfig: nil,
		}
	}

	err = ValidatePersistentConfig(persistentCfg)
	if err != nil {
		slog.Warn("volume /data/config.json is damaged")
		_ = RemovePersistentConfig()
		return AppConfig{
			Env:              envCfg,
			PersistentConfig: nil,
		}
	}

	return AppConfig{
		Env:              envCfg,
		PersistentConfig: persistentCfg,
	}
}

func configFromEnvs(envs EnvironmentVars) (*PersistentConfigV1, error) {

	var foundCount int

	acceptedCount := 5

	persistentCfg := &PersistentConfigV1{
		Version: CurrentVersion,
	}

	if v := envs.PgHost; v != "" {
		persistentCfg.Connection.Host = v
		foundCount++
	}

	if v := envs.PgPort; v > 0 {
		persistentCfg.Connection.Port = v
		foundCount++
	}

	if v := envs.PgUser; v != "" {
		persistentCfg.Connection.User = v
		foundCount++
	}

	if v := envs.PgPassword; v != "" {
		persistentCfg.Connection.Password = v
		foundCount++
	}

	if v := envs.PgDatabase; v != "" {
		persistentCfg.Connection.Database = v
		foundCount++
	}

	// if v := envs.PgSSLMode; v != "" {
	// 	persistentCfg.Connection.SSLMode = v
	// 	foundCount++
	// }

	if foundCount > 0 && foundCount != acceptedCount {
		return nil, errors.New("specify all necessary PG_* envs or remove all for manual Postgres authentication")
	}

	if foundCount == acceptedCount {
		return persistentCfg, nil
	}

	return nil, nil
}
