package vault

import (
	"context"
	"crypto/x509"

	"github.com/hashicorp/go-hclog"
	vapi "github.com/hashicorp/vault/api"
)

const (
	EnvVaultAddr       = "VAULT_ADDR"
	EnvVaultToken      = "VAULT_TOKEN"
	EnvVaultClientCert = "VAULT_CLIENT_CERT"
	EnvVaultClientKey  = "VAULT_CLIENT_KEY"
	EnvVaultCACert     = "VAULT_CACERT"
	EnvVaultNamespace  = "VAULT_NAMESPACE"
	// SPIRE-specific; not a standard Vault SDK environment variable.
	EnvVaultAppRoleID         = "VAULT_APPROLE_ID"
	EnvVaultAppRoleSecretID   = "VAULT_APPROLE_SECRET_ID" // #nosec G101
	EnvVaultTransitEnginePath = "VAULT_TRANSIT_ENGINE_PATH"

	defaultCertMountPoint    = "cert"
	defaultPKIMountPoint     = "pki"
	defaultTransitEnginePath = "transit"
	defaultAppRoleMountPoint = "approle"
	defaultK8sMountPoint     = "kubernetes"
)

type AuthMethod int

const (
	_ AuthMethod = iota
	CERT
	TOKEN
	APPROLE
	K8S
)

type TransitKeyType string

const (
	TransitKeyTypeRSA2048   TransitKeyType = "rsa-2048"
	TransitKeyTypeRSA4096   TransitKeyType = "rsa-4096"
	TransitKeyTypeECDSAP256 TransitKeyType = "ecdsa-p256"
	TransitKeyTypeECDSAP384 TransitKeyType = "ecdsa-p384"
)

type TransitHashAlgorithm string

const (
	TransitHashAlgorithmSHA256 TransitHashAlgorithm = "sha2-256"
	TransitHashAlgorithmSHA384 TransitHashAlgorithm = "sha2-384"
	TransitHashAlgorithmSHA512 TransitHashAlgorithm = "sha2-512"
	TransitHashAlgorithmNone   TransitHashAlgorithm = "none"
)

type TransitSignatureAlgorithm string

const (
	TransitSignatureAlgorithmNone     TransitSignatureAlgorithm = ""
	TransitSignatureAlgorithmPSS      TransitSignatureAlgorithm = "pss"
	TransitSignatureAlgorithmPKCS1v15 TransitSignatureAlgorithm = "pkcs1v15"
)

type KeyEntry struct {
	KeyName string
	// KeyType is the top-level type from Vault (e.g., "ecdsa-p256", "rsa-2048").
	KeyType string
	KeyData map[string]any
}

// ClientConfig represents configuration parameters for vault client
type ClientConfig struct {
	Logger hclog.Logger
	// vault client parameters
	ClientParams *ClientParams
}

type ClientParams struct {
	// A URL of Vault server. (e.g., https://vault.example.com:8443/)
	VaultAddr string
	// Name of mount point where PKI secret engine is mounted. (e.e., /<mount_point>/ca/pem )
	PKIMountPoint string
	// token string to use when auth method is 'token'
	Token string
	// Name of mount point where TLS Cert auth method is mounted. (e.g., /auth/<mount_point>/login )
	CertAuthMountPoint string
	// Name of the Vault role.
	// If given, the plugin authenticates against only the named role
	CertAuthRoleName string
	// Path to a client certificate file to be used when auth method is 'cert'
	ClientCertPath string
	// Path to a client private key file to be used when auth method is 'cert'
	ClientKeyPath string
	// Path to a CA certificate file to be used when client verifies a server certificate
	CACertPath string
	// Name of mount point where AppRole auth method is mounted. (e.g., /auth/<mount_point>/login )
	AppRoleAuthMountPoint string
	// An identifier of AppRole
	AppRoleID string
	// A credential set of AppRole
	AppRoleSecretID string
	// Name of the mount point where Kubernetes auth method is mounted. (e.g., /auth/<mount_point>/login)
	K8sAuthMountPoint string
	// Name of the Vault role.
	// The plugin authenticates against the named role.
	K8sAuthRoleName string
	// Path to a K8s Service Account Token to be used when auth method is 'k8s'
	K8sAuthTokenPath string
	// If true, client accepts any certificates.
	// It should be used only test environment so on.
	TLSSkipVerify bool
	// MaxRetries controls the number of times to retry to connect
	// Set to 0 to disable retrying.
	// If the value is nil, to use the default in hashicorp/vault/api.
	MaxRetries *int
	// Name of the Vault namespace
	Namespace string
	// TransitEnginePath specifies the path to the transit engine to perform key operations.
	TransitEnginePath string
}

type Client struct {
	vaultClient  *vapi.Client
	ClientParams *ClientParams
}

// SignCSRResponse includes certificates which are generates by Vault
type SignCSRResponse struct {
	// A certificate requested to sign
	CACertPEM string
	// A certificate of CA(Vault)
	UpstreamCACertPEM string
	// Set of Upstream CA certificates
	UpstreamCACertChainPEM []string
}

// NewClientConfig returns a new *ClientConfig with default parameters.
func NewClientConfig(cp *ClientParams, logger hclog.Logger) (*ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAuthenticatedClient returns a new authenticated vault client with given authentication method
func (c *ClientConfig) NewAuthenticatedClient(method AuthMethod, renewCh chan struct{}) (client *Client, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleRenewToken handles renewing the vault token.
// if the token is non-renewable or renew failed, renewCh will be closed.
func handleRenewToken(vc *vapi.Client, sec *vapi.Secret, renewCh chan struct{}, logger hclog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigureTLS Configures TLS for Vault Client
func (c *ClientConfig) configureTLS(vc *vapi.Config) error { _ = "STUB: not implemented"; return nil }

// VaultClient returns the underlying vault API client.
func (c *Client) VaultClient() *vapi.Client { _ = "STUB: not implemented"; return nil }

// SetToken wraps vapi.Client.SetToken()
func (c *Client) SetToken(v string) { _ = "STUB: not implemented"; return }

// Auth authenticates to vault server with TLS certificate method
func (c *Client) Auth(path string, body map[string]any) (*vapi.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) LookupSelf(token string) (*vapi.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// don't care any parameters

// SignIntermediate requests sign-intermediate endpoint to generate certificate.
// ttl = TTL for Intermediate CA Certificate
// csr = Certificate Signing Request
// see: https://www.vaultproject.io/api/secret/pki/index.html#sign-intermediate
func (c *Client) SignIntermediate(ttl string, csr *x509.CertificateRequest) (*SignCSRResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// expect to be empty case when Vault is Root CA.

// CreateKey creates a new key in the specified transit secret engine
// See: https://developer.hashicorp.com/vault/api-docs/secret/transit#create-key
func (c *Client) CreateKey(ctx context.Context, keyName string, keyType TransitKeyType) error {
	_ = "STUB: not implemented"
	return nil
}

// SPIRE keys are never exportable
// SPIRE manages rotation; disable Vault-side auto-rotation

// DeleteKey deletes a key in the specified transit secret engine
// See: https://developer.hashicorp.com/vault/api-docs/secret/transit#update-key-configuration and https://developer.hashicorp.com/vault/api-docs/secret/transit#delete-key
func (c *Client) DeleteKey(ctx context.Context, keyName string) error {
	_ = "STUB: not implemented"
	return nil
}

// First, we need to enable deletion of the key. This is disabled by default.

// SignData signs the data using the transit engine key with the key name.
// See: https://developer.hashicorp.com/vault/api-docs/secret/transit#sign-data
func (c *Client) SignData(ctx context.Context, keyName string, data []byte, hashAlgo TransitHashAlgorithm, signatureAlgo TransitSignatureAlgorithm) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Vault adds an application specific prefix that we need to remove

// GetKeys returns all the keys of the transit engine.
// See: https://developer.hashicorp.com/vault/api-docs/secret/transit#list-keys
func (c *Client) GetKeys(ctx context.Context) ([]*KeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetKey returns a specific key from the transit engine.
// See: https://developer.hashicorp.com/vault/api-docs/secret/transit#read-key
func (c *Client) GetKey(ctx context.Context, keyName string) (*KeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Vault SDK deserializes JSON numbers as json.Number in map[string]any.
