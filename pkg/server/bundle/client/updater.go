package client

import (
	"context"
	"sync"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type BundleUpdaterConfig struct {
	TrustDomain spiffeid.TrustDomain
	DataStore   datastore.DataStore

	TrustDomainConfig TrustDomainConfig

	// newClientHook is a test hook for injecting client behavior
	newClientHook func(ClientConfig) (Client, error)
}

type BundleUpdater interface {
	// UpdateBundle fetches the local bundle from the datastore and the
	// endpoint bundle from the endpoint. The function will return an error if
	// the local bundle cannot be fetched, the endpoint bundle cannot be
	// downloaded, or there is a problem persisting the bundle. The local
	// bundle will always be returned if it was fetched, independent of any
	// other failures performing the update. The endpoint bundle is ONLY
	// returned if it can be successfully downloaded, is different from the
	// local bundle, and is successfully stored.
	UpdateBundle(ctx context.Context) (*spiffebundle.Bundle, *spiffebundle.Bundle, error)

	// GetTrustDomainConfig returns the configuration for the updater
	GetTrustDomainConfig() TrustDomainConfig

	// SetTrustDomainConfig sets the configuration for the updater
	SetTrustDomainConfig(TrustDomainConfig) bool
}

type bundleUpdater struct {
	td            spiffeid.TrustDomain
	ds            datastore.DataStore
	newClientHook func(ClientConfig) (Client, error)

	trustDomainConfigMtx sync.Mutex
	trustDomainConfig    TrustDomainConfig
}

func NewBundleUpdater(config BundleUpdaterConfig) BundleUpdater {
	_ = "STUB: not implemented"
	return *new(BundleUpdater)
}

func (u *bundleUpdater) UpdateBundle(ctx context.Context) (*spiffebundle.Bundle, *spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (u *bundleUpdater) GetTrustDomainConfig() TrustDomainConfig {
	_ = "STUB: not implemented"
	return *new(TrustDomainConfig)
}

func (u *bundleUpdater) SetTrustDomainConfig(trustDomainConfig TrustDomainConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (u *bundleUpdater) newClient(ctx context.Context, trustDomainConfig TrustDomainConfig) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func fetchBundleIfExists(ctx context.Context, ds datastore.DataStore, trustDomain spiffeid.TrustDomain) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	// Load the current bundle and extract the root CA certificates
	return nil, nil
}
