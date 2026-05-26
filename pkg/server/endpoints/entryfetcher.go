package endpoints

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/cache/entrycache"
	"github.com/spiffe/spire/pkg/server/datastore"
)

var _ api.AuthorizedEntryFetcher = (*AuthorizedEntryFetcherWithFullCache)(nil)

type entryCacheBuilderFn func(ctx context.Context) (entrycache.Cache, error)

type AuthorizedEntryFetcherWithFullCache struct {
	buildCache           entryCacheBuilderFn
	cache                entrycache.Cache
	clk                  clock.Clock
	log                  logrus.FieldLogger
	ds                   datastore.DataStore
	mu                   sync.RWMutex
	cacheReloadInterval  time.Duration
	pruneEventsOlderThan time.Duration
}

func NewAuthorizedEntryFetcherWithFullCache(ctx context.Context, buildCache entryCacheBuilderFn, log logrus.FieldLogger, clk clock.Clock, ds datastore.DataStore, cacheReloadInterval, pruneEventsOlderThan time.Duration) (*AuthorizedEntryFetcherWithFullCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AuthorizedEntryFetcherWithFullCache) LookupAuthorizedEntries(ctx context.Context, agentID spiffeid.ID, entryIDs map[string]struct{}) (map[string]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AuthorizedEntryFetcherWithFullCache) FetchAuthorizedEntries(_ context.Context, agentID spiffeid.ID) ([]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunRebuildCacheTask starts a ticker which rebuilds the in-memory entry cache.
func (a *AuthorizedEntryFetcherWithFullCache) RunRebuildCacheTask(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// PruneEventsTask start a ticker which prunes old events
func (a *AuthorizedEntryFetcherWithFullCache) PruneEventsTask(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizedEntryFetcherWithFullCache) pruneEvents(ctx context.Context, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
