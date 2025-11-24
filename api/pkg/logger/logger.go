package logger

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type LoggerVariant string

const (
	LiveLoggerVariant LoggerVariant = "live"
	TestLoggerVariant LoggerVariant = "test"
)

func New(env LoggerVariant) *slog.Logger {

	var handler slog.Handler

	switch env {

	case LiveLoggerVariant:

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

	case TestLoggerVariant:
		handler = slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
			Level: slog.LevelError + 10,
		})
	}

	handler = withMiddleware(handler)

	log := slog.New(handler)

	slog.SetDefault(log)

	return log
}
