package keymanager

import (
	"context"

	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

func WithMetrics(km keymanager.KeyManager, metrics telemetry.Metrics) keymanager.KeyManager {
	_ = "STUB: not implemented"
	return *new(keymanager.KeyManager)
}

type keyManagerWrapper struct {
	catalog.PluginInfo
	km keymanager.KeyManager
	m  telemetry.Metrics
}

func (w keyManagerWrapper) GenerateKey(ctx context.Context, keyID string, keyType keymanager.KeyType) (_ keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), nil
}

func (w keyManagerWrapper) GetKey(ctx context.Context, keyID string) (_ keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), nil
}

func (w keyManagerWrapper) GetKeys(ctx context.Context) (_ []keymanager.Key, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
