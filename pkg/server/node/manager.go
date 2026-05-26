package node

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/datastore"
)

const (
	defaultJobInterval = time.Hour
	maxJitter          = 15 * time.Minute
)

type PruneArgs struct {
	ExpiredFor             time.Duration
	IncludeNonReattestable bool
}

type ManagerConfig struct {
	DataStore datastore.DataStore

	Log     logrus.FieldLogger
	Metrics telemetry.Metrics

	Clock    clock.Clock
	Interval time.Duration

	PruneArgs
}

type Manager struct {
	c       ManagerConfig
	log     logrus.FieldLogger
	metrics telemetry.Metrics

	pruneRequestedCh chan PruneArgs
}

func NewManager(c ManagerConfig) *Manager { _ = "STUB: not implemented"; return nil }

// Add random jitter: ±15 minutes (45-75 minutes range)
//nolint // gosec: no need for cryptographic randomness here

func (m *Manager) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) Prune(ctx context.Context, expiredFor time.Duration, includeNonReattestable bool) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) pruneEvery(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) prune(ctx context.Context, expiredBefore time.Time, includeNonReattestable bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
