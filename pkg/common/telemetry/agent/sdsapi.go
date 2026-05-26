package agent

import "github.com/spiffe/spire/pkg/common/telemetry"

// Counters (literal increments, not call counters)

// IncrSDSAPIConnectionCounter indicate SDS
// API connection (some connection is made, running total count)
func IncrSDSAPIConnectionCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// SetSDSAPIConnectionTotalGauge sets the number of active SDS connections
func SetSDSAPIConnectionTotalGauge(m telemetry.Metrics, connections int32) {
	_ = "STUB: not implemented"
	return
}

// End Counters
