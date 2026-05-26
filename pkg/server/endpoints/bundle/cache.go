package bundle

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/proto/spire/common"

	"github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	"github.com/spiffe/spire/pkg/server/datastore"
)

const (
	cacheExpiry = time.Second
)

type Cache struct {
	ds         datastore.DataStore
	bundlesMtx sync.Mutex
	bundles    map[spiffeid.TrustDomain]*bundleEntry
	clock      clock.Clock
}

func NewCache(ds datastore.DataStore, clk clock.Clock) *Cache {
	_ = "STUB: not implemented"
	return nil
}

type bundleEntry struct {
	mu         sync.Mutex
	ts         time.Time
	bundle     *common.Bundle
	x509Bundle *x509bundle.Bundle
}

func (c *Cache) FetchBundleX509(ctx context.Context, td spiffeid.TrustDomain) (*x509bundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache) deleteEntry(td spiffeid.TrustDomain) { _ = "STUB: not implemented"; return }

// parseBundle parses a *x509bundle.Bundle from a *common.bundle.
func parseBundle(td spiffeid.TrustDomain, commonBundle *common.Bundle) (*x509bundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
