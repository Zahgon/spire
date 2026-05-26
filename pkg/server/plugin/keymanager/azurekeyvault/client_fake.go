package azurekeyvault

import (
	"context"
	"crypto"
	"sync"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/andres-erbsen/clock"
)

type kmsClientFake struct {
	t               *testing.T
	store           fakeStore
	vaultURI        string
	trustDomain     string
	serverID        string
	mu              sync.RWMutex
	createKeyErr    error
	deleteKeyErr    error
	updateKeyErr    error
	getKeyErr       error
	listKeysErr     error
	getPublicKeyErr error
	signErr         error
}

type fakeStore struct {
	fakeKeys   map[string]*fakeKeyEntry
	ec256Key   crypto.Signer
	ec384Key   crypto.Signer
	rsa2048Key crypto.Signer
	rsa4096Key crypto.Signer
	mu         sync.RWMutex
	clk        *clock.Mock
}

type fakeKeyEntry struct {
	KeyBundle  azkeys.KeyBundle
	PrivateKey crypto.Signer
}

func newKMSClientFake(t *testing.T, vaultURI, trustDomain, serverID string, c *clock.Mock) *kmsClientFake {
	_ = "STUB: not implemented"
	return nil
}

func newFakeStore(c *clock.Mock, t *testing.T) fakeStore {
	_ = "STUB: not implemented"
	return *new(fakeStore)
}

func (fs *fakeStore) SaveKeyEntry(input *fakeKeyEntry) { _ = "STUB: not implemented"; return }

func (fs *fakeStore) DeleteKeyEntry(keyName string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setEntries(entries []fakeKeyEntry) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setCreateKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setGetKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setGetPublicKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setUpdateKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setDeleteKeyErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setSignDataErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setListKeysErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) CreateKey(_ context.Context, keyName string, parameters azkeys.CreateKeyParameters, _ *azkeys.CreateKeyOptions) (azkeys.CreateKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.CreateKeyResponse), nil
}

func (k *kmsClientFake) DeleteKey(_ context.Context, name string, _ *azkeys.DeleteKeyOptions) (azkeys.DeleteKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.DeleteKeyResponse), nil
}

func (k *kmsClientFake) UpdateKey(_ context.Context, name, _ string, _ azkeys.UpdateKeyParameters, _ *azkeys.UpdateKeyOptions) (azkeys.UpdateKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.UpdateKeyResponse), nil
}

func (k *kmsClientFake) GetKey(_ context.Context, keyName, _ string, _ *azkeys.GetKeyOptions) (azkeys.GetKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.GetKeyResponse), nil
}

func (k *kmsClientFake) NewListKeyPropertiesPager(_ *azkeys.ListKeyPropertiesOptions) *runtime.Pager[azkeys.ListKeyPropertiesResponse] {
	_ = "STUB: not implemented"
	return nil
}

func (k *kmsClientFake) Sign(_ context.Context, keyName, _ string, parameters azkeys.SignParameters, _ *azkeys.SignOptions) (azkeys.SignResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.SignResponse), nil
}

// This is to produce an IEEE-P1363 encoded signature since that's how the azure signature is encoded

func toRSAKey(publicKey crypto.PublicKey, kmsKeyID string, keyOperations []*azkeys.KeyOperation) *azkeys.JSONWebKey {
	_ = "STUB: not implemented"
	return nil
}

func toECKey(publicKey crypto.PublicKey, keyName string, curveName azkeys.CurveName, keyOperations []*azkeys.KeyOperation) *azkeys.JSONWebKey {
	_ = "STUB: not implemented"
	return nil
}

// ecdsa.PublicKey.Bytes returns an uncompressed point encoded per SEC 1
// v2.0 Section 2.3.3: 0x04 || X || Y.
// Here, X and Y are encoded as defined in Section 2.3.5 of SEC 1.
// JWK EC "x" and "y" are defined by RFC 7518 to use that same Section
// 2.3.5 encoding.
// Therefore, accessing encodedPoint sub-slices below is correct.
// Note that accessing ecdsaKey.X and ecdsaKey.Y would also work, but is
// deprecated.

func (fs *fakeStore) FetchKeyEntry(keyName string) (*fakeKeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *fakeStore) fetchKeyEntry(keyName string) (*fakeKeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *fakeStore) fetchKeyEntries() []fakeKeyEntry { _ = "STUB: not implemented"; return nil }

func getKeyOperations() []*azkeys.KeyOperation { _ = "STUB: not implemented"; return nil }
