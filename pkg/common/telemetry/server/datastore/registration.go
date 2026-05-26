package datastore

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartCountRegistrationCall return metric
// for server's datastore, on counting registrations.
func StartCountRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartCreateRegistrationCall return metric
// for server's datastore, on creating a registration.
func StartCreateRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartDeleteRegistrationCall return metric
// for server's datastore, on deleting a registration.
func StartDeleteRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartFetchRegistrationCall return metric
// for server's datastore, on creating a registration.
func StartFetchRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartListRegistrationCall return metric
// for server's datastore, on listing registrations.
func StartListRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartPruneRegistrationCall return metric
// for server's datastore, on pruning registrations.
func StartPruneRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartUpdateRegistrationCall return metric
// for server's datastore, on updating a registration.
func StartUpdateRegistrationCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters
