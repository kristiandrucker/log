package log

import (
	"context"

	"go.uber.org/fx/fxevent"
)

type FakeLogger struct{}

func (FakeLogger) PrepareCtx(ctx context.Context, _ string) context.Context {
	return ctx
}
func (FakeLogger) Info(_ context.Context, _ ...interface{})               {}
func (FakeLogger) Infof(_ context.Context, _ string, _ ...interface{})    {}
func (FakeLogger) Warning(_ context.Context, _ ...interface{})            {}
func (FakeLogger) Warningf(_ context.Context, _ string, _ ...interface{}) {}
func (FakeLogger) Error(_ context.Context, _ ...interface{})              {}
func (FakeLogger) Errorf(_ context.Context, _ string, _ ...interface{})   {}
func (FakeLogger) FxLogger() fxevent.Logger                               { return &fxevent.NopLogger }
