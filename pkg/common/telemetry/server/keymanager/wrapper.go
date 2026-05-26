package keymanager

import (
	"context"
	"crypto"
	"io"

	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
)

func WithMetrics(km keymanager.KeyManager, metrics telemetry.Metrics) keymanager.KeyManager {
	_ = "STUB: not implemented"
	return *new(keymanager.KeyManager)
}

type keyManagerWrapper struct {
	keymanager.KeyManager
	m telemetry.Metrics
}

func (w keyManagerWrapper) GenerateKey(ctx context.Context, id string, keyType keymanager.KeyType) (_ keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), nil
}

func (w keyManagerWrapper) GetKey(ctx context.Context, id string) (_ keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), nil
}

func (w keyManagerWrapper) GetKeys(ctx context.Context) (_ []keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keyWrapper struct {
	keymanager.Key
	m telemetry.Metrics
}

func (w keyWrapper) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (_ []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapKeys(m telemetry.Metrics, keys []keymanager.Key) []keymanager.Key {
	_ = "STUB: not implemented"
	return nil
}

func wrapKey(m telemetry.Metrics, key keymanager.Key) keymanager.Key {
	_ = "STUB: not implemented"
	return *new(keymanager.Key)
}
