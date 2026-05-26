package keymanagertest

import (
	"context"
	"crypto/elliptic"
	"crypto/x509"
	"testing"

	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
	keymanagerbase "github.com/spiffe/spire/pkg/server/plugin/keymanager/base"
)

type keyAlgorithm int

const (
	keyAlgorithmEC keyAlgorithm = iota
	keyAlgorithmRSA
)

var (
	ctx = context.Background()

	keyTypes = map[keymanager.KeyType]keyAlgorithm{
		keymanager.ECP256:  keyAlgorithmEC,
		keymanager.ECP384:  keyAlgorithmEC,
		keymanager.RSA2048: keyAlgorithmRSA,
		keymanager.RSA4096: keyAlgorithmRSA,
	}

	expectCurve = map[keymanager.KeyType]elliptic.Curve{
		keymanager.ECP256: elliptic.P256(),
		keymanager.ECP384: elliptic.P384(),
	}

	expectBits = map[keymanager.KeyType]int{
		keymanager.RSA2048: 2048,
		keymanager.RSA4096: 4096,
	}
)

func NewGenerator() keymanagerbase.Generator {
	_ = "STUB: not implemented"
	return *new(keymanagerbase.Generator)
}

type CreateFunc = func(t *testing.T) keymanager.KeyManager

type Config struct {
	Create CreateFunc

	// UnsupportedSignatureAlgorithms is a map of algorithms that are
	// unsupported for the given key type.
	UnsupportedSignatureAlgorithms map[keymanager.KeyType][]x509.SignatureAlgorithm

	signatureAlgorithms map[keymanager.KeyType][]x509.SignatureAlgorithm
}

func (config *Config) testKey(t *testing.T, key keymanager.Key, keyType keymanager.KeyType) {
	_ = "STUB: not implemented"
	return
}

func (config *Config) testKeyWithID(t *testing.T, key keymanager.Key, keyType keymanager.KeyType, expectID string) {
	_ = "STUB: not implemented"
	return
}

func Test(t *testing.T, config Config) {
	_ = "STUB: not implemented"
	// Build a convenient set to look up unsupported algorithms
	return
}

// build up the list of key types and hash algorithms to test

func testGenerateKey(t *testing.T, config Config) { _ = "STUB: not implemented"; return }

// Signing with oldKey should fail since it has been overwritten.

func testGetKey(t *testing.T, config Config) { _ = "STUB: not implemented"; return }

func testGetKeys(t *testing.T, config Config) { _ = "STUB: not implemented"; return }

func requireGenerateKey(t *testing.T, km keymanager.KeyManager, keyType keymanager.KeyType) keymanager.Key {
	_ = "STUB: not implemented"
	return *new(keymanager.Key)
}

func requireGenerateKeyWithID(t *testing.T, km keymanager.KeyManager, keyType keymanager.KeyType, id string) keymanager.Key {
	_ = "STUB: not implemented"
	return *new(keymanager.Key)
}

func requireGetKey(t *testing.T, km keymanager.KeyManager, id string) keymanager.Key {
	_ = "STUB: not implemented"
	return *new(keymanager.Key)
}

func requireGetKeys(t *testing.T, km keymanager.KeyManager) []keymanager.Key {
	_ = "STUB: not implemented"
	return nil
}

func assertECKey(t *testing.T, key keymanager.Key, curve elliptic.Curve) {
	_ = "STUB: not implemented"
	return
}

func assertRSAKey(t *testing.T, key keymanager.Key, bits int) { _ = "STUB: not implemented"; return }

func testSignCertificates(t *testing.T, key keymanager.Key, signatureAlgorithms []x509.SignatureAlgorithm) {
	_ = "STUB: not implemented"
	return
}

func assertSignCertificate(t *testing.T, key keymanager.Key, signatureAlgorithm x509.SignatureAlgorithm) {
	_ = "STUB: not implemented"
	return
}
