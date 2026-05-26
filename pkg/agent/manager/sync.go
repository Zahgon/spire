package manager

import (
	"context"
	"crypto"
	"crypto/x509"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/agent/workloadkey"
	"github.com/spiffe/spire/proto/spire/common"
)

type csrRequest struct {
	EntryID              string
	SpiffeID             string
	CurrentSVIDExpiresAt time.Time
}

type SVIDCache interface {
	// UpdateEntries updates entries on cache
	UpdateEntries(update *cache.UpdateEntries, checkSVID func(*common.RegistrationEntry, *common.RegistrationEntry, *cache.X509SVID) bool)

	// UpdateSVIDs updates SVIDs on provided records
	UpdateSVIDs(update *cache.UpdateSVIDs)

	// GetStaleEntries gets a list of records that need update SVIDs
	GetStaleEntries() []*cache.StaleEntry

	// TaintX509SVIDs marks all SVIDs signed by a tainted X.509 authority as tainted
	// to force their rotation.
	TaintX509SVIDs(ctx context.Context, taintedX509Authorities []*x509.Certificate)

	// TaintJWTSVIDs removes JWT-SVIDs with tainted authorities from the cache,
	// forcing the server to issue a new JWT-SVID when one with a tainted
	// authority is requested.
	TaintJWTSVIDs(ctx context.Context, taintedJWTAuthorities map[string]struct{})
}

func (m *manager) syncSVIDs(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// processTaintedAuthorities verifies if a new authority is tainted and forces rotation in all caches if required.
func (m *manager) processTaintedAuthorities(ctx context.Context, bundle *spiffebundle.Bundle, x509Authorities []string, jwtAuthorities map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Taint all regular X.509 SVIDs

// Taint all SVIDStore SVIDs

// Notify rotator about new tainted authorities

// Taint JWT-SVIDs in the cache

// synchronize fetches the authorized entries from the server, updates the
// cache, and fetches missing/expiring SVIDs.
func (m *manager) synchronize(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Process all tainted authorities. The bundle is shared between both caches using regular cache data.

// Set last success sync

func (m *manager) updateCache(ctx context.Context, update *cache.UpdateEntries, log logrus.FieldLogger, cacheType string, c SVIDCache) error {
	_ = "STUB: not implemented"
	// update the cache and build a list of CSRs that need to be processed
	// in this interval.
	//
	// the values in `update` now belong to the cache. DO NOT MODIFY.
	return nil
}

// no SVID

// SVID has an empty chain. this is not expected to happen.

// Registration entry has been updated

// SVID is good

// TODO: this values are not real, we may remove

func (m *manager) updateSVIDs(ctx context.Context, log logrus.FieldLogger, c SVIDCache) error {
	_ = "STUB: not implemented"
	return nil
}

// we've exceeded the CSR limit, don't make any more CSRs

// the values in `update` now belong to the cache. DO NOT MODIFY.

func (m *manager) fetchSVIDs(ctx context.Context, csrs []csrRequest) (_ *cache.UpdateSVIDs, err error) {
	_ = "STUB: not implemented"
	// Put all the CSRs in an array to make just one call with all the CSRs.
	return nil, nil
}

// Since entryIDs are unique, this shouldn't happen. Log just in case

// Reduce csr size for next invocation

// fetchEntries fetches entries that the agent is entitled to, divided in lists, one for regular entries and
// another one for storable entries
func (m *manager) fetchEntries(ctx context.Context) (_ *cache.UpdateEntries, _ *cache.UpdateEntries, err error) {
	_ = "STUB: not implemented"
	// Put all the CSRs in an array to make just one call with all the CSRs.
	return nil, nil, nil
}

// Get all Subject Key IDs and KeyIDs of tainted authorities

func newCSR(spiffeID spiffeid.ID, keyType workloadkey.KeyType) (crypto.Signer, []byte, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil, nil
}

func parseBundles(bundles map[string]*common.Bundle) (map[spiffeid.TrustDomain]*cache.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNewItemsFromSlice(current map[string]struct{}, items []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getNewItemsFromMap(current map[string]struct{}, items map[string]struct{}) []string {
	_ = "STUB: not implemented"
	return nil
}
