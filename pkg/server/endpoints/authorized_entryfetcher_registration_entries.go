package endpoints

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/authorizedentries"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type registrationEntries struct {
	cache   *authorizedentries.Cache
	clk     clock.Clock
	ds      datastore.DataStore
	log     logrus.FieldLogger
	metrics telemetry.Metrics

	eventsBeforeFirst map[uint]struct{}

	firstEvent     uint
	firstEventTime time.Time
	lastEvent      uint

	eventTracker *eventTracker
	eventTimeout time.Duration
	pageSize     int32

	fetchEntries map[string]struct{}

	// metrics change detection
	skippedEntryEvents int
	lastCacheStats     authorizedentries.CacheStats
}

func (a *registrationEntries) captureChangedEntries(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *registrationEntries) searchBeforeFirstEvent(ctx context.Context) error {
	_ = "STUB: not implemented"
	// First event detected, and startup was less than a transaction timout away.
	return nil
}

// if we have seen it before, don't reload it.

// zero out unused event tracker

func (a *registrationEntries) selectPolledEvents(ctx context.Context) {
	_ = "STUB: not implemented"
	// check if the polled events have appeared out-of-order
	return
}

func (a *registrationEntries) scanForNewEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// event time determines if we have seen the first event.

// track any skipped event ids, should they appear later.

// every event adds its entry to the entry fetch list.

func (a *registrationEntries) loadCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Build the cache
	return nil
}

// preliminary loading should not be done via read-replicas

// buildRegistrationEntriesCache Fetches all registration entries and adds them to the cache
func buildRegistrationEntriesCache(ctx context.Context, log logrus.FieldLogger, metrics telemetry.Metrics, ds datastore.DataStore, clk clock.Clock, cache *authorizedentries.Cache, pageSize int32, cacheReloadInterval, eventTimeout time.Duration) (*registrationEntries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// updateCache Fetches all the events since the last time this function was running and updates
// the cache with all the changes.
func (a *registrationEntries) updateCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// updateCacheEntry update/deletes/creates an individual registration entry in the cache.
func (a *registrationEntries) updateCachedEntries(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// fetchEntriesPage gets the range for the page starting at pageStart
func (a *registrationEntries) fetchEntriesPage(entryIds []string, pageStart int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (a *registrationEntries) emitMetrics() { _ = "STUB: not implemented"; return }
