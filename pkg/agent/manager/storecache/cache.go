package storecache

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/common"
)

// Record holds the latest cached SVID with its context
type Record struct {
	// ID holds entry ID
	ID string
	// Entry holds registration entry for record
	Entry *common.RegistrationEntry
	// ExpiresAt is the expiration time for SVID
	ExpiresAt time.Time
	// Svid holds a valid X509-SVID
	Svid *cache.X509SVID
	// Revision is the current cache record version
	Revision int64
	// Bundles holds trust domain bundle together with federated bundle
	Bundles map[spiffeid.TrustDomain]*spiffebundle.Bundle
	// HandledEntry holds the previous entry revision. It is useful to define
	// what changed between versions.
	HandledEntry *common.RegistrationEntry
}

// cachedRecord holds internal cached SVIDs
type cachedRecord struct {
	entry *common.RegistrationEntry
	svid  *cache.X509SVID

	revision     int64
	handled      int64
	handledEntry *common.RegistrationEntry
}

// Config is the store cache configuration
type Config struct {
	Log         logrus.FieldLogger
	TrustDomain spiffeid.TrustDomain
	Metrics     telemetry.Metrics
}

type Cache struct {
	c *Config

	mtx sync.RWMutex

	// bundles holds the latest bundles
	bundles map[spiffeid.TrustDomain]*spiffebundle.Bundle
	// records holds all the latest SVIDs with its entries
	records map[string]*cachedRecord

	// staleEntries holds stale registration entries
	staleEntries map[string]bool
}

func New(config *Config) *Cache { _ = "STUB: not implemented"; return nil }

// UpdateEntries using `UpdateEntries` updates and validates latest entries,
// record's revision number is incremented on each record based on:
// - Knowledge or when the SVID for that entry changes
// - Knowledge when the bundle changes
// - Knowledge when a federated bundle related to a storable entry changes
func (c *Cache) UpdateEntries(update *cache.UpdateEntries, checkSVID func(*common.RegistrationEntry, *common.RegistrationEntry, *cache.X509SVID) bool) {
	_ = "STUB: not implemented"
	return
}

// Remove bundles that no longer exist. The bundle for the agent trust
// domain should NOT be removed even if not present (which should only be
// the case if there is a bug on the server) since it is necessary to
// authenticate the server.

// bundle no longer exists.

// Update bundles with changes, populating a "changed" set that we can
// check when processing registration entries to know if they need to
// increment revision.

// Remove records of registration entries that no longer exist

// Record is marked as removed and already processed by store service,
// since the value of latest handled is equal to current revision

// Entry waiting to be removed on platform

// Mark the entry as removed, setting "entry" as 'nil'. The latest handled entry is set as current entry,
// and increment the revision.
// The record will be taken by the service to propagate it to SVID Stores.
// Once the SVID Store plugin removes it from the specific platform, 'revision' will be equal to 'handled'

// Add/update records for registration entries in the update

// TODO: may we separate cases to add more details about why we increment revision?

// Entry revision changed that means entry changed

// Increase the revision when the TD bundle changed

// Mark record as stale when a federated bundle changed

// Increase the revision when the federated bundle related with the entry is removed

// Related bundles or entry changed, mark this record as outdated

// TODO: in case where entry is updated may we not increment revision and just add it to stale?
// Then stale will be taken by sync and it will increment revision.

// Log when entry is updated or created.

// UpdateSVIDs updates cache with latest SVIDs
func (c *Cache) UpdateSVIDs(update *cache.UpdateSVIDs) { _ = "STUB: not implemented"; return }

// Add/update records for registration entries in the update

// Record is going to be deleted

// Increment revision since record changed

// Cache record is updated, remove it from stale map

func (c *Cache) TaintX509SVIDs(ctx context.Context, taintedX509Authorities []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}

// Skip nil or already tainted SVIDs

// Mark SVID as tainted by setting it to nil

func (c *Cache) TaintJWTSVIDs(ctx context.Context, taintedJWTAuthorities map[string]struct{}) {
	_ = "STUB: not implemented"
	// Nothing to do here
	return

	// GetStaleEntries obtains a list of stale entries, that needs new SVIDs
}

func (c *Cache) GetStaleEntries() []*cache.StaleEntry { _ = "STUB: not implemented"; return nil }

func (c *Cache) CountX509SVIDs() int { _ = "STUB: not implemented"; return 0 }

// ReadyToStore returns all records that are ready to be stored
func (c *Cache) ReadyToStore() []*Record { _ = "STUB: not implemented"; return nil }

// HandledRecord updates handled revision, and sets the latest processed entry
func (c *Cache) HandledRecord(handledEntry *common.RegistrationEntry, revision int64) {
	_ = "STUB: not implemented"
	return
}

// Records returns all the records in the cache.
// This function exists only to facilitate testing.
func (c *Cache) Records() []*Record { _ = "STUB: not implemented"; return nil }

// updateOrCreateRecord creates a new record if required or updates the existing record.
// In case that the record is updated, the old entry is returned.
func (c *Cache) updateOrCreateRecord(newEntry *common.RegistrationEntry) (*cachedRecord, *common.RegistrationEntry) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Revision will be incremented after validations

// isBundleChanged indicates whether any federated bundle changed or not
func isBundleChanged(federatesWith []string, bundleChanged map[spiffeid.TrustDomain]bool) bool {
	_ = "STUB: not implemented"
	return false
}

// There are logs on previous steps that already log this case

// In case that a single bundle changed, all the record is marked as outdated

// isBundleRemoved indicates if any federated bundle is now removed
func isBundleRemoved(federatesWith []string, bundleRemoved map[spiffeid.TrustDomain]bool) bool {
	_ = "STUB: not implemented"
	return false
}

// There are logs on previous steps that already log this case

// In case a single bundle is removed, all the record is marked as outdated

// recordFromCache parses cache record into storable Record
func recordFromCache(r *cachedRecord, bundles map[spiffeid.TrustDomain]*spiffebundle.Bundle) *Record {
	_ = "STUB: not implemented"
	return nil
}

// TODO: May we filter bundles based in TD and federated bundle?
