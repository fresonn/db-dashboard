package logger

import (
	"dashboard/api/internal/config"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func New(appCfg config.AppConfig) *slog.Logger {

	var handler slog.Handler

	switch appCfg.Env.Runtime {

	case config.LiveRuntime:

		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelInfo,
			TimeFormat: time.DateTime,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

				if a.Key == "error" {
					v, ok := a.Value.Any().(error)
					if !ok {
						return a
					}
					return tint.Err(v)
				}

				return a
			},
		})

	case config.DevRuntime:

		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

				if a.Key == "error" {
					v, ok := a.Value.Any().(error)
					if !ok {
						return a
					}
					return tint.Err(v)
				}

				return a
			},
		})

	case config.TestRuntime:

		handler = slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
			Level: slog.LevelError + 10,
		})
	}

	handler = withMiddleware(handler)

	log := slog.New(handler)

	slog.SetDefault(log)

	return log
}
