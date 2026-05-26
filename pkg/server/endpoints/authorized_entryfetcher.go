package endpoints

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/authorizedentries"
	"github.com/spiffe/spire/pkg/server/cache/nodecache"
	"github.com/spiffe/spire/pkg/server/datastore"
)

var _ api.AuthorizedEntryFetcher = (*AuthorizedEntryFetcherEvents)(nil)

const pageSize = 10000

type AuthorizedEntryFetcherEventsConfig struct {
	clk                     clock.Clock
	log                     logrus.FieldLogger
	cacheReloadInterval     time.Duration
	fullCacheReloadInterval time.Duration
	pruneEventsOlderThan    time.Duration
	eventTimeout            time.Duration
	ds                      datastore.DataStore
	nodeCache               *nodecache.Cache
	metrics                 telemetry.Metrics
}

type AuthorizedEntryFetcherEvents struct {
	c                   AuthorizedEntryFetcherEventsConfig
	cache               *authorizedentries.Cache
	registrationEntries eventsBasedCache
	attestedNodes       eventsBasedCache
	mu                  sync.RWMutex
	trustDomain         string
}

type eventsBasedCache interface {
	updateCache(ctx context.Context) error
}

func NewAuthorizedEntryFetcherEvents(ctx context.Context, trustDomain string, c AuthorizedEntryFetcherEventsConfig) (*AuthorizedEntryFetcherEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AuthorizedEntryFetcherEvents) LookupAuthorizedEntries(ctx context.Context, agentID spiffeid.ID, entryIDs map[string]struct{}) (map[string]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AuthorizedEntryFetcherEvents) FetchAuthorizedEntries(_ context.Context, agentID spiffeid.ID) ([]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunUpdateCacheTask starts a ticker which rebuilds the in-memory entry cache.
func (a *AuthorizedEntryFetcherEvents) RunUpdateCacheTask(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// PruneEventsTask start a ticker which prunes old events
func (a *AuthorizedEntryFetcherEvents) PruneEventsTask(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizedEntryFetcherEvents) pruneEvents(ctx context.Context, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizedEntryFetcherEvents) updateCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizedEntryFetcherEvents) buildCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizedEntryFetcherEvents) startTickers() (*clock.Ticker, *clock.Ticker) {
	_ = "STUB: not implemented"
	return nil, nil
}
