package logger

import (
	"context"
	"log/slog"
	"os"
)

type SlogLogger struct {
	log *slog.Logger
}

func NewSlogLogger(env string) Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return &SlogLogger{
		log: slog.New(handler),
	}
}

func (l *SlogLogger) withContext(ctx context.Context, attrs []any) []any {
	if rid, ok := RequestIDFromContext(ctx); ok {
		return append(attrs, "request_id", rid)
	}
	return attrs
}

func (l *SlogLogger) Info(ctx context.Context, msg string, attrs ...any) {
	l.log.Info(msg, l.withContext(ctx, attrs)...)
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, attrs ...any) {
	l.log.Warn(msg, l.withContext(ctx, attrs)...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, attrs ...any) {
	l.log.Error(msg, l.withContext(ctx, attrs)...)
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, attrs ...any) {
	l.log.Debug(msg, l.withContext(ctx, attrs)...)
}
