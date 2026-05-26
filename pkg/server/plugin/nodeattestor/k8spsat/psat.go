package k8spsat

import (
	"context"
	"sync"

	hclog "github.com/hashicorp/go-hclog"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/plugin/k8s/apiserver"
	"github.com/spiffe/spire/pkg/common/pluginconf"

	// Add auth providers to authenticate to clusters to verify tokens
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

const (
	pluginName = "k8s_psat"
)

var (
	defaultAudience = []string{"spire-server"}
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *AttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

// AttestorConfig contains a map of clusters that uses cluster name as key
type AttestorConfig struct {
	Clusters map[string]*ClusterConfig `hcl:"clusters"`
}

// ClusterConfig holds a single cluster configuration
type ClusterConfig struct {
	// Array of allowed service accounts names
	// Attestation is denied if coming from a service account that is not in the list
	ServiceAccountAllowList []string `hcl:"service_account_allow_list"`

	// Audience for PSAT token validation
	// If audience is not configured, defaultAudience will be used
	// If audience value is set to an empty slice, k8s apiserver audience will be used
	Audience *[]string `hcl:"audience"`

	// Kubernetes configuration file path
	// Used to create a k8s client to query the API server. If string is empty, in-cluster configuration is used
	KubeConfigFile string `hcl:"kube_config_file"`

	// Node labels that are allowed to use as selectors
	AllowedNodeLabelKeys []string `hcl:"allowed_node_label_keys"`

	// Pod labels that are allowed to use as selectors
	AllowedPodLabelKeys []string `hcl:"allowed_pod_label_keys"`
}

type attestorConfig struct {
	trustDomain string
	clusters    map[string]*clusterConfig
}

type clusterConfig struct {
	serviceAccounts      map[string]bool
	audience             []string
	client               apiserver.Client
	allowedNodeLabelKeys map[string]bool
	allowedPodLabelKeys  map[string]bool
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *attestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// AttestorPlugin is a PSAT (Projected SAT) node attestor plugin
type AttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mu     sync.RWMutex
	config *attestorConfig
	log    hclog.Logger
}

// New creates a new PSAT node attestor plugin
func New() *AttestorPlugin { _ = "STUB: not implemented"; return nil }

var _ nodeattestorv1.NodeAttestorServer = (*AttestorPlugin)(nil)

// SetLogger sets up plugin logging
func (p *AttestorPlugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *AttestorPlugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *AttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *AttestorPlugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *AttestorPlugin) getConfig() (*attestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
