package endpoints

import (
	"context"

	"github.com/spiffe/spire/pkg/common/api/middleware"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

func withPerServiceConnectionMetrics(metrics telemetry.Metrics) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

type connectionMetrics struct {
	metrics                   telemetry.Metrics
	workloadAPIConns          int32
	sdsAPIConns               int32
	debugAPIConns             int32
	delegatedIdentityAPIConns int32
}

func (m *connectionMetrics) Preprocess(ctx context.Context, _ string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Intentionally not emitting metrics for health and reflection services

func (m *connectionMetrics) Postprocess(ctx context.Context, _ string, _ bool, _ error) {
	_ = "STUB: not implemented"
	return
}

// Intentionally not emitting metrics for health and reflection services
