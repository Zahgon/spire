package gcpkms

import (
	"context"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/hashicorp/go-hclog"
	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
)

type keyFetcher struct {
	keyRing   string
	kmsClient cloudKeyManagementService
	log       hclog.Logger
	serverID  string
	tdHash    string
}

// fetchKeyEntries requests Cloud KMS to get the list of CryptoKeys that are
// active in this server. They are returned as a keyEntry array.
func (kf *keyFetcher) fetchKeyEntries(ctx context.Context) ([]*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Trigger a goroutine to get the details of the key

// Wait for all the detail gathering routines to finish.

// getKeyEntriesFromCryptoKey builds an array of keyEntry values from the provided
// CryptoKey. In order to do that, Cloud KMS is requested to list the
// CryptoKeyVersions of the CryptoKey. The public key of the CryptoKeyVersion is
// also retrieved from each CryptoKey to construct each keyEntry.
func (kf *keyFetcher) getKeyEntriesFromCryptoKey(ctx context.Context, cryptoKey *kmspb.CryptoKey, spireKeyID string) (keyEntries []*keyEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter by state, so only enabled keys are returned. This will leave
// out all the versions that have been rotated.

// getSPIREKeyIDFromCryptoKeyName parses a CryptoKey resource name to get the
// SPIRE Key ID. This Key ID is used in the Server KeyManager interface.
func getSPIREKeyIDFromCryptoKeyName(cryptoKeyName string) (string, bool) {
	_ = "STUB: not implemented"
	// cryptoKeyName is the resource name for the CryptoKey holding the SPIRE Key
	// in the format: projects/*/locations/*/keyRings/*/cryptoKeys/spire-key-*-*.
	// Example: projects/project-name/locations/us-east1/keyRings/key-ring-name/cryptoKeys/spire-key-1f2e225a-91d8-4589-a4fe-f88b7bb04bac-x509-CA-A
	return "", false
}

// Get the last element of the path.

// All CryptoKeys are under a Key Ring; not a valid Crypto Key name.

// The i index will indicate us where
// "spire-key-1f2e225a-91d8-4589-a4fe-f88b7bb04bac-x509-CA-A" starts.
// Now we have to get the position where the SPIRE Key ID starts.
// For that, we need to add the length of the CryptoKey name prefix that we
// are using, the UUID length, and the two "-" separators used in our format.
// 39 is the UUID length plus two '-' separators

// The index is out of range.

// keyTypeFromCryptoKeyVersionAlgorithm gets the KeyType that corresponds to the
// given CryptoKeyVersion_CryptoKeyVersionAlgorithm.
func keyTypeFromCryptoKeyVersionAlgorithm(algorithm kmspb.CryptoKeyVersion_CryptoKeyVersionAlgorithm) (keymanagerv1.KeyType, bool) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), false
}
