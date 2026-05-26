package gcpcas

import (
	"context"
	"sync"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	"cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	// The name of the plugin
	pluginName    = "gcp_cas"
	publicKeyType = "PUBLIC KEY"
)

// BuiltIn constructs a catalog Plugin using a new instance of this plugin.
func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type CertificateAuthoritySpec struct {
	Project    string `hcl:"project_name"`
	Location   string `hcl:"region_name"`
	CaPool     string `hcl:"ca_pool"`
	LabelKey   string `hcl:"label_key"`
	LabelValue string `hcl:"label_value"`
}

func (spec *CertificateAuthoritySpec) caParentPath(caPool string) string {
	_ = "STUB: not implemented"
	return ""
}

func (spec *CertificateAuthoritySpec) caPoolParentPath() string {
	_ = "STUB: not implemented"
	return ""
}

type Configuration struct {
	RootSpec CertificateAuthoritySpec `hcl:"root_cert_spec,block"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

// Without a project and location, we can never locate CAs

// Even LabelKey/Value pair is necessary

type CAClient interface {
	CreateCertificate(ctx context.Context, req *privatecapb.CreateCertificateRequest) (*privatecapb.Certificate, error)
	LoadCertificateAuthorities(ctx context.Context, spec CertificateAuthoritySpec) ([]*privatecapb.CertificateAuthority, error)
}

type Plugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	// mu is a mutex that protects the configuration. Plugins may at some point
	// need to support hot-reloading of configuration (by receiving another
	// call to Configure). So we need to prevent the configuration from
	// being used concurrently and make sure it is updated atomically.
	mu     sync.Mutex
	config *Configuration

	log hclog.Logger

	hook struct {
		getClient func(ctx context.Context) (CAClient, error)
	}
}

// These are compile time assertions that the plugin matches the interfaces the
// catalog requires to provide the plugin with a logger and host service
// broker as well as the UpstreamAuthority itself.
var _ pluginsdk.NeedsLogger = (*Plugin)(nil)
var _ upstreamauthorityv1.UpstreamAuthorityServer = (*Plugin)(nil)

func New() *Plugin { _ = "STUB: not implemented"; return nil }

// SetLogger will be called by the catalog system to provide the plugin with
// a logger when it is loaded. The logger is wired up to the SPIRE core
// logger
func (p *Plugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Mints an X.509 CA and responds with the signed X.509 CA certificate
	// chain and upstream X.509 roots. If supported by the implementation,
	// subsequent responses on the stream contain upstream X.509 root updates,
	// otherwise the RPC is completed after sending the initial response.
	//
	// Implementation note:
	// The stream should be kept open in the face of transient errors
	// encountered while tracking changes to the upstream X.509 roots as SPIRE
	// core will not reopen a closed stream until the next X.509 CA rotation.
	return
}

func (p *Plugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishJWTKeyAndSubscribe is not yet supported. It will return with GRPC Unimplemented error
func (p *Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	// Parse HCL config payload into config struct
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() (*Configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Plugin) mintX509CA(ctx context.Context, csr []byte, preferredTTL int32) (*upstreamauthorityv1.MintX509CAResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't want to use revoked, disabled or pending deletion CAs
// In short, we only need CAs that are in enabled state

// we want the CA that is expiring the earliest
// so sort and grab the first one

// All the CAs that are eligible for signing are still trusted

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#SubjectAltNames

// this is 0, golint complains if it's explicitly set to 0 since it's the default value of an int32

// privatecapb.CertificateAuthority.Name is the full GCP path but the request below expects only the CA's ID

// chosenPool will be in the form of projects/PROJECT/locations/LOCATION/caPools/POOL/certificateAuthorities/
// after the path.Split call above.  We need to trim off the /certificateAuthorities/ part for the request below

// certificate_id is required when using CertificateAuthority Enterprise tier. We generate a unique ID
// from the CSR public key. Same CSR will always produce the same ID.

// Convert the hash into a string that matches the `[a-zA-Z0-9_-]{1,63}` requirement for GCP API.

// https://pkg.go.dev/cloud.google.com/go/security/privateca/apiv1#CertificateAuthorityClient.CreateCertificate

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#Certificate

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#Certificate_Config

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CertificateConfig

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#PublicKey

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CertificateConfig_SubjectConfig

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#X509Parameters

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#X509Parameters_CaOptions

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#KeyUsage

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#KeyUsage_KeyUsageOptions

// All else comprises the chain (including the issued certificate)
// We don't include the root, since we pack that into the trust bundle.

// The last certificate returned from the chain is the root, so we seed the trust bundle with that.

// Then we append all the extra cert roots we loaded

// The last element in the PemCaCertificates is the root of this particular chain
// Note. We don't just use the CAs matched by labels from GCP because they could be
// intermediate CAs. If so, some of the libraries including OpenSSL will fail to
// validate them by default.
// Please refer to "X509_V_FLAG_PARTIAL_CHAIN" in
//    https://www.openssl.org/docs/man1.1.1/man3/X509_VERIFY_PARAM_set_flags.html

// We may well have specified multiple paths to the same root.

func getClient(ctx context.Context) (CAClient, error) {
	_ = "STUB: not implemented"
	// https://cloud.google.com/docs/authentication/production#go
	// The client creation implicitly uses Application Default Credentials (ADC) for authentication
	return *new(CAClient), nil
}

type gcpCAClient struct {
	pcaClient *privateca.CertificateAuthorityClient
}

func (client *gcpCAClient) CreateCertificate(ctx context.Context, req *privatecapb.CreateCertificateRequest) (*privatecapb.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *gcpCAClient) LoadCertificateAuthorities(ctx context.Context, spec CertificateAuthoritySpec) ([]*privatecapb.CertificateAuthority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the config has a ca pool provided only look for CAs in that pool, otherwise search each pool in the region

// https://pkg.go.dev/cloud.google.com/go/security/privateca/apiv1#CertificateAuthorityClient.ListCertificateAuthorities

// if there are cas in multiple pools that match our filter we need to throw an error

// There is "OrderBy" option, but it seems to work only for the name field
// So we will have to sort it by expiry timestamp at our end

func (client *gcpCAClient) listCaPools(ctx context.Context, spec CertificateAuthoritySpec) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterOutNonEnabledCAs(cas []*privatecapb.CertificateAuthority) []*privatecapb.CertificateAuthority {
	_ = "STUB: not implemented"
	return nil
}

// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CertificateAuthority_State
// Only CA in enabled state can issue certificates

// Sort in-place by ascending order of expiry time of CAs
func sortCAsByExpiryTime(cas []*privatecapb.CertificateAuthority) {
	_ = "STUB: not implemented"
	return
}
