package authorizedentries

type aliasRecord struct {
	// EntryID is the ID of the registration entry that defines this node
	// alias.
	EntryID string

	// AliasID is the SPIFFE ID of nodes that match this alias.
	AliasID string

	// Selector is the specific selector we use to fan out to this record
	// during the crawl.
	Selector Selector

	// AllSelectors is here out of convenience to verify that the agent
	// possesses a superset of the alias's selectors and is therefore
	// authorized for the alias.
	AllSelectors selectorSet
}

func aliasRecordByEntryID(a, b aliasRecord) bool { _ = "STUB: not implemented"; return false }

func aliasRecordBySelector(a, b aliasRecord) bool { _ = "STUB: not implemented"; return false }
