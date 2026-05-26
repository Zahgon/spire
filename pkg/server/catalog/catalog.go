package catalog

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/datastore"
	ds_sql "github.com/spiffe/spire/pkg/server/datastore/sqlstore"
	"github.com/spiffe/spire/pkg/server/hostservice/agentstore"
	"github.com/spiffe/spire/pkg/server/hostservice/identityprovider"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher"
	"github.com/spiffe/spire/pkg/server/plugin/credentialcomposer"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/server/plugin/notifier"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority"
)

const (
	bundlePublisherType    = "BundlePublisher"
	credentialComposerType = "CredentialComposer" //nolint: gosec // this is not a hardcoded credential...
	dataStoreType          = "DataStore"
	keyManagerType         = "KeyManager"
	nodeAttestorType       = "NodeAttestor"
	notifierType           = "Notifier"
	upstreamAuthorityType  = "UpstreamAuthority"
)

var ReconfigureTask = catalog.ReconfigureTask

type Catalog interface {
	GetBundlePublishers() []bundlepublisher.BundlePublisher
	GetCredentialComposers() []credentialcomposer.CredentialComposer
	GetDataStore() datastore.DataStore
	GetNodeAttestorNamed(name string) (nodeattestor.NodeAttestor, bool)
	GetKeyManager() keymanager.KeyManager
	GetNotifiers() []notifier.Notifier
	GetUpstreamAuthority() (upstreamauthority.UpstreamAuthority, bool)
}

type PluginConfigs = catalog.PluginConfigs

type Config struct {
	Log           logrus.FieldLogger
	TrustDomain   spiffeid.TrustDomain
	PluginConfigs PluginConfigs

	Metrics          telemetry.Metrics
	IdentityProvider *identityprovider.IdentityProvider
	AgentStore       *agentstore.AgentStore
	HealthChecker    health.Checker
}

type datastoreRepository struct{ datastore.Repository }

type Repository struct {
	bundlePublisherRepository
	credentialComposerRepository
	datastoreRepository
	keyManagerRepository
	nodeAttestorRepository
	notifierRepository
	upstreamAuthorityRepository

	log      logrus.FieldLogger
	dsCloser io.Closer
	catalog  *catalog.Catalog
}

type dsConfigurer struct {
	ds *ds_sql.Plugin
}

func (c *dsConfigurer) Configure(ctx context.Context, _ catalog.CoreConfig, configuration string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dsConfigurer) Validate(ctx context.Context, coreConfig catalog.CoreConfig, configuration string) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *Repository) Plugins() map[string]catalog.PluginRepo {
	_ = "STUB: not implemented"
	return nil
}

func (repo *Repository) Services() []catalog.ServiceRepo { _ = "STUB: not implemented"; return nil }

func (repo *Repository) Reconfigure(ctx context.Context) { _ = "STUB: not implemented"; return }

func (repo *Repository) Close() {
	_ = "STUB: not implemented"
	// Must close in reverse initialization order!
	return
}

func Load(ctx context.Context, config Config) (_ *Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Strip out the Datastore plugin configuration and load the SQL plugin
// directly. This allows us to bypass gRPC and get rid of response limits.

func ValidateConfig(ctx context.Context, config Config) (pluginNotes map[string][]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadSQLDataStore(ctx context.Context, config Config, coreConfig catalog.CoreConfig, datastoreConfigs catalog.PluginConfigs) (*ds_sql.Plugin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
