package azureimds

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/plugin/azure"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	nodeattestorbase "github.com/spiffe/spire/pkg/server/plugin/nodeattestor/base"
)

const (
	pluginName = "azure_imds"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IMDSAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type SecretAuthConfig struct {
	AppID     string `hcl:"app_id" json:"app_id"`
	AppSecret string `hcl:"app_secret" json:"app_secret"`
}

type TokenAuthConfig struct {
	TokenPath string `hcl:"token_path" json:"token_path"`
	AppID     string `hcl:"app_id" json:"app_id"`
}

type TenantConfig struct {
	AuthType                string            `hcl:"auth_type" json:"auth_type"`
	SecretAuth              *SecretAuthConfig `hcl:"secret_auth" json:"secret_auth"`
	TokenAuth               *TokenAuthConfig  `hcl:"token_auth" json:"token_auth"`
	AllowedTags             []string          `hcl:"allowed_vm_tags" json:"allowed_vm_tags"`
	RestrictToSubscriptions []*string         `hcl:"restrict_to_subscriptions" json:"restrict_to_subscriptions"`
}

type IMDSAttestorConfig struct {
	Tenants           map[string]*TenantConfig `hcl:"tenants" json:"tenants"`
	AgentPathTemplate string                   `hcl:"agent_path_template" json:"agent_path_template"`
}

type tenantConfig struct {
	client                  apiClient
	allowedTags             map[string]struct{}
	restrictToSubscriptions map[string]struct{}
}

type imdsAttestorConfig struct {
	td             spiffeid.TrustDomain
	tenants        map[string]*tenantConfig
	idPathTemplate *agentpathtemplate.Template
}

func (t *tenantConfig) subscriptionAllowed(subscriptionID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *IMDSAttestorPlugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *imdsAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// Use tenant-specific credentials for resolving selectors

// If credentials are not configured then selectors won't be gathered.

type IMDSAttestorPlugin struct {
	nodeattestorbase.Base
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	mu     sync.RWMutex
	config *imdsAttestorConfig

	hooks struct {
		tenantIdMap         map[string]string
		newClient           func(azcore.TokenCredential) (apiClient, error)
		fetchCredential     func(string) (azcore.TokenCredential, error)
		validateAttestedDoc func(context.Context, *azure.AttestedDocument) (*azure.AttestedDocumentContent, error)
		lookupTenantID      func(string) (string, error)
	}
}

var _ nodeattestorv1.NodeAttestorServer = (*IMDSAttestorPlugin)(nil)

func New() *IMDSAttestorPlugin { _ = "STUB: not implemented"; return nil }

func (p *IMDSAttestorPlugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *IMDSAttestorPlugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	// receive initial empty payload
	return nil
}

// create a 32byte nonce

// send nonce back to agent

// receive the attestation payload

// Get the challenge response which contains the attested document and metadata

// parse the document

// if the query hint has a domain look up the tenant id

// Before doing the work to validate the token, ensure that the vmID has not already been used.

func (p *IMDSAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IMDSAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IMDSAttestorPlugin) getConfig() (*imdsAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildSelectors(ctx context.Context, tenant *tenantConfig, vmssName *string, vmID string, subscriptionID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// build up a unique map of selectors. this is easier than deduping
		// individual selectors (e.g. the virtual network for each interface)
		nil
}

// Get the VMSS Instance or Virtual Machine

// add tag selectors

// add network interface selectors

// sort and return selectors
