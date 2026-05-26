package keymanager

import (
	"context"
)

// SVIDKeyManager is a wrapper around the key manager specifically used for
// managing the agent SVID.
type SVIDKeyManager interface {
	// GenerateKey generates a new key. The current key is passed, if available
	// so the key manager can determine which which "slot" to occupy (i.e.
	// which key ID to use for the new key).
	GenerateKey(ctx context.Context, currentKey Key) (Key, error)

	// GetKeys returns all keys managed by the KeyManager.
	GetKeys(ctx context.Context) ([]Key, error)
}

// Returns an SVIDKeyManager over the given KeyManager
func ForSVID(km KeyManager) SVIDKeyManager { _ = "STUB: not implemented"; return *new(SVIDKeyManager) }

type svidKeyManager struct {
	km KeyManager
}

func (s svidKeyManager) GenerateKey(ctx context.Context, currentKey Key) (Key, error) {
	_ = "STUB: not implemented"
	return *new(Key), nil
}

func (s svidKeyManager) GetKeys(ctx context.Context) ([]Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
