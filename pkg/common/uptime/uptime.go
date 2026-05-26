package uptime

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Report every 10 seconds.
const reportInterval = time.Second * 10

var (
	clk   = clock.New()
	start = clk.Now()
)

func Uptime() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func reportMetrics(ctx context.Context, interval time.Duration, m telemetry.Metrics) {
	_ = "STUB: not implemented"
	return
}

func ReportMetrics(ctx context.Context, metrics telemetry.Metrics) {
	_ = "STUB: not implemented"
	return
}
