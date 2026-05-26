package manager

import (
	"crypto/x509"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/catalog"
	managerCache "github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/agent/manager/storecache"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/agent/storage"
	"github.com/spiffe/spire/pkg/agent/trustbundlesources"
	"github.com/spiffe/spire/pkg/agent/workloadkey"
	"github.com/spiffe/spire/pkg/common/rotationutil"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
)

// Config holds a cache manager configuration
type Config struct {
	// Agent SVID and key resulting from successful attestation.
	SVID                     []*x509.Certificate
	SVIDKey                  keymanager.Key
	Bundle                   *managerCache.Bundle
	Reattestable             bool
	Catalog                  catalog.Catalog
	TrustDomain              spiffeid.TrustDomain
	Log                      logrus.FieldLogger
	Metrics                  telemetry.Metrics
	ServerAddr               string
	Storage                  storage.Storage
	TrustBundleSources       trustbundlesources.Bundle
	RebootstrapMode          string
	RebootstrapDelay         time.Duration
	WorkloadKeyType          workloadkey.KeyType
	SyncInterval             time.Duration
	UseSyncAuthorizedEntries bool
	RotationInterval         time.Duration
	SVIDStoreCache           *storecache.Cache
	X509SVIDCacheMaxSize     int
	JWTSVIDCacheMaxSize      int
	DisableLRUCache          bool
	NodeAttestor             nodeattestor.NodeAttestor
	RotationStrategy         *rotationutil.RotationStrategy
	TLSPolicy                tlspolicy.Policy

	// Clk is the clock the manager will use to get time
	Clk clock.Clock
}

// New creates a cache manager based on c's configuration
func New(c *Config) Manager { _ = "STUB: not implemented"; return *new(Manager) }

func newManager(c *Config) *manager { _ = "STUB: not implemented"; return nil }
