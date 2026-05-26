package authorizedentries

import (
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/google/btree"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/api"
)

const (
	// We can tweak these degrees to try and get optimal L1 cache use, but
	// it's probably not worth it unless we have benchmarks showing that it
	// is a problem at scale in production. Initial benchmarking by myself
	// at similar scale to some of our bigger, existing deployments didn't
	// seem to yield much difference. As such, these values are probably an
	// ok jumping off point.
	agentRecordDegree = 32
	aliasRecordDegree = 32
)

type Selector struct {
	Type  string
	Value string
}

func (s Selector) String() string { _ = "STUB: not implemented"; return "" }

type Cache struct {
	mu          sync.RWMutex
	clk         clock.Clock
	trustDomain string

	agentsByID        map[string]agentRecord
	agentsByExpiresAt *btree.BTreeG[agentRecord]

	aliasesByEntryID  *btree.BTreeG[aliasRecord]
	aliasesBySelector *btree.BTreeG[aliasRecord]

	entriesByEntryID  map[string]*types.Entry
	entriesByParentID map[string]map[string]*types.Entry
}

func NewCache(clk clock.Clock, trustDomain string) *Cache { _ = "STUB: not implemented"; return nil }

func (c *Cache) LookupAuthorizedEntries(agentID spiffeid.ID, requestedEntries map[string]struct{}) map[string]api.ReadOnlyEntry {
	_ = "STUB: not implemented"
	return nil
}

// Load up the agent selectors. If the agent info does not exist, it is
// likely that the cache is still catching up to a recent attestation.
// Since the calling agent has already been authorized and authenticated,
// it is safe to continue with the authorized entry crawl to obtain entries
// that are directly parented against the agent. Any entries that would be
// obtained via node aliasing will not be returned until the cache is
// updated with the node selectors for the agent.

func (c *Cache) GetAuthorizedEntries(agentID spiffeid.ID) []api.ReadOnlyEntry {
	_ = "STUB: not implemented"
	return nil
}

// Load up the agent selectors. If the agent info does not exist, it is
// likely that the cache is still catching up to a recent attestation.
// Since the calling agent has already been authorized and authenticated,
// it is safe to continue with the authorized entry crawl to obtain entries
// that are directly parented against the agent. Any entries that would be
// obtained via node aliasing will not be returned until the cache is
// updated with the node selectors for the agent.

func (c *Cache) UpdateEntry(entry *types.Entry) {
	_ = "STUB: not implemented"
	// Ensure that the trust domain of the entry matches the expected trust domain.
	// This allows us to use only the path component as a key in maps.
	return
}

func (c *Cache) RemoveEntry(entryID string) { _ = "STUB: not implemented"; return }

func (c *Cache) UpdateAgent(agentID string, expiresAt time.Time, selectors []*types.Selector) {
	_ = "STUB: not implemented"
	return
}

// Need to delete existing record from the ExpiresAt index first. Use
// the ID index to locate the existing record.

func (c *Cache) RemoveAgent(agentID string) { _ = "STUB: not implemented"; return }

func (c *Cache) PruneExpiredAgents() int { _ = "STUB: not implemented"; return 0 }

func (c *Cache) appendDescendents(records []api.ReadOnlyEntry, parentID string, parentSeen stringSet) []api.ReadOnlyEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) addDescendants(foundEntries map[string]api.ReadOnlyEntry, parentID string, requestedEntries map[string]struct{}, parentSeen stringSet) {
	_ = "STUB: not implemented"
	return
}

func (c *Cache) getAgentAliases(agentSelectors selectorSet) []aliasRecord {
	_ = "STUB: not implemented"
	// Keep track of which aliases have already been evaluated.
	return nil
}

// Figure out which aliases the agent belongs to.

func (c *Cache) updateEntry(entry *types.Entry) { _ = "STUB: not implemented"; return }

func (c *Cache) removeEntry(entryID string) { _ = "STUB: not implemented"; return }

// entry was a normal workload registration. No need to search the aliases.

func (c *Cache) Stats() CacheStats { _ = "STUB: not implemented"; return *new(CacheStats) }

func isNodeAlias(e *types.Entry) bool { _ = "STUB: not implemented"; return false }

type CacheStats struct {
	AgentsByID        int
	AgentsByExpiresAt int
	AliasesByEntryID  int
	AliasesBySelector int
	EntriesByEntryID  int
}
