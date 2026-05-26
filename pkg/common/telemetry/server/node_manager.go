package server

import "github.com/spiffe/spire/pkg/common/telemetry"

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartNodeManagerPruneAttestedExpiredNodesCall returns metric for
// for expired agent pruning
func StartNodeManagerPruneAttestedExpiredNodesCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
