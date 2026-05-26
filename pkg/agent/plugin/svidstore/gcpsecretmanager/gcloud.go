package gcpsecretmanager

import (
	"context" //nolint: gosec // We use sha1 to hash trust domain names in 128 bytes to avoid secret label restrictions
	"sync"

	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/hcl/token"
	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "gcp_secretmanager"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *SecretManagerPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

func New() *SecretManagerPlugin { _ = "STUB: not implemented"; return nil }

func newPlugin(newSecretManagerClient func(context.Context, string) (secretManagerClient, error)) *SecretManagerPlugin {
	_ = "STUB: not implemented"
	return nil
}

type Configuration struct {
	ServiceAccountFile string                 `hcl:"service_account_file" json:"service_account_file"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions" json:",omitempty"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

type SecretManagerPlugin struct {
	svidstorev1.UnsafeSVIDStoreServer
	configv1.UnsafeConfigServer

	log                 hclog.Logger
	mtx                 sync.RWMutex
	secretManagerClient secretManagerClient
	tdHash              string

	hooks struct {
		newSecretManagerClient func(context.Context, string) (secretManagerClient, error)
	}
}

func (p *SecretManagerPlugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Configure configures the SecretManagerPlugin.
	return
}

func (p *SecretManagerPlugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gcp secret manager does not allow ".", hash td as label
//nolint: gosec // We use sha1 to hash trust domain names in 128 bytes to avoid secret label restrictions

func (p *SecretManagerPlugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PutX509SVID puts the specified X509-SVID in the configured Google Cloud Secrets Manager
func (p *SecretManagerPlugin) PutX509SVID(ctx context.Context, req *svidstorev1.PutX509SVIDRequest) (*svidstorev1.PutX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get secret, if it does not exist, a secret is created

// Secret not found, create it

// DeleteX509SVID deletes a secret in the configured Google Cloud Secret manager
func (p *SecretManagerPlugin) DeleteX509SVID(ctx context.Context, req *svidstorev1.DeleteX509SVIDRequest) (*svidstorev1.DeleteX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSecret gets secret from Google Cloud and validates if it has `spire-svid` label with hashed trust domain as value,
// nil if not found
func getSecret(ctx context.Context, client secretManagerClient, secretName string, tdHash string) (*secretmanagerpb.Secret, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Verify that secret contains "spire-svid" label and it is enabled

func (p *SecretManagerPlugin) shouldSetPolicy(ctx context.Context, secretName string, opt *secretOptions, secretFound bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Expecting a single Service account as member

func (p *SecretManagerPlugin) setIamPolicy(ctx context.Context, secretName string, opt *secretOptions) error {
	_ = "STUB: not implemented"
	// Create a policy without conditions and a single binding
	return nil
}

type secretOptions struct {
	projectID      string
	name           string
	roleName       string
	serviceAccount string
	replication    *secretmanagerpb.Replication
}

// parent gets parent in the format `projects/*`
func (s *secretOptions) parent() string { _ = "STUB: not implemented"; return "" }

// secretName gets secret name in format `projects/*/secrets/*`
func (s *secretOptions) secretName() string { _ = "STUB: not implemented"; return "" }

func optionsFromSecretData(selectorData []string) (*secretOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Getting secret name and project, both are required.

// example: "serviceAccount:project-id@appspot.gserviceaccount.com"

// Avoid adding empty strings as region

func validateLabels(labels map[string]string, tdHash string) bool {
	_ = "STUB: not implemented"
	return false
}

// expectedBindingMembers ensures that there is exactly one binding member, and
// that it matches the provided service account name
func expectedBindingMembers(bindingMembers []string, serviceAccount string) bool {
	_ = "STUB: not implemented"
	return false
}
