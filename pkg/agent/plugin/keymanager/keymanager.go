package keymanager

import (
	"context"
	"crypto"

	"github.com/spiffe/spire/pkg/common/catalog"
)

// KeyManager provides either a single or multi-key key manager
type KeyManager interface {
	catalog.PluginInfo

	// GenerateKey generates a key with the given ID and key type. If a key
	// with that ID already exists, it is overwritten.
	GenerateKey(ctx context.Context, id string, keyType KeyType) (Key, error)

	// GetKey returns the key with the given ID. If a key with that ID does
	// not exist, a status of codes.NotFound is returned.
	GetKey(ctx context.Context, id string) (Key, error)

	// GetKeys returns all keys managed by the KeyManager.
	GetKeys(ctx context.Context) ([]Key, error)
}

// Key is a KeyManager-backed key
type Key interface {
	crypto.Signer

	// ID returns the ID of the key in the KeyManager.
	ID() string
}

// KeyType represents the types of keys that are supported by the KeyManager.
type KeyType int

const (
	KeyTypeUnset KeyType = iota
	ECP256
	ECP384
	RSA2048
	RSA4096
)

// GenerateSigner generates a new key for the given key type
func (keyType KeyType) GenerateSigner() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// String returns the string representation of the key type
func (keyType KeyType) String() string { _ = "STUB: not implemented"; return "" }
