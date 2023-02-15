package log

import (
	"context"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Initialize logger with fake logger to avoid any issues.
var logger Logger = FakeLogger{}

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger = NewZapLogger(l)
}

func PrepareCtx(ctx context.Context, projectID string) context.Context {
	return logger.PrepareCtx(ctx, projectID)
}

func Info(ctx context.Context, args ...interface{}) {
	logger.Info(ctx, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.Infof(ctx, format, args...)
}

func Warning(ctx context.Context, args ...interface{}) {
	logger.Warning(ctx, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	logger.Warningf(ctx, format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	logger.Error(ctx, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.Errorf(ctx, format, args...)
}

func FxLogger() fxevent.Logger {
	return logger.FxLogger()
}
