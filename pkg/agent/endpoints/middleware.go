package endpoints

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/api/middleware"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

const (
	workloadAPIMethodPrefix = "/SpiffeWorkloadAPI/"
)

func Middleware(log logrus.FieldLogger, metrics telemetry.Metrics) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func addWatcherPID(ctx context.Context, _ string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func verifySecurityHeader(ctx context.Context, fullMethod string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func isWorkloadAPIMethod(fullMethod string) bool { _ = "STUB: not implemented"; return false }

func hasSecurityHeader(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// discardAgentCallMetrics prevents RPC metrics from being emitted for calls
// made by the agent itself (e.g. health check loopback calls to the Workload
// API). This runs in Postprocess before the metrics middleware finalizes the
// call counter, so discarding here prevents the counter from emitting.
func discardAgentCallMetrics(ctx context.Context, _ string, _ bool, _ error) {
	_ = "STUB: not implemented"
	return
}
