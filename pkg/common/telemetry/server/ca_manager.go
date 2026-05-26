package server

import (
	"time"

	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartCAManagerPruneBundleCall returns metric for
// for server CA manager bundle pruning
func StartCAManagerPruneBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartServerCAManagerPrepareJWTKeyCall return metric for
// Server CA Manager preparing a JWT Key
func StartServerCAManagerPrepareJWTKeyCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartServerCAManagerPrepareWITKeyCall return metric for
// Server CA Manager preparing a WIT Key
func StartServerCAManagerPrepareWITKeyCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartServerCAManagerPrepareX509CACall return metric for
// Server CA Manager preparing an X509 CA
func StartServerCAManagerPrepareX509CACall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters

// Gauge (remember previous value set)

// SetX509CARotateGauge set gauge for X509 CA rotation,
// expiration time and TTL of CA for a specific TrustDomain
func SetX509CARotateGauge(m telemetry.Metrics, trustDomain string, expiration, now time.Time) {
	_ = "STUB: not implemented"
	return
}

// End Gauge

// Counters (literal increments, not call counters)

// IncrActivateJWTKeyManagerCounter indicate activation
// of JWT Key manager
func IncrActivateJWTKeyManagerCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrActivateX509CAManagerCounter indicate activation
// of X509 CA manager
func IncrActivateX509CAManagerCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrActivateWITKeyManagerCounter indicate activation
// of WIT Key manager
func IncrActivateWITKeyManagerCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrManagerPrunedBundleCounter indicate manager
// having pruned a bundle
func IncrManagerPrunedBundleCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrServerCASignJWTSVIDCounter indicate Server CA
// signed a JWT SVID.
func IncrServerCASignJWTSVIDCounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrServerCASignX509CACounter indicate Server CA
// signed an X509 CA SVID.
func IncrServerCASignX509CACounter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// IncrServerCASignX509Counter indicate Server CA
// signed an X509 SVID.
func IncrServerCASignX509Counter(m telemetry.Metrics) { _ = "STUB: not implemented"; return }

// End Counters
