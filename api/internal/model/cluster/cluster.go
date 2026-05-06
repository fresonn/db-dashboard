package cluster

import (
	"dashboard/api/internal/infra/postgres"
	"time"
)

type NewConnection struct {
	Host     string  `json:"host" validate:"required,hostname"`
	Port     int     `json:"port" validate:"gt=0"`
	User     string  `json:"user" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Database *string `json:"database,omitempty"`
	SSLMode  string  `json:"sslmode"`
}

type PostgresVersion struct {
	Version  string `json:"version"`
	Compiler string `json:"compiler"`
	BitDepth string `json:"bitDepth"`
}

type PostgresUptime struct {
	StartedAt time.Time `json:"startedAt" db:"cluster_started_at"`
}

type Status struct {
	CurrentUser      *string
	CurrentDatabase  *string
	ConnectionStatus postgres.ConnectionStatus
}
