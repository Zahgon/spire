package k8spsat

import (
	"context"
	"sync"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName       = "k8s_psat"
	defaultTokenPath = "/var/run/secrets/tokens/spire-agent" //nolint: gosec // false positive
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *AttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

// New creates a new PSAT attestor plugin
func New() *AttestorPlugin { _ = "STUB: not implemented"; return nil }

// AttestorPlugin is a PSAT (projected SAT) attestor plugin
type AttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mu     sync.RWMutex
	config *attestorConfig
}

// AttestorConfig holds configuration for AttestorPlugin
type AttestorConfig struct {
	// Cluster name where the agent lives
	Cluster string `hcl:"cluster"`
	// File path of PSAT
	TokenPath string `hcl:"token_path"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *attestorConfig {
	_ = "STUB: not implemented"
	return nil
}

type attestorConfig struct {
	cluster   string
	tokenPath string
}

// AidAttestation loads the PSAT token from the configured path
func (p *AttestorPlugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure decodes JSON config from request and populates AttestorPlugin with it
func (p *AttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (resp *configv1.ConfigureResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *AttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (resp *configv1.ValidateResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *AttestorPlugin) getConfig() (*attestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadTokenFromFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }
