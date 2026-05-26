package agent

import (
	"github.com/spiffe/spire/pkg/agent/client"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

const (
	CacheTypeWorkload  = "workload"
	CacheTypeSVIDStore = "svid_store"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartManagerFetchEntriesUpdatesCall returns metric for when agent's
// synchronization manager fetching latest entries information
// from server
func StartManagerFetchEntriesUpdatesCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartManagerFetchSVIDsUpdatesCall returns metric for when agent's
// synchronization manager fetching latest SVIDs information
// from server
func StartManagerFetchSVIDsUpdatesCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// End Call Counters

// Add Samples (metric on count of some object, entries, event...)

// AddCacheManagerExpiredSVIDsSample count of expiring SVIDs according to
// agent cache manager
func AddCacheManagerExpiredSVIDsSample(m telemetry.Metrics, cacheType string, count float32) {
	_ = "STUB: not implemented"
	return
}

// AddCacheManagerOutdatedSVIDsSample count of SVIDs with outdated attributes
// according to agent cache manager
func AddCacheManagerOutdatedSVIDsSample(m telemetry.Metrics, cacheType string, count float32) {
	_ = "STUB: not implemented"
	return
}

// AddCacheManagerTaintedX509SVIDsSample count of tainted X509-SVIDs according to
// agent cache manager
func AddCacheManagerTaintedX509SVIDsSample(m telemetry.Metrics, cacheType string, count float32) {
	_ = "STUB: not implemented"
	return
}

// AddCacheManagerTaintedJWTSVIDsSample count of tainted JWT-SVIDs according to
// agent cache manager
func AddCacheManagerTaintedJWTSVIDsSample(m telemetry.Metrics, cacheType string, count float32) {
	_ = "STUB: not implemented"
	return
}

// End Add Samples

func SetSyncStats(m telemetry.Metrics, stats client.SyncStats) { _ = "STUB: not implemented"; return }
