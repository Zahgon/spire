package workloadkey

import (
	"crypto"
)

func KeyTypeFromString(s string) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

// KeyType represents the types of keys that are supported by the KeyManager.
type KeyType int

const (
	KeyTypeUnset KeyType = iota
	ECP256
	RSA2048
	ECP384
)

// GenerateSigner generates a new key for the given key type
func (keyType KeyType) GenerateSigner() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// String returns the string representation of the key type
func (keyType KeyType) String() string { _ = "STUB: not implemented"; return "" }
