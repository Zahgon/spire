package agent

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartRotateAgentSVIDCall return metric for Agent's SVID
// Rotation.
func StartRotateAgentSVIDCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartReattestAgentCall return metric for Agent's
// Reattestation.
func StartReattestAgentCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
