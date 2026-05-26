package k8sconfigmap

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk/support/bundleformat"
	bundlepublisherv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/bundlepublisher/v1"
	"github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "k8s_configmap"
)

type pluginHooks struct {
	newK8sClientFunc func(string) (kubernetesClient, error)
}

// BuiltIn returns a new BundlePublisher built-in plugin.
func BuiltIn() catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *

	// New creates a new k8s_configmap BundlePublisher plugin instance.
	new(catalog.BuiltIn)
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

// Config holds the configuration of the plugin.
type Config struct {
	Clusters map[string]*Cluster `hcl:"clusters,block" json:"clusters"`
}

// Config holds the configuration of the plugin.
type Cluster struct {
	Format         string `hcl:"format" json:"format"`
	Namespace      string `hcl:"namespace" json:"namespace"`
	ConfigMapName  string `hcl:"configmap_name" json:"configmap_name"`
	ConfigMapKey   string `hcl:"configmap_key" json:"configmap_key"`
	KubeConfigPath string `hcl:"kubeconfig_path" json:"kubeconfig_path"`
	RefreshHint    string `hcl:"refresh_hint" json:"refresh_hint"`

	// bundleFormat is used to store the content of BundleFormat, parsed
	// as bundleformat.Format.
	bundleFormat bundleformat.Format

	// k8sClient is the Kubernetes client used to interact with the cluster, set
	// when the plugin is configured.
	k8sClient kubernetesClient

	// parsedRefreshHint is used to store the content of RefreshHint, parsed
	// as an int64.
	parsedRefreshHint int64
}

// buildConfig builds the plugin configuration from the provided HCL config.
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

	hooks pluginHooks
	log   hclog.Logger
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

// PublishBundle puts the bundle in the configured Kubernetes ConfigMap.
func (p *Plugin) PublishBundle(ctx context.Context, req *bundlepublisherv1.PublishBundleRequest) (*bundlepublisherv1.PublishBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle not changed. No need to publish.

// Validate validates the configuration of the plugin.
func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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
func newPlugin(newK8sClientFunc func(string) (kubernetesClient, error)) *Plugin {
	_ = "STUB: not implemented"
	return nil
}
