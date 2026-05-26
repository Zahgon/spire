package x509pop

import (
	"context"
	"crypto/x509"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	identityproviderv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/identityprovider/v1"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "x509pop"

	// Default security limits to prevent resource exhaustion attacks
	defaultMaxIntermediates = 4
	defaultMaxRSAKeySize    = 8192
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Config struct {
	Mode              string   `hcl:"mode"`
	SVIDPrefix        *string  `hcl:"spiffe_prefix"`
	CABundlePath      string   `hcl:"ca_bundle_path"`
	CABundlePaths     []string `hcl:"ca_bundle_paths"`
	AgentPathTemplate string   `hcl:"agent_path_template"`
	MaxIntermediates  *int     `hcl:"max_intermediates"`
	MaxRSAKeySize     *int     `hcl:"max_rsa_key_size"`
}

type configuration struct {
	mode             string
	svidPrefix       string
	trustDomain      spiffeid.TrustDomain
	trustBundle      *x509.CertPool
	pathTemplate     *agentpathtemplate.Template
	maxIntermediates int
	maxRSAKeySize    int
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *configuration {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	m                sync.Mutex
	config           *configuration
	identityProvider identityproviderv1.IdentityProviderServiceClient
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) BrokerHostServices(broker pluginsdk.ServiceBroker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

// build up leaf certificate and list of intermediates

// Check intermediate count limit before parsing to prevent resource exhaustion

// Validate leaf certificate key size

// Validate intermediate certificate key size

// verify the chain of trust

// now that the leaf certificate is trusted, issue a challenge to the node
// to prove possession of the private key.

// receive and validate the challenge response

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger sets this plugin's logger
func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) getTrustBundle(ctx context.Context) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() (*configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func buildSelectorValues(leaf *x509.Certificate, chains [][]*x509.Certificate, sanSelectors map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Used to avoid duplicating selectors.

// Iterate over all the certs in the chain (skip leaf at the 0 index)

// If the same fingerprint is generated, continue with the next certificate, because
// a selector should have been already created for it.

func (p *Plugin) parseUriSanSelectors(leaf *x509.Certificate, trustDomain string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Checks if an RSA public key exceeds the maximum allowed size.
func validateRSAKeySize(cert *x509.Certificate, maxRSAKeySize int) error {
	_ = "STUB: not implemented"
	return nil
}
