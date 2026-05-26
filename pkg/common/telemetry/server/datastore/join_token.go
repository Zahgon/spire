package datastore

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartCreateJoinTokenCall return metric
// for server's datastore, on creating a join token.
func StartCreateJoinTokenCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartDeleteJoinTokenCall return metric
// for server's datastore, on deleting a join token.
func StartDeleteJoinTokenCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartFetchJoinTokenCall return metric
// for server's datastore, on fetching a join token.
func StartFetchJoinTokenCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartPruneJoinTokenCall return metric
// for server's datastore, on pruning join tokens.
func StartPruneJoinTokenCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
