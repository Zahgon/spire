package keymanager

import (
	"context"
	"crypto"
	"io"

	keymanagerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/keymanager/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
)

type V1 struct {
	plugin.Facade

	keymanagerv1.KeyManagerPluginClient
}

func (v1 V1) GenerateKey(ctx context.Context, id string, keyType KeyType) (Key, error) {
	_ = "STUB: not implemented"
	return *new(Key), nil
}

func (v1 V1) GetKey(ctx context.Context, id string) (Key, error) {
	_ = "STUB: not implemented"
	return *new(Key), nil
}

func (v1 V1) GetKeys(ctx context.Context) ([]Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v1 V1) makeKey(id string, pb *keymanagerv1.PublicKey) (Key, error) {
	_ = "STUB: not implemented"
	return *new(Key), nil
}

func (v1 *V1) convertKeyType(t KeyType) (keymanagerv1.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(keymanagerv1.KeyType), nil
}

func (v1 *V1) convertHashAlgorithm(h crypto.Hash) keymanagerv1.HashAlgorithm {
	_ = "STUB: not implemented"
	// Hash algorithm constants are aligned.
	return *new(keymanagerv1.HashAlgorithm)
}

type v1Key struct {
	v1          V1
	id          string
	fingerprint string
	publicKey   crypto.PublicKey
}

func (s *v1Key) ID() string { _ = "STUB: not implemented"; return "" }

func (s *v1Key) Public() crypto.PublicKey { _ = "STUB: not implemented"; return *new(crypto.PublicKey) }

func (s *v1Key) Sign(_ io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	// rand is purposefully ignored since it can't be communicated between
	// the plugin boundary. The crypto.Signer interface implies this is ok
	// when it says "possibly using entropy from rand".
	return nil, nil
}

func (s *v1Key) signContext(ctx context.Context, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
