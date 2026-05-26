package datastore

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// StartSetCAJournal return metric for server's datastore, on setting a CA
// journal.
func StartSetCAJournal(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartFetchCAJournal return metric
// for server's datastore, on fetching a CA journal.
func StartFetchCAJournal(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartPruneCAJournalsCall return metric for server's datastore, on pruning CA
// journals.
func StartPruneCAJournalsCall(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// StartListCAJournalsForTesting return metric
// for server's datastore, on listing CA journals for testing.
func StartListCAJournalsForTesting(m telemetry.Metrics) *telemetry.CallCounter {
	_ = "STUB: not implemented"
	return nil
}
