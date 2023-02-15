package log

import (
	"context"
	"fmt"

	"go.uber.org/fx/fxevent"

	"go.uber.org/zap"
)

// Logger provides functions for logging at various levels.
type Logger interface {
	// PrepareCtx adds values to the context for logging if necessary.
	PrepareCtx(ctx context.Context, projectID string) context.Context
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	FxLogger() fxevent.Logger
}

type zapLogger struct {
	l *zap.Logger
}

// New provides a new instance of Logger
func New() (Logger, error) {
	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return NewZapLogger(l), nil
}

func NewZapLogger(log *zap.Logger) Logger {
	return &zapLogger{l: log.WithOptions(zap.AddCallerSkip(1))}
}

func (z zapLogger) PrepareCtx(ctx context.Context, _ string) context.Context {
	return ctx
}

func (z zapLogger) Info(_ context.Context, args ...interface{}) {
	z.l.Info(fmt.Sprint(args...))
}

func (z zapLogger) Infof(_ context.Context, format string, args ...interface{}) {
	z.l.Info(fmt.Sprintf(format, args...))
}

func (z zapLogger) Warning(_ context.Context, args ...interface{}) {
	z.l.Warn(fmt.Sprint(args...))
}

func (z zapLogger) Warningf(_ context.Context, format string, args ...interface{}) {
	z.l.Warn(fmt.Sprintf(format, args...))
}

func (z zapLogger) Error(_ context.Context, args ...interface{}) {
	z.l.Error(fmt.Sprint(args...))
}

func (z zapLogger) Errorf(_ context.Context, format string, args ...interface{}) {
	z.l.Info(fmt.Sprintf(format, args...))
}

func (z zapLogger) FxLogger() fxevent.Logger {
	return &fxevent.ZapLogger{Logger: z.l}
}
