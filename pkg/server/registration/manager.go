package registration

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/datastore"
)

const (
	_pruningCadence = 5 * time.Minute
)

// ManagerConfig is the config for the registration manager
type ManagerConfig struct {
	DataStore datastore.DataStore

	Log     logrus.FieldLogger
	Metrics telemetry.Metrics

	Clock clock.Clock
}

// Manager is the manager of registrations
type Manager struct {
	c       ManagerConfig
	log     logrus.FieldLogger
	metrics telemetry.Metrics
}

// NewManager creates a new registration manager
func NewManager(c ManagerConfig) *Manager { _ = "STUB: not implemented"; return nil }

// Run runs the registration manager
func (m *Manager) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) pruneEvery(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Log an error on failure unless we're shutting down

func (m *Manager) prune(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }
