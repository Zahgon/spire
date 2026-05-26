package keymanagerbase

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"sync"

	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
)

// KeyEntry is an entry maintained by the key manager
type KeyEntry struct {
	PrivateKey crypto.Signer
	*keymanagerv1.PublicKey
}

// Config is a collection of optional callbacks. Default implementations will be
// used when not provided.
type Config struct {
	// Generator is an optional key generator.
	Generator Generator

	// WriteEntries is an optional callback used to persist key entries
	WriteEntries func(ctx context.Context, entries []*KeyEntry) error
}

// Generator is a key generator
type Generator interface {
	GenerateRSA2048Key() (crypto.Signer, error)
	GenerateRSA4096Key() (crypto.Signer, error)
	GenerateEC256Key() (crypto.Signer, error)
	GenerateEC384Key() (crypto.Signer, error)
}

// Base is the base KeyManager implementation
type Base struct {
	keymanagerv1.UnsafeKeyManagerServer
	config Config

	mu      sync.RWMutex
	entries map[string]*KeyEntry
}

// New creates a new base key manager using the provided config.
func New(config Config) *Base { _ = "STUB: not implemented"; return nil }

// SetEntries is used to replace the set of managed entries. This is generally
// called by implementations when they are first loaded to set the initial set
// of entries.
func (m *Base) SetEntries(entries []*KeyEntry) { _ = "STUB: not implemented"; return }

// populate the fingerprints

// GenerateKey implements the KeyManager RPC of the same name.
func (m *Base) GenerateKey(ctx context.Context, req *keymanagerv1.GenerateKeyRequest) (*keymanagerv1.GenerateKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKey implements the KeyManager RPC of the same name.
func (m *Base) GetPublicKey(_ context.Context, req *keymanagerv1.GetPublicKeyRequest) (*keymanagerv1.GetPublicKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKeys implements the KeyManager RPC of the same name.
func (m *Base) GetPublicKeys(context.Context, *keymanagerv1.GetPublicKeysRequest) (*keymanagerv1.GetPublicKeysResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignData implements the KeyManager RPC of the same name.
func (m *Base) SignData(_ context.Context, req *keymanagerv1.SignDataRequest) (*keymanagerv1.SignDataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Base) generateKey(ctx context.Context, req *keymanagerv1.GenerateKeyRequest) (*keymanagerv1.GenerateKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Base) signData(req *keymanagerv1.SignDataRequest) (*keymanagerv1.SignDataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Base) getPrivateKeyAndFingerprint(id string) (crypto.Signer, string, bool) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), "", false
}

func (m *Base) generateKeyEntry(keyID string, keyType keymanagerv1.KeyType) (e *KeyEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeKeyEntry(keyID string, keyType keymanagerv1.KeyType, privateKey crypto.Signer) (*KeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeKeyEntryFromKey(id string, privateKey crypto.PrivateKey) (*KeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rsaKeyType(privateKey *rsa.PrivateKey) (keymanagerv1.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), nil
}

func ecdsaKeyType(privateKey *ecdsa.PrivateKey) (keymanagerv1.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), nil
}

type defaultGenerator struct{}

func (defaultGenerator) GenerateRSA2048Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func (defaultGenerator) GenerateRSA4096Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func (defaultGenerator) GenerateEC256Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func (defaultGenerator) GenerateEC384Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func entriesSliceFromMap(entriesMap map[string]*KeyEntry) (entriesSlice []*KeyEntry) {
	_ = "STUB: not implemented"
	return nil
}

func entriesMapFromSlice(entriesSlice []*KeyEntry) map[string]*KeyEntry {
	_ = "STUB: not implemented"
	// return keys in sorted order for consistency
	return nil
}

func clonePublicKey(publicKey *keymanagerv1.PublicKey) *keymanagerv1.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func makeFingerprint(pkixData []byte) string { _ = "STUB: not implemented"; return "" }

func SortKeyEntries(entries []*KeyEntry) { _ = "STUB: not implemented"; return }

func prefixStatus(err error, prefix string) error { _ = "STUB: not implemented"; return nil }
