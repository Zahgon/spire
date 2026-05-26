package gcsbundle

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	identityproviderv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/identityprovider/v1"
	notifierv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/notifier/v1"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtIn(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type bucketClient interface {
	GetObjectGeneration(ctx context.Context, bucket, object string) (int64, error)
	PutObject(ctx context.Context, bucket, object string, data []byte, generation int64) error
	Close() error
}

type configuration struct {
	Bucket             string `hcl:"bucket"`
	ObjectPath         string `hcl:"object_path"`
	ServiceAccountFile string `hcl:"service_account_file"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *configuration {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	notifierv1.UnsafeNotifierServer
	configv1.UnsafeConfigServer

	mu               sync.RWMutex
	log              hclog.Logger
	config           *configuration
	identityProvider identityproviderv1.IdentityProviderServiceClient

	hooks struct {
		newBucketClient func(ctx context.Context, configPath string) (bucketClient, error)
	}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) BrokerHostServices(broker pluginsdk.ServiceBroker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Notify(ctx context.Context, req *notifierv1.NotifyRequest) (*notifierv1.NotifyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignore the bundle presented in the request. see updateBundleObject for details on why.

func (p *Plugin) NotifyAndAdvise(ctx context.Context, req *notifierv1.NotifyAndAdviseRequest) (*notifierv1.NotifyAndAdviseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignore the bundle presented in the request. see updateBundleObject for details on why.

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (resp *configv1.ConfigureResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (resp *configv1.ValidateResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() (*configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Plugin) updateBundleObject(ctx context.Context, c *configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Get the bundle object generation that we can use to resolve
// conflicts racing on updates from other servers.

// Load bundle data from the identity provider. The bundle has to
// be loaded after fetching the generation so we can properly detect
// and correct a race updating the bundle (i.e. read-modify-write
// semantics).

// Upload the bundle, handling version conflicts

// If there is a conflict then some other server won the race updating
// the object. We need to retrieve the latest bundle and try again.

type gcsBucketClient struct {
	client *storage.Client
}

func newGCSBucketClient(ctx context.Context, serviceAccountFile string) (bucketClient, error) {
	_ = "STUB: not implemented"
	return *new(bucketClient), nil
}

func (c *gcsBucketClient) GetObjectGeneration(ctx context.Context, bucket, object string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *gcsBucketClient) PutObject(ctx context.Context, bucket, object string, data []byte, generation int64) error {
	_ = "STUB: not implemented"
	// If for whatever reason we don't make it to w.Close(), canceling the
	// context will cleanly release resources held by the writer.
	return nil
}

func (c *gcsBucketClient) Close() error { _ = "STUB: not implemented"; return nil }

// bundleData formats the bundle data for storage in GCS
func bundleData(bundle *plugintypes.Bundle) []byte { _ = "STUB: not implemented"; return nil }

// no need to check the error since we're encoding into a memory buffer

func isConditionNotMetError(err error) bool { _ = "STUB: not implemented"; return false }
