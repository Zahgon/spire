package agent

import "github.com/spiffe/spire/pkg/common/telemetry"

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartNodeAttestorNewSVIDCall return metric
// for agent node attestor call to get new SVID
// for the agent
func StartNodeAttestorNewSVIDCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
