package azuremsi

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
	pluginName = "azure_msi"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *MSIAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type MSIAttestorConfig struct {
	// ResourceID assigned to the MSI token. This value is the intended
	// audience of the token, in other words, which service the token can be
	// used to authenticate with. Ideally deployments use the ID of an
	// application they registered with the active directory to limit the scope
	// of use of the token. A bogus value cannot be used; Azure makes sure the
	// resource ID is either an azure service ID or a registered app ID.
	ResourceID string `hcl:"resource_id"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *MSIAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

type MSIAttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mu     sync.RWMutex
	config *MSIAttestorConfig

	hooks struct {
		fetchMSIToken func(azure.HTTPClient, string) (string, error)
	}
}

func New() *MSIAttestorPlugin { _ = "STUB: not implemented"; return nil }

func (p *MSIAttestorPlugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Obtain an MSI token from the Azure Instance Metadata Service

func (p *MSIAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *MSIAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *MSIAttestorPlugin) getConfig() (*MSIAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
