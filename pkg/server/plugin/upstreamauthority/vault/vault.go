package vault

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"

	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"github.com/spiffe/spire/pkg/server/common/vault"
)

const (
	pluginName = "vault"

	PluginConfigMalformed = "plugin configuration is malformed"
)

// BuiltIn constructs a catalog.BuiltIn using a new instance of this plugin.
func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Configuration struct {
	vault.BaseConfiguration `hcl:",squash"`

	// Name of the mount point where PKI secret engine is mounted. (e.g., /<mount_point>/ca/pem)
	PKIMountPoint string `hcl:"pki_mount_point" json:"pki_mount_point"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add field validations

// TODO: consider moving some elements of parseAuthMethod into config checking
// TODO: consider moving some elements of genClientParams into config checking
// TODO: consider moving some elements of NewClientConfig into config checking

type Plugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	mtx    *sync.RWMutex
	logger hclog.Logger

	authMethod vault.AuthMethod
	cc         *vault.ClientConfig
	vc         *vault.Client

	hooks struct {
		lookupEnv func(string) (string, bool)
	}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set PKI mount point

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) MintX509CAAndSubscribe(req *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// if renewCh has been closed, the token can not be renewed and may expire,
// it needs to re-authenticate to the Vault.

// Parse CACert in PEM format

// Parse PEM format data to get DER format data

// Since Vault v1.11.0, the signed CA certificate appears within the ca_chain
// https://github.com/hashicorp/vault/blob/v1.11.0/changelog/15524.txt

// PublishJWTKeyAndSubscribe is not implemented by the wrapper and returns a codes.Unimplemented status
func (*Plugin) PublishJWTKeyAndSubscribe(*upstreamauthorityv1.PublishJWTKeyRequest, upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}
