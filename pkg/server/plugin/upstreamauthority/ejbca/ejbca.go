package ejbca

import (
	"context"
	"crypto/x509"
	"sync"

	ejbcaclient "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

var (
	// This compile-time assertion ensures the plugin conforms properly to the
	// pluginsdk.NeedsLogger interface.
	_ pluginsdk.NeedsLogger = (*Plugin)(nil)
)

const (
	pluginName = "ejbca"
)

type newEjbcaAuthenticatorFunc func(*Config) (ejbcaclient.Authenticator, error)
type getEnvFunc func(string) string
type readFileFunc func(string) ([]byte, error)

// Plugin implements the UpstreamAuthority plugin
type Plugin struct {
	// UnimplementedUpstreamAuthorityServer is embedded to satisfy gRPC
	upstreamauthorityv1.UnimplementedUpstreamAuthorityServer

	// UnimplementedConfigServer is embedded to satisfy gRPC
	configv1.UnimplementedConfigServer

	config    *Config
	configMtx sync.RWMutex

	// The logger received from the framework via the SetLogger method
	logger hclog.Logger

	client ejbcaClient

	hooks struct {
		newAuthenticator newEjbcaAuthenticatorFunc
		getEnv           getEnvFunc
		readFile         readFileFunc
	}
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

// Config defines the configuration for the plugin.
type Config struct {
	Hostname               string `hcl:"hostname" json:"hostname"`
	CaCertPath             string `hcl:"ca_cert_path" json:"ca_cert_path"`
	ClientCertPath         string `hcl:"client_cert_path" json:"client_cert_path"`
	ClientCertKeyPath      string `hcl:"client_cert_key_path" json:"client_cert_key_path"`
	CAName                 string `hcl:"ca_name" json:"ca_name"`
	EndEntityProfileName   string `hcl:"end_entity_profile_name" json:"end_entity_profile_name"`
	CertificateProfileName string `hcl:"certificate_profile_name" json:"certificate_profile_name"`
	DefaultEndEntityName   string `hcl:"end_entity_name" json:"end_entity_name"`
	AccountBindingID       string `hcl:"account_binding_id" json:"account_binding_id"`
}

func (p *Plugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

// If ClientCertPath or ClientCertKeyPath were not found in the main server conf file,
// load them from the environment.

// If ClientCertPath or ClientCertKeyPath were not present in either the conf file or
// the environment, return an error.

// New returns an instantiated EJBCA UpstreamAuthority plugin
func New() *Plugin { _ = "STUB: not implemented"; return nil }

// Configure configures the EJBCA UpstreamAuthority plugin. This is invoked by SPIRE when the plugin is
// first loaded. After the first invocation, it may be used to reconfigure the plugin.
func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger is called by the framework when the plugin is loaded and provides
// the plugin with a logger wired up to SPIRE's logging facilities.
func (p *Plugin) SetLogger(logger hclog.Logger) {
	_ = "STUB: not implemented"

	// MintX509CAAndSubscribe implements the UpstreamAuthority MintX509CAAndSubscribe RPC. Mints an X.509 CA and responds
	// with the signed X.509 CA certificate chain and upstream X.509 roots. The stream is kept open but new roots will
	// not be published unless the CA is rotated and a new X.509 CA is minted.
	//
	// Implementation note:
	//   - It's important that the EJBCA Certificate Profile and End Entity Profile are properly configured before
	//     using this plugin. The plugin does not attempt to configure these profiles.
	return
}

func (p *Plugin) MintX509CAAndSubscribe(req *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure the request using local state and the CSR

// x509CertificateChain contains the leaf CA certificate, then any intermediates up to but not including the root CA.

// The EJBCA UpstreamAuthority plugin does not support publishing JWT keys.
func (p *Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

// setConfig replaces the configuration atomically under a write lock.
func (p *Plugin) setConfig(config *Config) { _ = "STUB: not implemented"; return }

// getConfig gets the configuration under a read lock.
func (p *Plugin) getConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// setClient replaces the client atomically under a write lock.
func (p *Plugin) setClient(client ejbcaClient) { _ = "STUB: not implemented"; return }

// getEndEntityName calculates the End Entity Name based on the default_end_entity_name from the EJBCA UpstreamAuthority
// configuration. The possible values are:
// - cn: Uses the Common Name from the CSR's Distinguished Name.
// - dns: Uses the first DNS Name from the CSR's Subject Alternative Names (SANs).
// - uri: Uses the first URI from the CSR's Subject Alternative Names (SANs).
// - ip: Uses the first IP Address from the CSR's Subject Alternative Names (SANs).
// - Custom Value: Any other string will be directly used as the End Entity Name.
// If the default_end_entity_name is not set, the plugin will determine the End Entity Name in the same order as above.
func (p *Plugin) getEndEntityName(config *Config, csr *x509.CertificateRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 1. If the endEntityName option is set, determine the end entity name based on the option
// 2. If the endEntityName option is not set, determine the end entity name based on the CSR

// cn: Use the CommonName from the CertificateRequest's DN

// dns: Use the first DNSName from the CertificateRequest's DNSNames SANs

// uri: Use the first URI from the CertificateRequest's URI Sans

// ip: Use the first IPAddress from the CertificateRequest's IPAddresses SANs

// End of defaults; if the endEntityName option is set to anything but cn, dns, or uri, use the option as the end entity name

// If we get here, we were unable to determine the end entity name

// parseEjbcaError parses an error returned by the EJBCA API and returns a gRPC status error.
func (p *Plugin) parseEjbcaError(detail string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// generateRandomString generates a random string of the specified length
func generateRandomString(length int) (string, error) { _ = "STUB: not implemented"; return "", nil }
