package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/rchmachina/rach-fw/internal/utils/helper"
)

type Field struct {
	Key   string
	Value any
}

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)

	With(fields ...Field) Logger
	WithCtx(ctx context.Context) Logger
}
type slogLogger struct {
	l *slog.Logger
}

func NewSlogLogger(component string, isEnv bool) Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{}

	if isEnv {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	base := slog.New(handler)

	return &slogLogger{
		l: base.With("component", component),
	}
}

func (s *slogLogger) Info(msg string, fields ...Field) {
	s.l.Info(msg, toArgs(fields)...)
}

func (s *slogLogger) Error(msg string, fields ...Field) {
	s.l.Error(msg, toArgs(fields)...)
}

func (s *slogLogger) Debug(msg string, fields ...Field) {
	s.l.Debug(msg, toArgs(fields)...)
}

func (s *slogLogger) Warn(msg string, fields ...Field) {
	s.l.Warn(msg, toArgs(fields)...)
}

func (s *slogLogger) With(fields ...Field) Logger {
	return &slogLogger{
		l: s.l.With(toArgs(fields)...),
	}
}

func toArgs(fields []Field) []any {
	args := make([]any, 0, len(fields)*2)

	for _, f := range fields {
		args = append(args, f.Key, f.Value)
	}
	return args
}

func ToField(k string, value any) Field {
	return Field{
		Key:   k,
		Value: value,
	}
}

func (s *slogLogger) WithCtx(ctx context.Context) Logger {
	reqID := helper.GetRequestID(ctx)

	if reqID == "" {
		return s
	}

	return &slogLogger{
		l: s.l.With("req_id", reqID),
	}
}
