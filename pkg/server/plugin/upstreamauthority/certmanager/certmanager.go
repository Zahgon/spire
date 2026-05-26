package certmanager

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	pluginName = "cert-manager"
)

// BuiltIn constructs a catalog.BuiltIn using a new instance of this plugin.
func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Configuration struct {
	// Options which are used for configuring the target issuer to sign requests.
	// The CertificateRequest will be created in the configured namespace.
	IssuerName  string `hcl:"issuer_name" json:"issuer_name"`
	IssuerKind  string `hcl:"issuer_kind" json:"issuer_kind"`
	IssuerGroup string `hcl:"issuer_group" json:"issuer_group"`
	Namespace   string `hcl:"namespace" json:"namespace"`

	// File path to the kubeconfig used to build the generic Kubernetes client.
	KubeConfigFilePath string `hcl:"kube_config_file" json:"kube_config_file"`
}

func (p *Plugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

// namespace is a required field

// issuer_name is a required field

// If no issuer_kind given, default to Issuer

// If no issuer_group given, default to cert-manager.io

// Event hooks used by unit tests to coordinate goroutines
type hooks struct {
	newClient         func(configPath string) (client.Client, error)
	onCreateCR        func()
	onCleanupStaleCRs func()
}

type Plugin struct {
	// gRPC requires embedding either the "Unimplemented" or "Unsafe" stub as
	// a way of opting in or out of forward build compatibility.
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	log    hclog.Logger
	config *Configuration
	mtx    sync.RWMutex

	// trustDomain is the trust domain of this SPIRE server. Used to label
	// CertificateRequests to be cleaned-up
	trustDomain string

	// cmclient is a generic Kubernetes client for interacting with the
	// cert-manager APIs
	cmclient client.Client

	// Used for synchronization in unit tests
	hooks hooks
}

func New() *Plugin {
	_ = "STUB: not implemented"

	// noop hooks to avoid nil checks
	return nil
}

// SetLogger will be called by the catalog system to provide the plugin with
// a logger when it is loaded. The logger is wired up to the SPIRE core
// logger
func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Used for adding labels to created CertificateRequests, which can be listed
// for cleanup.

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Build the CertificateRequest object and create it

// Poll the CertificateRequest until it is signed. If not signed after 300
// polls, error.

// ~1.25 mins

// If the request has been denied, then return error here

// If the request has failed, then return error here

// If the Certificate exists on the request then it is ready.

// Parse signed certificate chain and CA certificate from CertificateRequest

// If the configured issuer did not populate the CA on the request we cannot
// build the upstream roots. We can only error here.

// PublishJWTKey is not implemented by the wrapper and returns a codes.Unimplemented status
func (*Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func newCertManagerClient(configPath string) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// Build a generic Kubernetes client which has the cert-manager.io schemas
// installed

func getKubeConfig(configPath string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
