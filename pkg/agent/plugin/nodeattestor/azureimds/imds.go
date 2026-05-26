package azureimds

import (
	"context"
	"sync"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/plugin/azure"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "azure_imds"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IMDSAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type IMDSAttestorConfig struct {
	// TenantDomain is the domain of the tenant in which the VM is running.
	TenantDomain string `hcl:"tenant_domain"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *IMDSAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

type IMDSAttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mu     sync.RWMutex
	config *IMDSAttestorConfig

	hooks struct {
		fetchAttestedDocument func(azure.HTTPClient, string) (*azure.AttestedDocument, error)
		fetchComputeMetadata  func(azure.HTTPClient) (*azure.InstanceMetadata, error)
	}
}

func New() *IMDSAttestorPlugin { _ = "STUB: not implemented"; return nil }

func (p *IMDSAttestorPlugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// send initial payload, this is just so we can receive a challenge containing the nonce

// receive challenge containing the nonce which we will use to fetch the attested document

// Get the attested document

// Get the compute metadata

func (p *IMDSAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IMDSAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IMDSAttestorPlugin) getConfig() (*IMDSAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
