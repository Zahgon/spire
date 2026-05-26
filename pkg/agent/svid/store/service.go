package store

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/manager/storecache"
	"github.com/spiffe/spire/pkg/agent/plugin/svidstore"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	defaultInterval = 5 * time.Second
)

type Cache interface {
	// ReadyToStore is a list of store cache records that are ready to be stored on specific SVID Store
	ReadyToStore() []*storecache.Record
	// HandledRecord sets a revision to record on cache
	HandledRecord(entry *common.RegistrationEntry, revision int64)
}

type Config struct {
	Clk         clock.Clock
	Log         logrus.FieldLogger
	TrustDomain spiffeid.TrustDomain
	Cache       Cache
	Catalog     catalog.Catalog
	Metrics     telemetry.Metrics
}

type SVIDStoreService struct {
	clk clock.Clock
	log logrus.FieldLogger
	// trustDomain is the trust domain of the agent
	trustDomain spiffeid.TrustDomain
	// cache is the store cache
	cache   Cache
	cat     catalog.Catalog
	metrics telemetry.Metrics

	hooks struct {
		// test hook used to verify if a cycle finished
		storeFinished chan struct{}
	}
}

func New(c *Config) *SVIDStoreService { _ = "STUB: not implemented"; return nil }

// SetStoreFinishedHook used for testing only
func (s *SVIDStoreService) SetStoreFinishedHook(storeFinished chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// Run starts SVID Store service
func (s *SVIDStoreService) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// deleteSVID deletes a stored SVID that uses the SVIDStore plugin. It gets the plugin name from entry selectors
func (s *SVIDStoreService) deleteSVID(ctx context.Context, log logrus.FieldLogger, entry *common.RegistrationEntry) bool {
	_ = "STUB: not implemented"
	return false
}

// storeSVID creates or updates an SVID using SVIDStore plugin. It get the plugin name from entry selectors
func (s *SVIDStoreService) storeSVID(ctx context.Context, log logrus.FieldLogger, record *storecache.Record) {
	_ = "STUB: not implemented"
	return

	// Svid is not yet provided.
}

// Set revision, since SVID was updated successfully

// TODO: may we change log.Error for debug?
func (s *SVIDStoreService) processRecords(ctx context.Context) { _ = "STUB: not implemented"; return }

// Check if entry is marked to be deleted

// TODO: add a retry backoff

// Deleted successfully. update revision

// Entries with changes on selectors must be removed before SVID is stored.

// Verify if selector changed. If it changed, delete the SVID from store before updating

// TODO: add retry, and maybe fail update until it is deleted?

// requestFromRecord parses a cache record to a *svidstore.X509SVID
func (s *SVIDStoreService) requestFromRecord(record *storecache.Record, metadata []string) (*svidstore.X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is purely defensive since federatedID should be valid

// Do not add the agent's trust domain to the federated bundles

// Federated bundle not found, no action taken

// getStoreNameWithMetadata gets SVIDStore plugin name from entry selectors and selectors metadata, it fails in case an entry
func getStoreNameWithMetadata(selectors []*common.Selector) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
