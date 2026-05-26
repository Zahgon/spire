package dscache

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	datastoreCacheExpiry = time.Second
)

type useCache struct{}

func WithCache(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type bundleEntry struct {
	mu     sync.Mutex
	ts     time.Time
	bundle *common.Bundle
}

type DatastoreCache struct {
	datastore.DataStore
	clock clock.Clock

	bundlesMu sync.Mutex
	bundles   map[string]*bundleEntry
}

func New(ds datastore.DataStore, clock clock.Clock) *DatastoreCache {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DatastoreCache) FetchBundle(ctx context.Context, trustDomain string) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't cache bundle "misses"

func (ds *DatastoreCache) PruneBundle(ctx context.Context, trustDomainID string, expiresBefore time.Time) (changed bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ds *DatastoreCache) AppendBundle(ctx context.Context, b *common.Bundle) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DatastoreCache) UpdateBundle(ctx context.Context, b *common.Bundle, mask *common.BundleMask) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DatastoreCache) DeleteBundle(ctx context.Context, td string, mode datastore.DeleteMode) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DatastoreCache) SetBundle(ctx context.Context, b *common.Bundle) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DatastoreCache) TaintX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToTaint string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DatastoreCache) RevokeX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToRevoke string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DatastoreCache) TaintJWTKey(ctx context.Context, trustDomainID string, authorityID string) (taintedKey *common.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DatastoreCache) RevokeJWTKey(ctx context.Context, trustDomainID string, authorityID string) (revokedKey *common.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DatastoreCache) invalidateBundleEntry(trustDomainID string) {
	_ = "STUB: not implemented"
	return
}
