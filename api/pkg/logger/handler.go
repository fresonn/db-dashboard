package logger

import (
	"context"
	"dashboard/api/pkg/ctxkeys"
	"log/slog"
)

type handlerMiddlware struct {
	next slog.Handler
}

func withMiddleware(next slog.Handler) *handlerMiddlware {
	return &handlerMiddlware{next: next}
}

func (h *handlerMiddlware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}

func (h *handlerMiddlware) Handle(ctx context.Context, rec slog.Record) error {

	if requestID, ok := ctxkeys.RequestID(ctx); ok {
		if requestID != "" {
			rec.AddAttrs(slog.String("request_id", requestID))
		}
	}

	return h.next.Handle(ctx, rec)
}

func (h *handlerMiddlware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handlerMiddlware{next: h.next.WithAttrs(attrs)}
}

func (h *handlerMiddlware) WithGroup(name string) slog.Handler {
	return &handlerMiddlware{next: h.next.WithGroup(name)}
}
