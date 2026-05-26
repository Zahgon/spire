package adminapi

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartFirstX509SVIDUpdateLatency returns Latency metric
// for SubscribeToX509SVIDs API fetching the first update from cache.
func StartFirstX509SVIDUpdateLatency(m telemetry.Metrics) *telemetry.Latency {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters

// Counters (literal increments, not call counters)

// IncrDelegatedIdentityAPIConnectionCounter indicate Delegated Identity
// API connection (some connection is made, running total count)
func IncrDelegatedIdentityAPIConnectionCounter(m telemetry.Metrics) {
	_ = "STUB: not implemented"
	return
}

// SetDelegatedIdentityAPIConnectionGauge sets the number of active connections
func SetDelegatedIdentityAPIConnectionGauge(m telemetry.Metrics, connections int32) {
	_ = "STUB: not implemented"
	return
}

// End Counters
