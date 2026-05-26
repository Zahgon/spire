package workloadapi

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartAttestationCall return metric
// for agent's Workload API Attestor for overall attestation
func StartAttestationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartAttestorCall return metric
// for agent's Workload API Attestor for a specific attestor
func StartAttestorCall(m telemetry.Metrics, aType string) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters

// Counters (literal increments, not call counters)

// IncrConnectionCounter indicate Workload
// API connection (some connection is made, running total count)
func IncrConnectionCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// SetConnectionTotalGauge sets the number of active Workload API connections
func SetConnectionTotalGauge(m telemetry.Metrics, connections int32) {
	_ = "STUB: not implemented"
	return
}

// End Counters

// Add Samples (metric on count of some object, entries, event...)

// AddDiscoveredSelectorsSample count of discovered selectors
// during an agent Workload Attest call
func AddDiscoveredSelectorsSample(m telemetry.Metrics, count float32) {
	_ = "STUB: not implemented"
	return
}

// End Add Samples
