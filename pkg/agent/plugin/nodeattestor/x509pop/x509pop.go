package x509pop

import (
	"context"
	"crypto"
	"sync"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "x509pop"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type configData struct {
	privateKey         crypto.PrivateKey
	attestationPayload []byte
}

type Config struct {
	PrivateKeyPath       string `hcl:"private_key_path"`
	CertificatePath      string `hcl:"certificate_path"`
	IntermediatesPath    string `hcl:"intermediates_path"`
	SpiffeEndpointSocket string `hcl:"spiffe_endpoint_socket"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	m sync.Mutex
	c *Config
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// send the attestation data back to the agent

// receive challenge

// calculate and send the challenge response

func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure the configuration produces valid data

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() *Config { _ = "STUB: not implemented"; return nil }

func (p *Plugin) loadConfigData(ctx context.Context) (*configData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: this needs more attention.  Parts of it might belong in buildConfig
func loadConfigData(ctx context.Context, config *Config, inAttest bool) (*configData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Append intermediate certificates if IntermediatesPath is set.
