package middleware

import (
	"context"

	"github.com/spiffe/spire/pkg/common/telemetry"
)

// WithMetrics adds per-call metrics to each RPC call. It emits both a call
// counter and sample with the call timing. RPC handlers can add their own
// labels to be attached to the per-call metrics via the
// rpccontext.AddMetricsLabel function. If unset, it also provides name
// metadata on to the handler context.
func WithMetrics(metrics telemetry.Metrics) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

type metricsMiddleware struct {
	metrics telemetry.Metrics
}

func (m metricsMiddleware) Preprocess(ctx context.Context, fullMethod string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (m metricsMiddleware) Postprocess(ctx context.Context, _ string, _ bool, rpcErr error) {
	_ = "STUB: not implemented"
	return
}
