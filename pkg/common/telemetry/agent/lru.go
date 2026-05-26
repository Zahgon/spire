package agent

import "github.com/spiffe/spire/pkg/common/telemetry"

func IncrementEntriesAdded(m telemetry.Metrics, entriesAdded int) {
	_ = "STUB: not implemented"
	return
}

func IncrementEntriesUpdated(m telemetry.Metrics, entriesUpdated int) {
	_ = "STUB: not implemented"
	return
}

func IncrementEntriesRemoved(m telemetry.Metrics, entriesRemoved int) {
	_ = "STUB: not implemented"
	return
}

func SetEntriesMapSize(m telemetry.Metrics, recordMapSize int) { _ = "STUB: not implemented"; return }

func SetSVIDMapSize(m telemetry.Metrics, svidMapSize int) { _ = "STUB: not implemented"; return }
