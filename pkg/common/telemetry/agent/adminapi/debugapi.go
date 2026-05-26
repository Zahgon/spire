package adminapi

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Counters (literal increments, not call counters)

// IncrDebugAPIConnectionCounter indicate Debug
// API connection (some connection is made, running total count)
func IncrDebugAPIConnectionCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// SetDebugAPIConnectionGauge sets the number of active Debug API connections
func SetDebugAPIConnectionGauge(m telemetry.Metrics, connections int32) {
	_ = "STUB: not implemented"
	return
}

// End Counters
