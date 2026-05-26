package endpoints

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"

	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/authorizedentries"
	"github.com/spiffe/spire/pkg/server/cache/nodecache"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type attestedNodes struct {
	cache     *authorizedentries.Cache
	nodeCache *nodecache.Cache
	clk       clock.Clock
	ds        datastore.DataStore
	log       logrus.FieldLogger
	metrics   telemetry.Metrics

	eventsBeforeFirst map[uint]struct{}

	firstEvent     uint
	firstEventTime time.Time
	lastEvent      uint

	eventTracker *eventTracker
	eventTimeout time.Duration

	fetchNodes map[string]struct{}

	// metrics change detection
	skippedNodeEvents int
	lastCacheStats    authorizedentries.CacheStats
}

func (a *attestedNodes) captureChangedNodes(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *attestedNodes) searchBeforeFirstEvent(ctx context.Context) error {
	_ = "STUB: not implemented"
	// First event detected, and startup was less than a transaction timout away.
	return nil
}

// if we have seen it before, don't reload it.

// zero out unused event tracker

func (a *attestedNodes) selectPolledEvents(ctx context.Context) {
	_ = "STUB: not implemented"
	// check if the polled events have appeared out-of-order
	return
}

func (a *attestedNodes) scanForNewEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// event time determines if we have seen the first event.

// track any skipped event ids, should they appear later.

// every event adds its entry to the entry fetch list.

func (a *attestedNodes) loadCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	// TODO: determine if this needs paging
	return nil
}

// buildAttestedNodesCache fetches all attested nodes and adds the unexpired ones to the cache.
// It runs once at startup.
func buildAttestedNodesCache(ctx context.Context, log logrus.FieldLogger, metrics telemetry.Metrics, ds datastore.DataStore, clk clock.Clock, cache *authorizedentries.Cache, nodeCache *nodecache.Cache, cacheReloadInterval, eventTimeout time.Duration) (*attestedNodes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize gauges to nonsense values to force a change.

// updateCache Fetches all the events since the last time this function was running and updates
// the cache with all the changes.
func (a *attestedNodes) updateCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *attestedNodes) updateCachedNodes(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Node was deleted

func (a *attestedNodes) emitMetrics() { _ = "STUB: not implemented"; return }

// AgentsByID and AgentsByExpiresAt should be the same.
