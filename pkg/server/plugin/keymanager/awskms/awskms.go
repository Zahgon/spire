package awskms

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/hashicorp/go-hclog"
	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName  = "aws_kms"
	aliasPrefix = "alias/SPIRE_SERVER/"

	keyArnTag    = "key_arn"
	aliasNameTag = "alias_name"
	reasonTag    = "reason"

	refreshAliasesFrequency = time.Hour * 6
	disposeAliasesFrequency = time.Hour * 24
	aliasThreshold          = time.Hour * 24 * 14 // two weeks

	disposeKeysFrequency = time.Hour * 48
	keyThreshold         = time.Hour * 48
)

var (
	validTagKeyPattern   = regexp.MustCompile(`^[\p{L}\p{N}\s+\-=._:/@]+$`)
	validTagValuePattern = regexp.MustCompile(`^[\p{L}\p{N}\s+\-=._:/@]*$`)
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type keyEntry struct {
	Arn       string
	AliasName string
	PublicKey *keymanagerv1.PublicKey
}

type pluginHooks struct {
	newKMSClient func(aws.Config) (kmsClient, error)
	newSTSClient func(aws.Config) (stsClient, error)
	clk          clock.Clock
	// just for testing
	scheduleDeleteSignal chan error
	refreshAliasesSignal chan error
	disposeAliasesSignal chan error
	disposeKeysSignal    chan error
}

// Plugin is the main representation of this keymanager plugin
type Plugin struct {
	keymanagerv1.UnsafeKeyManagerServer
	configv1.UnsafeConfigServer

	log            hclog.Logger
	mu             sync.RWMutex
	entries        map[string]keyEntry
	kmsClient      kmsClient
	stsClient      stsClient
	trustDomain    string
	serverID       string
	scheduleDelete chan string
	cancelTasks    context.CancelFunc
	hooks          pluginHooks
	keyPolicy      *string
	keyTags        []types.Tag
}

// Config provides configuration context for the plugin
type Config struct {
	AccessKeyID        string            `hcl:"access_key_id" json:"access_key_id"`
	SecretAccessKey    string            `hcl:"secret_access_key" json:"secret_access_key"`
	Region             string            `hcl:"region" json:"region"`
	KeyIdentifierFile  string            `hcl:"key_identifier_file" json:"key_identifier_file"`
	KeyIdentifierValue string            `hcl:"key_identifier_value" json:"key_identifier_value"`
	KeyPolicyFile      string            `hcl:"key_policy_file" json:"key_policy_file"`
	KeyTags            map[string]string `hcl:"key_tags" json:"key_tags"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

// New returns an instantiated plugin
func New() *Plugin { _ = "STUB: not implemented"; return nil }

func newPlugin(
	newKMSClient func(aws.Config) (kmsClient, error),
	newSTSClient func(aws.Config) (stsClient, error),
) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

// SetLogger sets a logger
func (p *Plugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Configure sets up the plugin
	return
}

func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cancels previous tasks in case of re-configure

// start tasks

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateKey creates a key in KMS. If a key already exists in the local storage, it is updated.
func (p *Plugin) GenerateKey(ctx context.Context, req *keymanagerv1.GenerateKeyRequest) (*keymanagerv1.GenerateKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignData creates a digital signature for the data to be signed
func (p *Plugin) SignData(ctx context.Context, req *keymanagerv1.SignDataRequest) (*keymanagerv1.SignDataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKey returns the public key for a given key
func (p *Plugin) GetPublicKey(_ context.Context, req *keymanagerv1.GetPublicKeyRequest) (*keymanagerv1.GetPublicKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKeys return the publicKey for all the keys
func (p *Plugin) GetPublicKeys(context.Context, *keymanagerv1.GetPublicKeysRequest) (*keymanagerv1.GetPublicKeysResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) createKey(ctx context.Context, spireKeyID string, keyType keymanagerv1.KeyType) (*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) assignAlias(ctx context.Context, entry *keyEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// create alias

// update alias

func (p *Plugin) setCache(keyEntries []*keyEntry) {
	_ = "STUB: not implemented"
	// clean previous cache
	return
}

// add results to cache

// scheduleDeleteTask ia a long-running task that deletes keys that were rotated
func (p *Plugin) scheduleDeleteTask(ctx context.Context) { _ = "STUB: not implemented"; return }

// refreshAliasesTask will update the alias of all keys in the cache every 6 hours.
// Aliases will be updated to the same key they already have.
// The consequence of this is that the field LastUpdatedDate in each alias belonging to the server will be set to the current date.
// This is all with the goal of being able to detect keys that are not in use by any server.
func (p *Plugin) refreshAliasesTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *Plugin) refreshAliases(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Resolve the alias's current target to avoid reverting a rotation
// performed by another replica sharing the same key_identifier_value.

// disposeAliasesTask will be run every 24hs.
// It will delete aliases that have a LastUpdatedDate value older than two weeks.
// It will also delete the keys associated with them.
// It will only delete aliases belonging to the current trust domain but not the current server.
// disposeAliasesTask relies on how aliases are built with prefixes to do all this.
// Alias example: `alias/SPIRE_SERVER/{TRUST_DOMAIN}/{SERVER_ID}/{KEY_ID}`
func (p *Plugin) disposeAliasesTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *Plugin) disposeAliases(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// if alias does not belong to trust domain skip

// if alias belongs to current server skip

// disposeKeysTask will be run every 48hs.
// It will delete keys that have a CreationDate value older than 48hs.
// It will only delete keys belonging to the current trust domain and without an alias.
// disposeKeysTask relies on how the keys description is built to do all this.
// Key description example: `SPIRE_SERVER/{TRUST_DOMAIN}`
// Keys belonging to a server should never be without an alias.
// The goal of this task is to remove keys that ended in this invalid state during a failure on alias assignment.
func (p *Plugin) disposeKeysTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *Plugin) disposeKeys(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// if key does not belong to trust domain, skip it

// if key has alias, skip it

func (p *Plugin) aliasFromSpireKeyID(spireKeyID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Plugin) descriptionFromSpireKeyID(spireKeyID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Plugin) descriptionPrefixForTrustDomain() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) aliasPrefixForServer() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) aliasPrefixForTrustDomain() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) notifyDelete(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyRefreshAliases(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyDisposeAliases(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) notifyDisposeKeys(err error) { _ = "STUB: not implemented"; return }

func (p *Plugin) createDefaultPolicy(ctx context.Context) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the server has not assumed any role, use default KMS policy and log a warn message

// roleNameFromARN returns the role name included in an ARN. If no role name exist
// an error is returned.
// ARN example: "arn:aws:sts::123456789:assumed-role/the-role-name/i-0001f4f25acfd1234",
func roleNameFromARN(arn string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func sanitizeTrustDomain(trustDomain string) string { _ = "STUB: not implemented"; return "" }

func signingAlgorithmForKMS(keyType keymanagerv1.KeyType, signerOpts any) (types.SigningAlgorithmSpec, error) {
	_ = "STUB: not implemented"
	return *new(types.SigningAlgorithmSpec), nil
}

// opts.PssOptions.SaltLength is handled by KMS. The salt length matches the bits of the hashing algorithm.

func keyTypeFromKeySpec(keySpec types.KeySpec) (keymanagerv1.KeyType, bool) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), false
}

func keySpecFromKeyType(keyType keymanagerv1.KeyType) (types.KeySpec, bool) {
	_ = "STUB: not implemented"
	return *new(types.KeySpec), false
}

func getOrCreateServerID(idPath string) (string, error) {
	_ = "STUB: not implemented"
	// get id from path
	return "", nil
}

// validate what we got is a uuid

func createServerID(idPath string) (string, error) {
	_ = "STUB: not implemented"
	// generate id
	return "", nil
}

// persist id

func makeFingerprint(pkixData []byte) string { _ = "STUB: not implemented"; return "" }

func validateTags(tags map[string]string) error { _ = "STUB: not implemented"; return nil }

func buildKeyTags(tags map[string]string) []types.Tag { _ = "STUB: not implemented"; return nil }

// encodeKeyID maps "." and "+" characters to the asciihex value using "_" as
// escape character. Currently, KMS does not support those characters to be used
// as alias name.
func encodeKeyID(keyID string) string { _ = "STUB: not implemented"; return "" }

// decodeKeyID decodes "." and "+" from the asciihex value using "_" as
// escape character.
func decodeKeyID(keyID string) string { _ = "STUB: not implemented"; return "" }
