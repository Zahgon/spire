package gcpkms

import (
	"context" //nolint: gosec // We use sha1 to hash trust domain names in 128 bytes to avoid label value restrictions
	"sync"
	"time"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/andres-erbsen/clock"
	"github.com/hashicorp/go-hclog"
	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"google.golang.org/api/option"
)

const (
	pluginName = "gcp_kms"

	algorithmTag             = "algorithm"
	cryptoKeyNameTag         = "crypto_key_name"
	cryptoKeyVersionNameTag  = "crypto_key_version_name"
	cryptoKeyVersionStateTag = "crypto_key_version_state"
	scheduledDestroyTimeTag  = "scheduled_destroy_time"
	reasonTag                = "reason"

	disposeCryptoKeysFrequency    = time.Hour * 48
	keepActiveCryptoKeysFrequency = time.Hour * 6
	maxStaleDuration              = time.Hour * 24 * 14 // Two weeks.

	cryptoKeyNamePrefix = "spire-key"
	labelNameServerID   = "spire-server-id"
	labelNameLastUpdate = "spire-last-update"
	labelNameServerTD   = "spire-server-td"
	labelNameActive     = "spire-active"

	getPublicKeyMaxAttempts = 10
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type keyEntry struct {
	cryptoKey            *kmspb.CryptoKey
	cryptoKeyVersionName string
	publicKey            *keymanagerv1.PublicKey
}

type pluginHooks struct {
	newKMSClient func(context.Context, ...option.ClientOption) (cloudKeyManagementService, error)

	clk clock.Clock

	// Used for testing only.
	disposeCryptoKeysSignal    chan error
	enqueueDestructionSignal   chan error
	keepActiveCryptoKeysSignal chan error
	scheduleDestroySignal      chan error
	setInactiveSignal          chan error
}

type pluginData struct {
	customPolicy *iam.Policy3
	serverID     string
	tdHash       string
}

// Plugin is the main representation of this keymanager plugin.
type Plugin struct {
	keymanagerv1.UnsafeKeyManagerServer
	configv1.UnsafeConfigServer

	cancelTasks context.CancelFunc

	config    *Config
	configMtx sync.RWMutex

	entries    map[string]keyEntry
	entriesMtx sync.RWMutex

	pd    *pluginData
	pdMtx sync.RWMutex

	hooks           pluginHooks
	kmsClient       cloudKeyManagementService
	log             hclog.Logger
	scheduleDestroy chan string
}

// Config provides configuration context for the plugin.
type Config struct {
	// File path location where information about generated keys will be persisted.
	KeyIdentifierFile string `hcl:"key_identifier_file" json:"key_identifier_file"`

	// Key metadata used by the plugin.
	KeyIdentifierValue string `hcl:"key_identifier_value" json:"key_identifier_value"`

	// File path location to a custom IAM Policy (v3) that will be set to
	// created CryptoKeys.
	KeyPolicyFile string `hcl:"key_policy_file" json:"key_policy_file"`

	// KeyRing is the resource ID of the key ring where the keys managed by this
	// plugin reside, in the format projects/*/locations/*/keyRings/*.
	KeyRing string `hcl:"key_ring" json:"key_ring"`

	// Path to the service account file used to authenticate with the Cloud KMS
	// API. If not specified, the value of the GOOGLE_APPLICATION_CREDENTIALS
	// environment variable is used.
	ServiceAccountFile string `hcl:"service_account_file" json:"service_account_file"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

// New returns an instantiated plugin.
func New() *Plugin { _ = "STUB: not implemented"; return nil }

// newPlugin returns a new plugin instance.
func newPlugin(
	newKMSClient func(context.Context, ...option.ClientOption) (cloudKeyManagementService, error),
) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Close() error { _ = "STUB: not implemented"; return nil }

// Configure sets up the plugin.
func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Label values do not allow "." and have a maximum length of 63 characters.
// https://cloud.google.com/kms/docs/creating-managing-labels#requirements
// Hash the trust domain name to avoid restrictions.
//nolint: gosec // We use sha1 to hash trust domain names in 128 bytes to avoid label restrictions

// Cancel previous tasks in case of re-configure.

// Start long-running tasks.

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateKey creates a key in KMS. If a key already exists in the local storage,
// it is updated.
func (p *Plugin) GenerateKey(ctx context.Context, req *keymanagerv1.GenerateKeyRequest) (*keymanagerv1.GenerateKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKey returns the public key for a given key
func (p *Plugin) GetPublicKey(_ context.Context, req *keymanagerv1.GetPublicKeyRequest) (*keymanagerv1.GetPublicKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKeys returns the publicKey for all the keys.
func (p *Plugin) GetPublicKeys(context.Context, *keymanagerv1.GetPublicKeysRequest) (*keymanagerv1.GetPublicKeysResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger sets a logger.
func (p *Plugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// SignData creates a digital signature for the data to be signed.
	return
}

func (p *Plugin) SignData(ctx context.Context, req *keymanagerv1.SignDataRequest) (*keymanagerv1.SignDataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RSASSA-PSS is not supported by this plugin.
// See the comment in cryptoKeyVersionAlgorithmFromKeyType function for
// more details.

// Perform integrity verification.

// createKey creates a new CryptoKey with a new CryptoKeyVersion in Cloud KMS
// if there is not already a cached entry with the specified SPIRE Key ID.
// If the cache already has an entry with this SPIRE Key ID, a new
// CryptoKeyVersion is added to the corresponding CryptoKey in Cloud KMS and the
// old CryptoKeyVersion is enqueued for destruction.
// If there is a specified IAM policy through the KeyPolicyFile configuration,
// that policy is set to the created CryptoKey. If there is no IAM policy specified,
// a default policy is constructed and attached. This function requests Cloud KMS
// to get the public key of the created CryptoKeyVersion. A keyEntry is returned
// with the CryptoKey, CryptoKeyVersion and public key.
func (p *Plugin) createKey(ctx context.Context, spireKeyID string, keyType keymanagerv1.KeyType) (*keymanagerv1.PublicKey, error) {
	_ = "STUB: not implemented"
	// If we already have this SPIRE Key ID cached, a new CryptoKeyVersion is
	// added to the existing CryptoKey and the cache is updated. The old
	// CryptoKeyVersion is enqueued for destruction.
	return nil, nil
}

// addCryptoKeyVersionToCachedEntry adds a new CryptoKeyVersion to an existing
// CryptoKey, updating the cached entries.
func (p *Plugin) addCryptoKeyVersionToCachedEntry(ctx context.Context, entry keyEntry, spireKeyID string, keyType keymanagerv1.KeyType) (*keymanagerv1.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the algorithm has changed and update if needed.

// disposeCryptoKeys looks for active CryptoKeys that haven't been updated
// during the maxStaleDuration time window. Those keys are then enqueued for
// destruction.
func (p *Plugin) disposeCryptoKeys(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If the CryptoKey doesn't have any enabled CryptoKeyVersion, mark it
// as inactive so it's not returned future calls.

// No more enabled CryptoKeyVersions in this CryptoKey.

// disposeCryptoKeysTask will be run every 24hr.
// It will schedule the destruction of CryptoKeyVersions that have a
// spire-last-update label value older than two weeks.
// It will only schedule the destruction of CryptoKeyVersions belonging to the
// current trust domain but not the current server. The spire-server-td and
// spire-server-id labels are used to identify the trust domain and server.
func (p *Plugin) disposeCryptoKeysTask(ctx context.Context) { _ = "STUB: not implemented"; return }

// enqueueDestruction enqueues the specified CryptoKeyVersion for destruction.
func (p *Plugin) enqueueDestruction(cryptoKeyVersionName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// getAuthenticatedServiceAccount gets the email of the authenticated service
// account that is interacting with the Cloud KMS Service.
func (p *Plugin) getAuthenticatedServiceAccount() (email string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getConfig gets the configuration of the plugin.
func (p *Plugin) getConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// getCryptoKeyLabels gets the labels that must be set to a new CryptoKey
// that is being created.
func (p *Plugin) getCryptoKeyLabels() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getDisposeCryptoKeysFilter gets the filter to be used to get the list of
// CryptoKeys that are stale but are still marked as active.
func (p *Plugin) getDisposeCryptoKeysFilter() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getKeyEntry gets the entry from the cache that matches the provided
// SPIRE Key ID
func (p *Plugin) getKeyEntry(keyID string) (ke keyEntry, ok bool) {
	_ = "STUB: not implemented"
	return *new(keyEntry), false
}

// getPluginData gets the pluginData structure maintained by the plugin.
func (p *Plugin) getPluginData() (*pluginData, error) { _ = "STUB: not implemented"; return nil, nil }

// setIamPolicy sets the IAM policy specified in the KeyPolicyFile to the given
// resource. If there is no KeyPolicyFile specified, a default policy is constructed
// and set to the resource.
func (p *Plugin) setIamPolicy(ctx context.Context, cryptoKeyName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Get the handle to be able to inspect and change the policy of the
// CryptoKey.

// We use V3 for policies.

// Get the policy.

// We expect the policy to be empty.

// The policy is not empty, log the situation and do not replace it.

// There is a custom policy defined.

// No custom policy defined. Build the default policy.

// setKeyEntry gets the entry from the cache that matches the provided
// SPIRE Key ID
func (p *Plugin) setKeyEntry(keyID string, ke keyEntry) { _ = "STUB: not implemented"; return }

// setPluginData sets the pluginData structure maintained by the plugin.
func (p *Plugin) setPluginData(pd *pluginData) { _ = "STUB: not implemented"; return }

// keepActiveCryptoKeys keeps CryptoKeys managed by this plugin active updating
// the spire-last-update label with the current Unix time.
func (p *Plugin) keepActiveCryptoKeys(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// keepActiveCryptoKeysTask updates the CryptoKeys in the cache every 6 hours,
// setting the spire-last-update label to the current (Unix) time.
// This is done to be able to detect CryptoKeys that are inactive (not in use
// by any server).
func (p *Plugin) keepActiveCryptoKeysTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyDestroy(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyDisposeCryptoKeys(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyEnqueueDestruction(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifySetInactive(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyKeepActiveCryptoKeys(err error) { _ = "STUB: not implemented"; return }

// scheduleDestroyTask is a long-running task that schedules the destruction
// of inactive CryptoKeyVersions and sets the corresponding CryptoKey as inactive.
func (p *Plugin) scheduleDestroyTask(ctx context.Context) { _ = "STUB: not implemented"; return }

// CryptoKeyVersion is not found, no CryptoKeyVersion to destroy

// There was an error in the DestroyCryptoKeyVersion call.
// Try to get the CryptoKeyVersion to know the state of the
// CryptoKeyVersion and if we need to re-enqueue.

// Purely defensive. We don't really expect this situation,
// because this should have been captured during the
// DestroyCryptoKeyVersion call that was just performed.

// Something external to the plugin modified the state
// of the CryptoKeyVersion. Do not try to schedule it for
// destruction.

// The GetCryptoKeyVersion call failed. Log this and re-enqueue
// the CryptoKey for destruction. Hopefully, this is a
// recoverable error.

// setInactive updates the spire-active label in the specified CryptoKey to
// indicate that is inactive.
func (p *Plugin) setInactive(ctx context.Context, cryptoKey *kmspb.CryptoKey) {
	_ = "STUB: not implemented"
	return
}

// setCache sets the cached entries with the provided entries.
func (p *Plugin) setCache(keyEntries []*keyEntry) { _ = "STUB: not implemented"; return }

// createServerID creates a randomly generated UUID to be used as a server ID
// and stores it in the specified idPath.
func createServerID(idPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// cryptoKeyVersionAlgorithmFromKeyType gets the corresponding algorithm of the
// CryptoKeyVersion from the provided key type.
// The returned CryptoKeyVersion_CryptoKeyVersionAlgorithm indicates the
// parameters that must be used for signing.
func cryptoKeyVersionAlgorithmFromKeyType(keyType keymanagerv1.KeyType) (kmspb.CryptoKeyVersion_CryptoKeyVersionAlgorithm, error) {
	_ = "STUB: not implemented"
	// CryptoKeyVersion_CryptoKeyVersionAlgorithm specifies the padding algorithm
	// and the digest algorithm for RSA signatures. The key type in the Key
	// Manager interface does not contain the information about these parameters
	// for signing. Currently, there is no way in SPIRE to specify custom
	// parameters when signing through the ca.ServerCA interface and
	// x509.CreateCertificate defaults to RSASSA-PKCS-v1_5 as the padding
	// algorithm and a SHA256 digest. Therefore, for RSA signing keys we
	// choose the corresponding CryptoKeyVersion_CryptoKeyVersionAlgorithm using
	// RSASSA-PKCS-v1_5 for padding and a SHA256 digest.
	return *new(kmspb.CryptoKeyVersion_CryptoKeyVersionAlgorithm), nil
}

// generateCryptoKeyID returns a new identifier to be used as a CryptoKeyID.
// The returned identifier has the form: spire-key-<UUID>-<SPIRE-KEY-ID>,
// where UUID is a new randomly generated UUID and SPIRE-KEY-ID is provided
// through the spireKeyID parameter.
func (p *Plugin) generateCryptoKeyID(spireKeyID string) (cryptoKeyID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// crc32Checksum returns the CRC-32 checksum of data using the polynomial
// represented by the  table constructed from the specified data.
// This is used to perform integrity verification of the result when that's
// available in the Cloud Key Management Service API.
// https://cloud.google.com/kms/docs/data-integrity-guidelines
func crc32Checksum(data []byte) uint32 { _ = "STUB: not implemented"; return 0 }

// generateUniqueID returns a randomly generated UUID.
func generateUniqueID() (id string, err error) { _ = "STUB: not implemented"; return "", nil }

// getOrCreateServerID gets the server ID from the specified file path or creates
// a new server ID if the file does not exist.
func getOrCreateServerID(idPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// getPublicKeyFromCryptoKeyVersion requests Cloud KMS to get the public key
// of the specified CryptoKeyVersion.
func getPublicKeyFromCryptoKeyVersion(ctx context.Context, log hclog.Logger, kmsClient cloudKeyManagementService, cryptoKeyVersionName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the CryptoKeyVersion is still being generated or
// if it is now enabled.
// Longer generation times can be observed when using algorithms
// with large key sizes. (e.g. when rsa-4096 keys are used).
// One or two additional attempts is usually enough to find the
// CryptoKeyVersion enabled.

// This is a recoverable error.

// The CryptoKeyVersion may be ready to be used now.

// We cannot recover if it's in a different status.

// Perform integrity verification.

func makeFingerprint(pkixData []byte) string { _ = "STUB: not implemented"; return "" }

func validateCharacters(str string) bool { _ = "STUB: not implemented"; return false }

// parsePolicyFile parses a file containing iam.Policy3 data in JSON format.
func parsePolicyFile(policyFile string) (*iam.Policy3, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
