package server

import "github.com/spiffe/spire/pkg/common/telemetry"

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartRegistrationManagerPruneEntryCall returns metric for
// for server registration manager entry pruning
func StartRegistrationManagerPruneEntryCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
