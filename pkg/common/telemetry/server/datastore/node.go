package datastore

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartCountNodeCall return metric
// for server's datastore, on counting nodes.
func StartCountNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartCreateNodeCall return metric
// for server's datastore, on creating a node.
func StartCreateNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartDeleteNodeCall return metric
// for server's datastore, on deleting a node.
func StartDeleteNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartPruneAttestedExpiredNodes return metric
// for server's datastore, on pruning expired attested nodes.
func StartPruneAttestedExpiredNodes(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartFetchNodeCall return metric
// for server's datastore, on fetching a node.
func StartFetchNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartListNodeCall return metric
// for server's datastore, on listing nodes.
func StartListNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartGetNodeSelectorsCall return metric
// for server's datastore, on getting selectors for a node.
func StartGetNodeSelectorsCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartListNodeSelectorsCall return metric
// for server's datastore, on getting selectors for a node.
func StartListNodeSelectorsCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartSetNodeSelectorsCall return metric
// for server's datastore, on setting selectors for a node.
func StartSetNodeSelectorsCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartUpdateNodeCall return metric
// for server's datastore, on updating a node.
func StartUpdateNodeCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
