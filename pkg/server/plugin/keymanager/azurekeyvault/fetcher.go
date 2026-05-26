package azurekeyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/hashicorp/go-hclog"
	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
)

type keyFetcher struct {
	keyVaultClient cloudKeyManagementService
	log            hclog.Logger
	serverID       string
	trustDomain    string
}

// fetchKeyEntries requests Key Vault to get the list of keys that are
// active in this server. They are returned as a keyEntry array.
func (kf *keyFetcher) fetchKeyEntries(ctx context.Context) ([]*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all the key from the configured key vault URL

// Skip keys that do not belong this server

// trigger a goroutine to get the details of the key

// Wait for all the detail gathering routines to finish.

func (kf *keyFetcher) keyBelongsToServer(key *azkeys.KeyProperties) bool {
	_ = "STUB: not implemented"
	return false
}

func (kf *keyFetcher) fetchKeyEntryDetails(ctx context.Context, keyProperties *azkeys.KeyProperties, spireKeyID string) (*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this means something external to the plugin, disabled the key
// returning an error provides the opportunity of reverting this in azure key vault

func keyTypeFromKeySpec(keyBundle azkeys.KeyBundle) (keymanagerv1.KeyType, bool) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), false
}

// spireKeyIDFromKeyName parses a Key Vault key name to get the
// SPIRE Key ID. This Key ID is used in the Server KeyManager interface.
func spireKeyIDFromKeyName(keyName string) (string, bool) {
	_ = "STUB: not implemented"
	// A key name would have the format spire-key-${UUID}-x509-CA-A.
	// first we find the position where the SPIRE Key ID starts.
	// For that, we need to add the length of the key name prefix that we
	// are using, the UUID length, and the two "-" separators used in our format.
	return "", false
}

// 39 is the UUID length plus two '-' separators

// The index is out of range.
