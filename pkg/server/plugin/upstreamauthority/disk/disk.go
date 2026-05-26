package disk

import (
	"context"
	"crypto/x509"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"

	"github.com/andres-erbsen/clock"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"github.com/spiffe/spire/pkg/common/x509svid"
)

const (
	CoreConfigRequired             = "server core configuration is required"
	CoreConfigTrustDomainRequired  = "server core configuration must contain trust_domain"
	CoreConfigTrustDomainMalformed = "server core configuration trust_domain is malformed"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Configuration struct {
	trustDomain spiffeid.TrustDomain

	CertFilePath   string `hcl:"cert_file_path" json:"cert_file_path"`
	KeyFilePath    string `hcl:"key_file_path" json:"key_file_path"`
	BundleFilePath string `hcl:"bundle_file_path" json:"bundle_file_path"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	mtx        sync.Mutex
	config     *Configuration
	certs      *caCerts
	upstreamCA *x509svid.UpstreamCA

	// test hooks
	clock clock.Clock
}

type caCerts struct {
	certChain   []*x509.Certificate
	trustBundle []*x509.Certificate
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set local vars from config struct

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: provide more granular status codes

func (*Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) reloadCA() (*x509svid.UpstreamCA, *caCerts, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO: perhaps load this into the config
func (p *Plugin) loadUpstreamCAAndCerts(config *Configuration) (*x509svid.UpstreamCA, *caCerts, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// pemutil guarantees at least 1 cert

// If there is no bundle path configured then we assume we have
// a self-signed cert. We enforce this by requiring that there is
// exactly one cert. This cert is reused for the trust bundle and
// config.BundleFilePath is ignored

// Validate cert matches private key
