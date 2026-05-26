package entrycache

import (
	"context"
	"sync"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/api"
)

var (
	seenSetPool = sync.Pool{
		New: func() any {
			return make(seenSet)
		},
	}

	stringSetPool = sync.Pool{
		New: func() any {
			return make(stringSet)
		},
	}
)

var _ Cache = (*FullEntryCache)(nil)

// Cache contains a snapshot of all registration entries and Agent selectors from the data source
// at a particular moment in time.
type Cache interface {
	LookupAuthorizedEntries(agentID spiffeid.ID, entries map[string]struct{}) map[string]api.ReadOnlyEntry
	GetAuthorizedEntries(agentID spiffeid.ID) []api.ReadOnlyEntry
}

// Selector is a key-value attribute of a node or workload.
type Selector struct {
	// Type is the type of the selector.
	Type string
	// Value is the value of the selector.
	Value string
}

// EntryIterator is used to iterate through registration entries from a data source.
// The usage pattern of the iterator is as follows:
//
//	for it.Next() {
//	    entry := it.Entry()
//	    // process entry
//	}
//
//	if it.Err() {
//	    // handle error
//	}
type EntryIterator interface {
	// Next returns true if there are any remaining registration entries in the data source and returns false otherwise.
	Next(ctx context.Context) bool
	// Entry returns the next entry from the data source.
	Entry() *types.Entry
	// Err returns an error encountered when attempting to process entries from the data source.
	Err() error
}

// AgentIterator is used to iterate through Agent selectors from a data source.
// The usage pattern of the iterator is as follows:
//
//	for it.Next() {
//	    agent := it.Agent()
//	    // process agent
//	}
//
//	if it.Err() {
//	    // handle error
//	}
type AgentIterator interface {
	// Next returns true if there are any remaining agents in the data source and returns false otherwise.
	Next(ctx context.Context) bool
	// Agent returns the next agent from the data source.
	Agent() Agent
	// Err returns an error encountered when attempting to process agents from the data source.
	Err() error
}

// Agent represents the association of selectors to an agent SPIFFE ID.
type Agent struct {
	// ID is the Agent's SPIFFE ID.
	ID spiffeid.ID
	// Selectors is the Agent's selectors.
	Selectors []*types.Selector
}

type FullEntryCache struct {
	aliases map[string][]aliasEntry
	entries map[string][]*types.Entry
}

type selectorSet map[Selector]struct{}
type seenSet map[string]struct{}
type stringSet map[string]struct{}

type aliasEntry struct {
	id    string
	entry *types.Entry
}

// Build queries the data source for all registration entries and Agent selectors and builds an in-memory
// representation of the data that can be used for efficient lookups.
func Build(ctx context.Context, trustDomain string, entryIter EntryIterator, agentIter AgentIterator) (*FullEntryCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// track which aliases we've evaluated so far to make sure we don't
// add one twice.

func (c *FullEntryCache) LookupAuthorizedEntries(agentID spiffeid.ID, requestedEntries map[string]struct{}) map[string]api.ReadOnlyEntry {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthorizedEntries gets all authorized registration entries for a given Agent SPIFFE ID.
func (c *FullEntryCache) GetAuthorizedEntries(agentID spiffeid.ID) []api.ReadOnlyEntry {
	_ = "STUB: not implemented"
	return nil
}

// Crawl the list of registration entries calling the visit function on all of them.
// visit(entry) returns a boolean indicating if we should continue iterating (if true)
// or if we should terminate the crawl (if false).
func (c *FullEntryCache) crawl(parentID string, seen map[string]struct{}, visit func(*types.Entry) bool) {
	_ = "STUB: not implemented"
	return
}

func selectorSetFromProto(selectors []*types.Selector) selectorSet {
	_ = "STUB: not implemented"
	return *new(selectorSet)
}

func allocSeenSet() seenSet { _ = "STUB: not implemented"; return *new(seenSet) }

func freeSeenSet(set seenSet) { _ = "STUB: not implemented"; return }

func clearSeenSet(set seenSet) { _ = "STUB: not implemented"; return }

func allocStringSet() stringSet { _ = "STUB: not implemented"; return *new(stringSet) }

func freeStringSet(set stringSet) { _ = "STUB: not implemented"; return }

func clearStringSet(set stringSet) { _ = "STUB: not implemented"; return }

func isSubset(sub, whole selectorSet) bool { _ = "STUB: not implemented"; return false }
