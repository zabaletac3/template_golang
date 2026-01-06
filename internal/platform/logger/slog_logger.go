package logger

import (
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

	l := slog.New(handler)

	return &SlogLogger{log: l}
}

func (l *SlogLogger) Info(msg string, attrs ...any) {
	l.log.Info(msg, attrs...)
}

func (l *SlogLogger) Warn(msg string, attrs ...any) {
	l.log.Warn(msg, attrs...)
}

func (l *SlogLogger) Error(msg string, attrs ...any) {
	l.log.Error(msg, attrs...)
}

func (l *SlogLogger) Debug(msg string, attrs ...any) {
	l.log.Debug(msg, attrs...)
}
