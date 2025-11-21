package logger

import "log/slog"

func WithScopeLogger(logger *slog.Logger, domain string) *slog.Logger {
	return logger.With("scope_name", domain)
}
