package gcpiit

import (
	"context"
	"sync"

	"github.com/go-jose/go-jose/v4"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/plugin/gcp"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	nodeattestorbase "github.com/spiffe/spire/pkg/server/plugin/nodeattestor/base"
	"google.golang.org/api/compute/v1"
)

const (
	pluginName                  = "gcp_iit"
	tokenAudience               = "spire-gcp-node-attestor" //nolint: gosec // false positive
	googleCertURL               = "https://www.googleapis.com/oauth2/v1/certs"
	defaultMaxMetadataValueSize = 128
)

// Per GCP documentation, IITs are always signed using the RS256 signature algorithm:
// https://cloud.google.com/compute/docs/instances/verifying-instance-identity#verify_signature
var allowedJWTSignatureAlgorithms = []jose.SignatureAlgorithm{jose.RS256}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IITAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type jwksRetriever interface {
	retrieveJWKS(context.Context) (*jose.JSONWebKeySet, error)
}

type computeEngineClient interface {
	fetchInstanceMetadata(ctx context.Context, instanceMetadata gcp.ComputeEngine, serviceAccountFile string) (*compute.Instance, error)
}

// IITAttestorPlugin implements node attestation for agents running in GCP.
type IITAttestorPlugin struct {
	nodeattestorbase.Base
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	config        *IITAttestorConfig
	log           hclog.Logger
	mtx           sync.Mutex
	jwksRetriever jwksRetriever
	client        computeEngineClient
}

// IITAttestorConfig is the config for IITAttestorPlugin.
type IITAttestorConfig struct {
	idPathTemplate      *agentpathtemplate.Template
	trustDomain         spiffeid.TrustDomain
	allowedLabelKeys    map[string]bool
	allowedMetadataKeys map[string]bool

	ProjectIDAllowList   []string `hcl:"projectid_allow_list"`
	AgentPathTemplate    string   `hcl:"agent_path_template"`
	UseInstanceMetadata  bool     `hcl:"use_instance_metadata"`
	AllowedLabelKeys     []string `hcl:"allowed_label_keys"`
	AllowedMetadataKeys  []string `hcl:"allowed_metadata_keys"`
	MaxMetadataValueSize int      `hcl:"max_metadata_value_size"`
	ServiceAccountFile   string   `hcl:"service_account_file"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *IITAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new IITAttestorPlugin.
func New() *IITAttestorPlugin { _ = "STUB: not implemented"; return nil }

// SetLogger sets up plugin logging
func (p *IITAttestorPlugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Attest implements the server side logic for the gcp iit node attestation plugin.
	return
}

func (p *IITAttestorPlugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure configures the IITAttestorPlugin.
func (p *IITAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IITAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IITAttestorPlugin) getConfig() (*IITAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInstanceSelectorValues(config *IITAttestorConfig, instance *compute.Instance) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keyValue struct {
	key   string
	value string
}

func validateAttestationAndExtractIdentityMetadata(stream nodeattestorv1.NodeAttestor_AttestServer, jwks *jose.JSONWebKeySet) (gcp.IdentityToken, error) {
	_ = "STUB: not implemented"
	return *new(gcp.IdentityToken), nil
}

func getInstanceTags(instance *compute.Instance) []string { _ = "STUB: not implemented"; return nil }

func getInstanceLabels(instance *compute.Instance, allowedKeys map[string]bool) []keyValue {
	_ = "STUB: not implemented"
	return nil
}

func getInstanceMetadata(instance *compute.Instance, allowedKeys map[string]bool, maxValueSize int) ([]keyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeSelectorValue(key string, value ...string) string { _ = "STUB: not implemented"; return "" }

type googleComputeEngineClient struct{}

func (c googleComputeEngineClient) fetchInstanceMetadata(ctx context.Context, instanceMetadata gcp.ComputeEngine, serviceAccountFile string) (*compute.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c googleComputeEngineClient) getService(ctx context.Context, serviceAccountFile string) (*compute.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
