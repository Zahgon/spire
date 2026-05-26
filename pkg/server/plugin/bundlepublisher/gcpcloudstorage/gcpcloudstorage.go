package gcpcloudstorage

import (
	"context"
	"io"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk/support/bundleformat"
	bundlepublisherv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/bundlepublisher/v1"
	"github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"google.golang.org/api/option"
)

const (
	pluginName = "gcp_cloudstorage"
)

type pluginHooks struct {
	newGCSClientFunc     func(ctx context.Context, opts ...option.ClientOption) (gcsService, error)
	newStorageWriterFunc func(ctx context.Context, o *storage.ObjectHandle) io.WriteCloser
	wroteObjectFunc      func() // Test hook called when an object was written.
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func New() *Plugin { _ = "STUB: not implemented"; return nil }

// Config holds the configuration of the plugin.
type Config struct {
	BucketName         string `hcl:"bucket_name" json:"bucket_name"`
	ObjectName         string `hcl:"object_name" json:"object_name"`
	Format             string `hcl:"format" json:"format"`
	ServiceAccountFile string `hcl:"service_account_file" json:"service_account_file"`
	RefreshHint        string `hcl:"refresh_hint" json:"refresh_hint"`

	// bundleFormat is used to store the content of Format, parsed
	// as bundleformat.Format.
	bundleFormat bundleformat.Format

	// parsedRefreshHint is used to store the content of RefreshHint, parsed
	// as an int64.
	parsedRefreshHint int64
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

// Only some bundleformats are supported by this plugin.

// Plugin is the main representation of this bundle publisher plugin.
type Plugin struct {
	bundlepublisherv1.UnsafeBundlePublisherServer
	configv1.UnsafeConfigServer

	config    *Config
	configMtx sync.RWMutex

	bundle    *types.Bundle
	bundleMtx sync.RWMutex

	hooks     pluginHooks
	gcsClient gcsService
	log       hclog.Logger
}

// SetLogger sets a logger in the plugin.
func (p *Plugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Configure configures the plugin.
	return
}

func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublishBundle puts the bundle in the configured GCS bucket and object name.
func (p *Plugin) PublishBundle(ctx context.Context, req *bundlepublisherv1.PublishBundleRequest) (*bundlepublisherv1.PublishBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle not changed. No need to publish.

// Purely defensive, the Bucket function implemented in GCS always returns a BucketHandle.

// Purely defensive, the Object function implemented in GCS always returns an ObjectHandle.

// Purely defensive, the NewWriter function implemented in GCS always returns a storage writer

// The number of bytes written can be safely ignored. To determine if an
// object was successfully uploaded, we need to look at the error returned
// from storageWriter.Close().

// Close the storage writer before returning.

// Close is called when the plugin is unloaded. Closes the client.
func (p *Plugin) Close() error { _ = "STUB: not implemented"; return nil }

// getBundle gets the latest bundle that the plugin has.
func (p *Plugin) getBundle() *types.Bundle { _ = "STUB: not implemented"; return nil }

// getConfig gets the configuration of the plugin.
func (p *Plugin) getConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// setBundle updates the current bundle in the plugin with the provided bundle.
func (p *Plugin) setBundle(bundle *types.Bundle) { _ = "STUB: not implemented"; return }

// setConfig sets the configuration for the plugin.
func (p *Plugin) setConfig(config *Config) { _ = "STUB: not implemented"; return }

// builtin creates a new BundlePublisher built-in plugin.
func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

// newPlugin returns a new plugin instance.
func newPlugin(newGCSClientFunc func(ctx context.Context, opts ...option.ClientOption) (gcsService, error),
	newStorageWriterFunc func(ctx context.Context, o *storage.ObjectHandle) io.WriteCloser) *Plugin {
	_ = "STUB: not implemented"
	return nil
}
