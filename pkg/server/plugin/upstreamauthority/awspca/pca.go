package awspca

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/hashicorp/go-hclog"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	// The name of the plugin
	pluginName = "aws_pca"
	// The header and footer type for a PEM-encoded CSR
	csrRequestType = "CERTIFICATE REQUEST"
	// The default CA signing template to use.
	// The SPIRE server intermediate CA can sign end-entity SVIDs only.
	defaultCASigningTemplateArn = "arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen0/V1"
	// Max certificate issuance wait duration
	maxCertIssuanceWaitDur = 3 * time.Minute
)

type newACMPCAClientFunc func(context.Context, *Configuration) (PCAClient, error)
type certificateIssuedWaitRetryFunc func(context.Context, *acmpca.GetCertificateInput, *acmpca.GetCertificateOutput, error) (bool, error)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *PCAPlugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

// Configuration provides configuration context for the plugin
type Configuration struct {
	Region                  string `hcl:"region" json:"region"`
	Endpoint                string `hcl:"endpoint" json:"endpoint"`
	CertificateAuthorityARN string `hcl:"certificate_authority_arn" json:"certificate_authority_arn"`
	SigningAlgorithm        string `hcl:"signing_algorithm" json:"signing_algorithm"`
	CASigningTemplateARN    string `hcl:"ca_signing_template_arn" json:"ca_signing_template_arn"`
	AssumeRoleARN           string `hcl:"assume_role_arn" json:"assume_role_arn"`
	SupplementalBundlePath  string `hcl:"supplemental_bundle_path" json:"supplemental_bundle_path"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

// PCAPlugin is the main representation of this upstreamauthority plugin
type PCAPlugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	mtx       sync.Mutex
	pcaClient PCAClient
	config    *configuration

	hooks struct {
		clock       clock.Clock
		newClient   newACMPCAClientFunc
		waitRetryFn certificateIssuedWaitRetryFunc
	}
}

type configuration struct {
	certificateAuthorityArn string
	signingAlgorithm        string
	caSigningTemplateArn    string
	supplementalBundle      []*x509.Certificate
}

// New returns an instantiated plugin
func New() *PCAPlugin { _ = "STUB: not implemented"; return nil }

func newPlugin(newClient newACMPCAClientFunc, waitRetryFn certificateIssuedWaitRetryFunc) *PCAPlugin {
	_ = "STUB: not implemented"
	return nil
}

func (p *PCAPlugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Configure sets up the plugin for use as an upstream authority
	return
}

func (p *PCAPlugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the client

// Perform a check for the presence of the CA

// Ensure the CA is set to ACTIVE

// If a signing algorithm has been provided, use it.
// Otherwise, fall back to the pre-configured value on the CA

// If a CA signing template ARN has been provided, use it.
// Otherwise, fall back to the default value (PathLen=0)

// Set local vars

func (p *PCAPlugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MintX509CA mints an X509CA by submitting the CSR to ACM to be signed by the certificate authority
func (p *PCAPlugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Have ACM sign the certificate

// Using the output of the `IssueCertificate` call, poll ACM until
// the certificate has been issued

// Finally get the certificate contents

// Parse the cert from the response

// Parse the chain from the response

// ACM's API outputs the certificate chain from a GetCertificate call in the following
// order: A (signed by B) -> B (signed by ROOT) -> ROOT.
// For SPIRE, the certificate chain will always include at least one certificate (the root),
// but may include other intermediates between SPIRE and the ROOT.
// See https://docs.aws.amazon.com/cli/latest/reference/acm-pca/import-certificate-authority-certificate.html
// and https://docs.aws.amazon.com/cli/latest/reference/acm-pca/get-certificate.html

// The last certificate returned from the chain is the root.

// All else comprises the chain (including the issued certificate)

// PublishJWTKey is not implemented by the wrapper and returns a codes.Unimplemented status
func (*PCAPlugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PCAPlugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PCAPlugin) getConfig() (*configuration, error) { _ = "STUB: not implemented"; return nil, nil }
