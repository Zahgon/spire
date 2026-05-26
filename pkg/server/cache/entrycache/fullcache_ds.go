package entrycache

import (
	"context"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

var (
	_ EntryIterator = (*entryIteratorDS)(nil)
	_ AgentIterator = (*agentIteratorDS)(nil)
	// 10,000 was chosen to balance # of requests sent to spire-db and timeouts to the database.
	// Too large of a page size incurs large latencies while listing registrations.
	// Too small incurs too many requests sent to the DB.
	// Pagination only affects large entry counts within spire-db. Smaller deployments of spire-db should remain
	// unaffected as the latency spent sending multiple requests is more expensive than the call itself.
	listEntriesRequestPageSize int32 = 10000
)

// BuildFromDataStore builds a Cache using the provided datastore as the data source
func BuildFromDataStore(ctx context.Context, trustDomain string, ds datastore.DataStore) (*FullEntryCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type entryIteratorDS struct {
	ds              datastore.DataStore
	entries         []*types.Entry
	next            int
	err             error
	paginationToken string
}

func makeEntryIteratorDS(ds datastore.DataStore) EntryIterator {
	_ = "STUB: not implemented"
	return *new(EntryIterator)
}

func (it *entryIteratorDS) Next(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (it *entryIteratorDS) filterEntries(in []*common.RegistrationEntry) []*common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}

// Filter out entries with invalid SPIFFE IDs. Operators are notified
// that they are ignored on server startup (see
// pkg/server/scanentries.go)

func (it *entryIteratorDS) Entry() *types.Entry { _ = "STUB: not implemented"; return nil }

func (it *entryIteratorDS) Err() error { _ = "STUB: not implemented"; return nil }

type agentIteratorDS struct {
	ds     datastore.DataStore
	agents []Agent
	next   int
	err    error
}

func makeAgentIteratorDS(ds datastore.DataStore) AgentIterator {
	_ = "STUB: not implemented"
	return *new(AgentIterator)
}

func (it *agentIteratorDS) Next(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (it *agentIteratorDS) Agent() Agent { _ = "STUB: not implemented"; return *new(Agent) }

func (it *agentIteratorDS) Err() error {
	_ = "STUB: not implemented"

	// Fetches all agent selectors from the datastore and stores them in the iterator.
	return nil
}

func (it *agentIteratorDS) fetchAgents(ctx context.Context) ([]Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
