package client

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/datastore"
)

const (
	// attemptsPerRefreshHint is the number of attempts within the returned
	// refresh hint period that the manager will attempt to refresh the
	// bundle. It is important to try more than once within a refresh hint
	// period so we can be resilient to temporary downtime or failures.
	attemptsPerRefreshHint = 4

	// configRefreshInterval is how often the manager reloads trust domain
	// configs from the source and reconciles it against the current bundle
	// updaters.
	configRefreshInterval = time.Second * 10

	// defaultRefreshInterval is how often the manager reloads the trust bundle
	// for a trust domain if that trust domain does not specify a refresh hint in
	// its current trust bundle.
	defaultRefreshInterval = time.Minute * 5
)

type TrustDomainConfig struct {
	// EndpointURL is the URL used to fetch the bundle of the federated
	// trust domain. Is served by a SPIFFE bundle endpoint server.
	EndpointURL string

	// EndpointProfile is the bundle endpoint profile used by the
	// SPIFFE bundle endpoint server.
	EndpointProfile EndpointProfileInfo
}

type EndpointProfileInfo interface {
	// The name of the endpoint profile (e.g. "https_spiffe").
	Name() string
}

type HTTPSWebProfile struct{}

func (p HTTPSWebProfile) Name() string { _ = "STUB: not implemented"; return "" }

type HTTPSSPIFFEProfile struct {
	// EndpointSPIFFEID is the expected SPIFFE ID of the bundle endpoint server.
	EndpointSPIFFEID spiffeid.ID
}

func (p HTTPSSPIFFEProfile) Name() string { _ = "STUB: not implemented"; return "" }

type ManagerConfig struct {
	Log       logrus.FieldLogger
	Metrics   telemetry.Metrics
	DataStore datastore.DataStore
	Clock     clock.Clock
	Source    TrustDomainConfigSource

	// newBundleUpdater is a test hook to inject updater behavior
	newBundleUpdater func(BundleUpdaterConfig) BundleUpdater

	// configRefreshedCh is a test hook to learn when the trust domain config
	// has been refreshed and be apprised of the next scheduled refresh.
	configRefreshedCh chan time.Duration

	// bundleRefreshedCh is a test hook to learn when a bundle has been
	// refreshed and be apprised of the next scheduled refresh.
	bundleRefreshedCh chan time.Duration
}

type Manager struct {
	log              logrus.FieldLogger
	metrics          telemetry.Metrics
	clock            clock.Clock
	ds               datastore.DataStore
	source           TrustDomainConfigSource
	configRefreshCh  chan struct{}
	configRefreshMtx sync.Mutex
	updatersMtx      sync.RWMutex
	updaters         map[spiffeid.TrustDomain]*managedBundleUpdater

	// test hooks
	newBundleUpdater  func(BundleUpdaterConfig) BundleUpdater
	configRefreshedCh chan time.Duration
	bundleRefreshedCh chan time.Duration
}

type managedBundleUpdater struct {
	BundleUpdater

	wg     sync.WaitGroup
	cancel context.CancelFunc
	runCh  chan chan error
}

func (m *managedBundleUpdater) Stop() { _ = "STUB: not implemented"; return }

func NewManager(config ManagerConfig) *Manager { _ = "STUB: not implemented"; return nil }

func (m *Manager) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Initialize the timer that will reload the configs. The initial duration
	// isn't very important since we'll reset it after the reload has
	// completed.
	return nil
}

// TriggerConfigReload triggers the manager to reload the configuration
func (m *Manager) TriggerConfigReload() { _ = "STUB: not implemented"; return }

// RefreshBundleFor refreshes the trust domain bundle for the given trust
// domain. If the trust domain is not managed by the manager, false is returned.
func (m *Manager) RefreshBundleFor(ctx context.Context, td spiffeid.TrustDomain) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Manager) refreshConfigs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Duplicate the configs map since we're going to mutate it while figuring
// out what needs to be started/updated/stopped.

// Updater still needed. Update the configuration and remove it
// from the configs list since so a new updater isn't started for
// this trust domain.

// Updater no longer needed. Stage it to be stopped and remove it
// from the updaters list.

// The remaining configs are for newly managed trust domains. Create and
// start up an updater for it.

func (m *Manager) runUpdater(ctx context.Context, trustDomain spiffeid.TrustDomain, updater BundleUpdater) {
	_ = "STUB: not implemented"
	// Initialize the timer. The initial duration does not matter since it will
	// be reset with the actual refresh interval before first use.
	return
}

// Notify the test hook

func (m *Manager) runUpdateOnce(ctx context.Context, log *logrus.Entry, trustDomain spiffeid.TrustDomain, updater BundleUpdater) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// We have no bundle to use to calculate the refresh hint. Since
// the endpoint cannot be reached without the local bundle (until
// we implement web auth), we can retry more aggressively. This
// refresh period determines how fast we'll respond to the local
// bundle being bootstrapped.
// TODO: reevaluate once we support web auth

func (m *Manager) notifyConfigRefreshed(ctx context.Context, nextRefresh time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) notifyBundleRefreshed(ctx context.Context, nextRefresh time.Duration) {
	_ = "STUB: not implemented"
	return
}

func calculateNextUpdate(b *spiffebundle.Bundle) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func cloneTrustDomainConfigs(configs map[spiffeid.TrustDomain]TrustDomainConfig) map[spiffeid.TrustDomain]TrustDomainConfig {
	_ = "STUB: not implemented"
	return nil
}
