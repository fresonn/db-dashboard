package config

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	CurrentVersion = 1
	ConfigPath     = "./data/config.json"
)

var (
	ErrNoConfig = errors.New("no config available")
)

type Connection struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

type PersistentConfigV1 struct {
	Version    int        `json:"version"`
	Connection Connection `json:"connection"`
}

// LoadConfig loads config.json → applies ENV → returns the finished config
func LoadPersistentConfig() (*PersistentConfigV1, error) {

	cfg, err := loadFromFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	migratedCfg, changed := Migrate(cfg)
	if changed {
		slog.Warn("config.json outdated, try to migrate")
		_ = migratedCfg.Save() // it's not critical if it's not saved
	}

	return cfg, nil
}

func loadFromFile(path string) (*PersistentConfigV1, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg PersistentConfigV1
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Migrate(cfg *PersistentConfigV1) (*PersistentConfigV1, bool) {
	changed := false

	// if cfg.Version < CurrentVersion {
	// 	if cfg.SSLMode == "" {
	// 		cfg.SSLMode = "disable"
	// 	}
	// 	cfg.Version = 1
	// 	changed = true
	// }

	if cfg.Version < CurrentVersion {
		cfg.Version = CurrentVersion
		changed = true
	}

	return cfg, changed
}

func (p *PersistentConfigV1) Save() error {
	dir := filepath.Dir(ConfigPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigPath, data, 0o600)
}

func RemovePersistentConfig() error {
	err := os.Remove(ConfigPath)

	if err != nil {
		return err
	}

	return nil
}

func ValidatePersistentConfig(cfg *PersistentConfigV1) error {

	c := cfg.Connection

	if c.Host == "" {
		return errors.New("host is required")
	}

	if c.Port <= 0 {
		return errors.New("port is required")
	}
	if c.User == "" {
		return errors.New("user is required")
	}
	if c.Database == "" {
		return errors.New("database is required")
	}

	if c.Password == "" {
		return errors.New("password is required")
	}

	// SSLMode can also be empty — Postgres will decide the fallback itself

	return nil
}
