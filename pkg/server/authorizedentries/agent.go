package authorizedentries

type agentRecord struct {
	ID string

	// ExpiresAt is seconds since unix epoch. Using instead of time.Time for
	// reduced memory usage and better cache locality.
	ExpiresAt int64

	Selectors selectorSet
}

func agentRecordByID(a, b agentRecord) bool { _ = "STUB: not implemented"; return false }

func agentRecordByExpiresAt(a, b agentRecord) bool { _ = "STUB: not implemented"; return false }
