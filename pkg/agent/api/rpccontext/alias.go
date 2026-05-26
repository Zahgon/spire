package rpccontext

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/api"
)

func WithLogger(ctx context.Context, log logrus.FieldLogger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Logger(ctx context.Context) logrus.FieldLogger {
	_ = "STUB: not implemented"
	return *new(logrus.FieldLogger)
}

func WithCallCounter(ctx context.Context, counter api.CallCounter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CallCounter(ctx context.Context) api.CallCounter {
	_ = "STUB: not implemented"
	return *new(api.CallCounter)
}

func AddMetricsLabel(ctx context.Context, name, value string) { _ = "STUB: not implemented"; return }

func WithNames(ctx context.Context, names api.Names) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Names(ctx context.Context) (api.Names, bool) {
	_ = "STUB: not implemented"
	return *new(api.Names), false
}
