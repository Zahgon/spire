package awssecret

import (
	"context"
	"crypto/x509"
	"sync"

	"github.com/andres-erbsen/clock"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"github.com/spiffe/spire/pkg/common/x509svid"
)

const (
	pluginName = "awssecret"

	CoreConfigRequired             = "server core configuration is required"
	CoreConfigTrustdomainRequired  = "server core configuration must contain trust_domain"
	CoreConfigTrustdomainMalformed = "server core configuration trust_domain is malformed"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Configuration struct {
	Region          string `hcl:"region" json:"region"`
	CertFileARN     string `hcl:"cert_file_arn" json:"cert_file_arn"`
	KeyFileARN      string `hcl:"key_file_arn" json:"key_file_arn"`
	BundleFileARN   string `hcl:"bundle_file_arn" json:"bundle_file_arn"`
	AccessKeyID     string `hcl:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key" json:"secret_access_key"`
	SecurityToken   string `hcl:"secret_token" json:"secret_token"`
	AssumeRoleARN   string `hcl:"assume_role_arn" json:"assume_role_arn"`
}

func (p *Plugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	mtx           sync.RWMutex
	upstreamCerts []*x509.Certificate
	bundleCerts   []*x509.Certificate
	upstreamCA    *x509svid.UpstreamCA

	hooks struct {
		clock     clock.Clock
		getenv    func(string) string
		newClient func(ctx context.Context, config *Configuration, region string) (secretsManagerClient, error)
	}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func newPlugin(newClient func(ctx context.Context, config *Configuration, region string) (secretsManagerClient, error)) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: determine if the items before the lock contain configuration validation.

// set the AWS configuration and reset clients +
// Set local vars from config struct

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MintX509CAAndSubscribe mints an X509CA by signing presented CSR with root CA fetched from AWS Secrets Manager
func (p *Plugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishJWTKeyAndSubscribe is not implemented by the wrapper and returns a codes.Unimplemented status
func (p *Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) loadUpstreamCAAndCerts(trustDomain spiffeid.TrustDomain, keyPEMstr, certsPEMstr, bundleCertsPEMstr string) (*x509svid.UpstreamCA, []*x509.Certificate, []*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// pemutil guarantees at least one cert

// If there is no bundle payload configured then the value of certs
// must be a self-signed cert. We enforce this by requiring that there is
// exactly one certificate; this certificate is reused for the trust
// bundle and bundleCertsPEMstr is ignored

// If there is a bundle, instead of using the payload of cert_file_arn
// to populate the trust bundle, we assume that certs is a chain of
// intermediates and populate the trust bundle with roots from
// bundle_file_arn

// If we get to this point we've successfully validated that:
// - cert_file_arn contains a single self-signed certificate OR
// - cert_file_arn contains a chain of certificates which terminate at a root
//   which is provided in bundle_file_arn

func fetchFromSecretsManager(ctx context.Context, config *Configuration, sm secretsManagerClient) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}
