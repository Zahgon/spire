package azuremsi

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/jwtutil"
	"github.com/spiffe/spire/pkg/common/plugin/azure"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	nodeattestorbase "github.com/spiffe/spire/pkg/server/plugin/nodeattestor/base"
)

const (
	pluginName = "azure_msi"

	// MSI tokens have the not-before ("nbf") claim. If there are clock
	// differences between the agent and server then token validation may fail
	// unless we give a little leeway. Tokens are valid for 8 hours, so a few
	// minutes extra in that direction does not seem like a big deal.
	tokenLeeway = time.Minute * 5

	keySetRefreshInterval = time.Hour
	azureOIDCIssuer       = "https://login.microsoftonline.com/common/"
)

var (
	reVirtualMachineID       = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Compute/virtualMachines/([^/]+)$`)
	reNetworkSecurityGroupID = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/networkSecurityGroups/([^/]+)$`)
	reNetworkInterfaceID     = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/networkInterfaces/([^/]+)$`)
	reVirtualNetworkSubnetID = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/virtualNetworks/([^/]+)/subnets/([^/]+)$`)
	// Azure doesn't appear to publicly document which signature algorithms they use for MSI tokens,
	// but a couple examples online were showing RS256.
	// To ensure compatibility, accept the most common signature algorithms that are known to be secure.
	allowedJWTSignatureAlgorithms = []jose.SignatureAlgorithm{
		jose.RS256,
		jose.RS384,
		jose.RS512,
		jose.ES256,
		jose.ES384,
		jose.ES512,
		jose.PS256,
		jose.PS384,
		jose.PS512,
	}
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *MSIAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type TenantConfig struct {
	ResourceID     string `hcl:"resource_id" json:"resource_id"`
	SubscriptionID string `hcl:"subscription_id" json:"subscription_id"`
	AppID          string `hcl:"app_id" json:"app_id"`
	AppSecret      string `hcl:"app_secret" json:"app_secret"`
}

type MSIAttestorConfig struct {
	Tenants           map[string]*TenantConfig `hcl:"tenants" json:"tenants"`
	AgentPathTemplate string                   `hcl:"agent_path_template" json:"agent_path_template"`
}

type tenantConfig struct {
	resourceID string
	client     apiClient
}

type msiAttestorConfig struct {
	td             spiffeid.TrustDomain
	tenants        map[string]*tenantConfig
	idPathTemplate *agentpathtemplate.Template
}

func (p *MSIAttestorPlugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *msiAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// Use tenant-specific credentials for resolving selectors

// If credentials are not configured then selectors won't be gathered.

type MSIAttestorPlugin struct {
	nodeattestorbase.Base
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	log hclog.Logger

	mu     sync.RWMutex
	config *msiAttestorConfig

	hooks struct {
		now                   func() time.Time
		keySetProvider        jwtutil.KeySetProvider
		newClient             func(string, azcore.TokenCredential) (apiClient, error)
		fetchInstanceMetadata func(azure.HTTPClient) (*azure.InstanceMetadata, error)
		fetchCredential       func(string) (azcore.TokenCredential, error)
	}
}

var _ nodeattestorv1.NodeAttestorServer = (*MSIAttestorPlugin)(nil)

func New() *MSIAttestorPlugin { _ = "STUB: not implemented"; return nil }

func (p *MSIAttestorPlugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *MSIAttestorPlugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Before doing the work to validate the token, ensure that this MSI token
// has not already been used to attest an agent.

func (p *MSIAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *MSIAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *MSIAttestorPlugin) getConfig() (*msiAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *MSIAttestorPlugin) resolve(ctx context.Context, client apiClient, principalID string) ([]string, error) {
	_ = "STUB: not implemented"
	// Retrieve the resource belonging to the principal id.
	return nil, nil
}

// parse out the resource group and vm name from the resource ID

// build up a unique map of selectors. this is easier than deduping
// individual selectors (e.g. the virtual network for each interface)

// pull the VM information and gather selectors

// sort and return selectors

func getNetworkProfileSelectors(ctx context.Context, client apiClient, networkProfile *armcompute.NetworkProfile) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNetworkInterfaceSelectors(networkInterface *armnetwork.Interface) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseVirtualMachineID(id string) (resourceGroup, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseNetworkSecurityGroupID(id string) (resourceGroup, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseNetworkInterfaceID(id string) (resourceGroup, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseVirtualNetworkSubnetID(id string) (resourceGroup, networkName, subnetName string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func resourceGroupName(resourceGroup, name string) string { _ = "STUB: not implemented"; return "" }

func selectorValue(parts ...string) string { _ = "STUB: not implemented"; return "" }

func getTokenKeyID(token *jwt.JSONWebToken) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
