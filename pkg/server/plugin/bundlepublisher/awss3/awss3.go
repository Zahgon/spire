package awss3

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk/support/bundleformat"
	bundlepublisherv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/bundlepublisher/v1"
	"github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "aws_s3"
)

type pluginHooks struct {
	newS3ClientFunc func(c aws.Config) (simpleStorageService, error)
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func New() *Plugin { _ = "STUB: not implemented"; return nil }

// Config holds the configuration of the plugin.
type Config struct {
	AccessKeyID     string `hcl:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key" json:"secret_access_key"`
	Region          string `hcl:"region" json:"region"`
	Bucket          string `hcl:"bucket" json:"bucket"`
	ObjectKey       string `hcl:"object_key" json:"object_key"`
	Format          string `hcl:"format" json:"format"`
	Endpoint        string `hcl:"endpoint" json:"endpoint"`
	RefreshHint     string `hcl:"refresh_hint" json:"refresh_hint"`

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

// This plugin only supports some bundleformats.

// Plugin is the main representation of this bundle publisher plugin.
type Plugin struct {
	bundlepublisherv1.UnsafeBundlePublisherServer
	configv1.UnsafeConfigServer

	config    *Config
	configMtx sync.RWMutex

	bundle    *types.Bundle
	bundleMtx sync.RWMutex

	hooks    pluginHooks
	s3Client simpleStorageService
	log      hclog.Logger
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

// seems wrong to change plugin s3Client before config change

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublishBundle puts the bundle in the configured S3 bucket name and
// object key.
func (p *Plugin) PublishBundle(ctx context.Context, req *bundlepublisherv1.PublishBundleRequest) (*bundlepublisherv1.PublishBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle not changed. No need to publish.

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
func newPlugin(newS3ClientFunc func(c aws.Config) (simpleStorageService, error)) *Plugin {
	_ = "STUB: not implemented"
	return nil
}
