// Package pubmanager manages the publishing of the trust bundle to external
// stores through the configured BundlePublisher plugins.
package pubmanager

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	// refreshInterval is the interval to check for an updated trust bundle.
	refreshInterval = 30 * time.Second
)

// NewManager creates a new bundle publishing manager.
func NewManager(c *ManagerConfig) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil,

		// ManagerConfig is the config for the bundle publishing manager.
		nil
}

type ManagerConfig struct {
	BundlePublishers []bundlepublisher.BundlePublisher
	DataStore        datastore.DataStore
	Clock            clock.Clock
	Log              logrus.FieldLogger
	TrustDomain      spiffeid.TrustDomain
}

// Manager is the manager for bundle publishing. It implements the PubManager
// interface.
type Manager struct {
	bundleUpdatedCh  chan struct{}
	bundlePublishers []bundlepublisher.BundlePublisher
	clock            clock.Clock
	dataStore        datastore.DataStore
	log              logrus.FieldLogger
	trustDomain      spiffeid.TrustDomain

	hooks struct {
		// Test hook used to indicate an attempt to publish a bundle using a
		// specific bundle publisher.
		publishResultCh chan *publishResult

		// Test hook used to indicate when the action of publishing a bundle
		// has finished.
		publishedCh chan error
	}
}

// Run runs the bundle publishing manager.
func (m *Manager) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// BundleUpdated tells the bundle publishing manager that the bundle has been
// updated and forces a PublishBundle operation on all the plugins.
func (m *Manager) BundleUpdated() { _ = "STUB: not implemented"; return }

// callPublishBundle calls the publishBundle function and logs if there was an
// error.
func (m *Manager) callPublishBundle(ctx context.Context) { _ = "STUB: not implemented"; return }

// publishBundle iterates through the configured bundle publishers and calls
// PublishBundle with the fetched bundle. This function only returns an error
// if bundle publishers can't be called due to a failure fetching the bundle
// from the datastore.
func (m *Manager) publishBundle(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PublishBundle was called on all the plugins. Is the responsibility of
// each plugin to handle failure conditions and implement a retry logic if
// needed.

// triggerPublishResultHook is called to know when the publish action using a
// specific bundle publisher has happened. It informs the result of calling the
// PublishBundle method to a bundle publisher.
func (m *Manager) triggerPublishResultHook(result *publishResult) {
	_ = "STUB: not implemented"
	return
}

// publishDone is called to know when a publish action has finished and informs
// if there was an error in the overall action (not specific to a bundle
// publisher). A publish action happens periodically (every refreshInterval) and
// also when BundleUpdated() is called.
func (m *Manager) publishDone(err error) { _ = "STUB: not implemented"; return }

// publishResult holds information about the result of trying to publish a
// bundle using a specific bundle publisher.
type publishResult struct {
	pluginName string
	bundle     *common.Bundle
	err        error
}

func (m *Manager) drainBundleUpdated() { _ = "STUB: not implemented"; return }

func newManager(c *ManagerConfig) (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }
