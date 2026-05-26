package rpccontext

import (
	"context"

	"github.com/spiffe/spire/pkg/common/api"
)

type callCounterKey struct{}

func WithCallCounter(ctx context.Context, counter api.CallCounter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CallCounter(ctx context.Context) api.CallCounter {
	_ = "STUB: not implemented"
	return *new(api.CallCounter)
}

func AddMetricsLabel(ctx context.Context, name, value string) { _ = "STUB: not implemented"; return }
