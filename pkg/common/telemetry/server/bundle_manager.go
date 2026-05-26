package server

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Counters (literal increments, not call counters)

// IncrBundleManagerUpdateFederatedBundleCounter indicate
// the number of updating federated bundle by bundle manager
func IncrBundleManagerUpdateFederatedBundleCounter(m telemetry.Metrics, trustDomain string) {
	_ = "STUB: not implemented"
	return
}

// End Counters

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartBundleManagerFetchFederatedBundleCall return metric for Server's federated bundle fetch.
func StartBundleManagerFetchFederatedBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
