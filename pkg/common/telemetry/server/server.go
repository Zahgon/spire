package server

import "github.com/spiffe/spire/pkg/common/telemetry"

// SetEntryDeletedGauge emits a gauge with the number of entries that will
// be deleted in the entry cache.
func SetEntryDeletedGauge(m telemetry.Metrics, deleted int) { _ = "STUB: not implemented"; return }

// SetAgentsByIDCacheCountGauge emits a gauge with the number of agents by ID that are
// currently in the node cache.
func SetAgentsByIDCacheCountGauge(m telemetry.Metrics, size int) { _ = "STUB: not implemented"; return }

// SetAgentsByExpiresAtCacheCountGauge emits a gauge with the number of agents by expiresAt that are
// currently in the node cache.
func SetAgentsByExpiresAtCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}

// SetSkippedNodeEventIDsCacheCountGauge emits a gauge with the number of entries that are
// currently in the skipped-node events cache.
func SetSkippedNodeEventIDsCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}

// SetNodeAliasesByEntryIDCacheCountGauge emits a gauge with the number of Node Aliases by EntryID that are
// currently in the entry cache.
func SetNodeAliasesByEntryIDCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}

// SetNodeAliasesBySelectorCacheCountGauge emits a gauge with the number of Node Aliases by Selector that are
// currently in the entry cache.
func SetNodeAliasesBySelectorCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}

// SetEntriesByEntryIDCacheCountGauge emits a gauge with the number of entries by entryID that are
// currently in the entry cache.
func SetEntriesByEntryIDCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}

// SetSkippedEntryEventIDsCacheCountGauge emits a gauge with the number of entries that are
// currently in the skipped-entry events cache.
func SetSkippedEntryEventIDsCacheCountGauge(m telemetry.Metrics, size int) {
	_ = "STUB: not implemented"
	return
}
