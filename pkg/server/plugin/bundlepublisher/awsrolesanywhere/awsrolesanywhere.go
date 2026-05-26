package awsrolesanywhere

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/go-hclog"
	bundlepublisherv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/bundlepublisher/v1"
	"github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "aws_rolesanywhere_trustanchor"
)

type pluginHooks struct {
	newRolesAnywhereClientFunc func(c aws.Config) (rolesAnywhere, error)
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func New() *Plugin { _ = "STUB: not implemented"; return nil }

// Config holds the configuration of the plugin.
type Config struct {
	AccessKeyID     string `hcl:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key" json:"secret_access_key"`
	Region          string `hcl:"region" json:"region"`
	TrustAnchorID   string `hcl:"trust_anchor_id" json:"trust_anchor_id"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

// Plugin is the main representation of this bundle publisher plugin.
type Plugin struct {
	bundlepublisherv1.UnsafeBundlePublisherServer
	configv1.UnsafeConfigServer

	config    *Config
	configMtx sync.RWMutex

	bundle    *types.Bundle
	bundleMtx sync.RWMutex

	hooks               pluginHooks
	rolesAnywhereClient rolesAnywhere
	log                 hclog.Logger
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

// PublishBundle puts the bundle in the Roles Anywhere trust anchor, with
// the configured id.
func (p *Plugin) PublishBundle(ctx context.Context, req *bundlepublisherv1.PublishBundleRequest) (*bundlepublisherv1.PublishBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle not changed. No need to publish.

// To prevent flooding of the logs in the case that the bundle is
// too large.

// Update the trust anchor that was found

// getBundle gets the latest bundle that the plugin has.
func (p *Plugin) getBundle() *types.Bundle { _ = "STUB: not implemented"; return nil }

// getConfig gets the configuration of the plugin.
func (p *Plugin) getConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// setBundle updates the current bundle in the plugin with the provided bundle.
func (p *Plugin) setBundle(bundle *types.Bundle) { _ = "STUB: not implemented"; return }

// builtin creates a new BundlePublisher built-in plugin.
func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

// newPlugin returns a new plugin instance.
func newPlugin(newRolesAnywhereClientFunc func(c aws.Config) (rolesAnywhere, error)) *Plugin {
	_ = "STUB: not implemented"
	return nil
}
