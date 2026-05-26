package datastore

import (
	"context"
	"time"

	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

// WithBundleUpdateCallback wraps a datastore interface and provides updates to
// bundle publishers in operations that modify the local bundle.
func WithBundleUpdateCallback(ds datastore.DataStore, bundleUpdated func()) datastore.DataStore {
	_ = "STUB: not implemented"
	return *new(datastore.DataStore)
}

type datastoreWrapper struct {
	datastore.DataStore
	bundleUpdated func()
}

func (w datastoreWrapper) AppendBundle(ctx context.Context, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w datastoreWrapper) PruneBundle(ctx context.Context, trustDomainID string, expiresBefore time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w datastoreWrapper) RevokeX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToRevoke string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w datastoreWrapper) RevokeJWTKey(ctx context.Context, trustDomainID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
