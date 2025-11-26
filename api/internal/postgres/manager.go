package postgres

import (
	"context"
	"dashboard/api/internal/config"
	"fmt"
	"log/slog"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type ConnectionStatus string

const (
	StatusDisconnected ConnectionStatus = "disconnected"
	StatusConnecting   ConnectionStatus = "connecting"
	StatusConnected    ConnectionStatus = "connected"
	StatusError        ConnectionStatus = "error"
)

type Manager struct {
	config   config.AppConfig
	logger   *slog.Logger
	mu       sync.RWMutex
	sqlxDB   *sqlx.DB
	retries  int
	retryDur time.Duration
	status   ConnectionStatus
}

func New(config config.AppConfig, logger *slog.Logger) *Manager {

	m := &Manager{
		config:   config,
		logger:   logger,
		retries:  3,
		retryDur: 1 * time.Second,
	}

	if config.PersistentConfig != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := m.UpdateConnection(ctx, config.PersistentConfig.Connection); err != nil {
			logger.Error("failed to connect from sqlc", "error", err)
		} else {
			logger.Info("✅ connected to Postgres from config")
		}
	}

	return m
}

func (m *Manager) UpdateConnection(ctx context.Context, newConn config.Connection) error {

	if newConn == (config.Connection{}) {
		m.mu.Lock()

		if m.sqlxDB != nil {
			m.sqlxDB.Close()
		}
		m.sqlxDB = nil
		if m.config.PersistentConfig != nil {
			m.config.PersistentConfig.Connection = config.Connection{}
		}

		m.status = StatusDisconnected

		m.mu.Unlock()

		return nil
	}

	m.mu.Lock()
	m.status = StatusConnecting
	m.mu.Unlock()

	dsn := connectionDSN(newConn)
	var db *sqlx.DB
	var err error

	for i := 0; i <= m.retries; i++ {
		db, err = sqlx.Open("pgx", dsn)

		if err != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(m.retryDur):
				continue
			}
		}

		pingErr := db.PingContext(ctx)
		if pingErr == nil {
			break
		}

		db.Close()
		err = pingErr

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(m.retryDur):
		}
	}

	if err != nil {
		m.mu.Lock()
		m.status = StatusDisconnected
		m.mu.Unlock()
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	m.mu.Lock()

	if m.sqlxDB != nil {
		m.sqlxDB.Close()
	}

	m.sqlxDB = db

	if m.config.PersistentConfig != nil {
		m.config.PersistentConfig.Connection = newConn
		m.config.PersistentConfig.Save()
	} else {
		m.config.PersistentConfig = &config.PersistentConfigV1{
			Version:    config.CurrentVersion,
			Connection: newConn,
		}

		m.config.PersistentConfig.Save()
	}

	m.status = StatusConnected

	m.mu.Unlock()

	return nil
}

func (m *Manager) SQLX() (*sqlx.DB, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.sqlxDB == nil {
		return nil, ErrNotConnected
	}
	return m.sqlxDB, nil
}

func (m *Manager) HealthCheck(ctx context.Context) error {
	m.mu.RLock()
	db := m.sqlxDB
	m.mu.RUnlock()
	if db == nil {
		return ErrNotConnected
	}
	return db.PingContext(ctx)
}

func (m *Manager) Status() ConnectionStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.status
}

func (m *Manager) IsConnected() bool {
	return m.Status() == StatusConnected
}

func (m *Manager) Connecting() bool {
	return m.Status() == StatusConnecting
}
