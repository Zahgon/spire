package rpccontext

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, log logrus.FieldLogger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Logger(ctx context.Context) logrus.FieldLogger {
	_ = "STUB: not implemented"
	return *new(logrus.FieldLogger)
}
