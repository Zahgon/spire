package gcpiit

import (
	"context"
	"sync"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	defaultIdentityTokenHost     = "metadata.google.internal"
	identityTokenURLPathTemplate = "/computeMetadata/v1/instance/service-accounts/%s/identity"
	identityTokenAudience        = "spire-gcp-node-attestor" //nolint: gosec // false positive
	defaultServiceAccount        = "default"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IITAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

// IITAttestorPlugin implements GCP nodeattestation in the agent.
type IITAttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mtx    sync.RWMutex
	config *IITAttestorConfig
}

// IITAttestorConfig configures a IITAttestorPlugin.
type IITAttestorConfig struct {
	IdentityTokenHost string `hcl:"identity_token_host"`
	ServiceAccount    string `hcl:"service_account"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *IITAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// NewIITAttestorPlugin creates a new IITAttestorPlugin.
func New() *IITAttestorPlugin { _ = "STUB: not implemented"; return nil }

// AidAttestation fetches attestation data from the GCP metadata server and sends an attestation response
// on given stream.
func (p *IITAttestorPlugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *IITAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IITAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IITAttestorPlugin) getConfig() (*IITAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// identityTokenURL creates the URL to find an instance identity document given the
// host of the GCP metadata server and the service account the instance is running as.
func identityTokenURL(host, serviceAccount string) string { _ = "STUB: not implemented"; return "" }

func retrieveInstanceIdentityToken(url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
