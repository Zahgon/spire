package rotator

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/server/ca/manager"
)

const (
	rotateInterval          = 10 * time.Second
	pruneBundleInterval     = 6 * time.Hour
	pruneCAJournalsInterval = 8 * time.Hour
)

type CAManager interface {
	NotifyBundleLoaded(ctx context.Context) error
	ProcessBundleUpdates(ctx context.Context)

	GetCurrentX509CASlot() manager.Slot
	GetNextX509CASlot() manager.Slot

	PrepareX509CA(ctx context.Context) error
	ActivateX509CA(ctx context.Context)
	RotateX509CA(ctx context.Context)

	GetCurrentJWTKeySlot() manager.Slot
	GetNextJWTKeySlot() manager.Slot

	PrepareJWTKey(ctx context.Context) error
	ActivateJWTKey(ctx context.Context)
	RotateJWTKey(ctx context.Context)

	GetCurrentWITKeySlot() manager.Slot
	GetNextWITKeySlot() manager.Slot

	PrepareWITKey(ctx context.Context) error
	ActivateWITKey(ctx context.Context)
	RotateWITKey(ctx context.Context)

	SubscribeToLocalBundle(ctx context.Context) error

	PruneBundle(ctx context.Context) error
	PruneCAJournals(ctx context.Context) error
}

type Config struct {
	Manager       CAManager
	Log           logrus.FieldLogger
	Clock         clock.Clock
	HealthChecker health.Checker
}

type Rotator struct {
	c Config

	// For keeping track of number of failed rotations.
	failedRotationNum uint64
}

func NewRotator(c Config) *Rotator { _ = "STUB: not implemented"; return nil }

func (r *Rotator) Initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Rotator) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// notifyOnBundleUpdate does not fail but rather logs any errors
// encountered while notifying

func (r *Rotator) rotateEvery(ctx context.Context, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// rotate() errors are logged by rotate() and shouldn't cause the
// manager run task to bail so ignore them here. The error returned
// by rotate is used by the unit tests, so we need to keep it for
// now.

func (r *Rotator) rotate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Preparation of the X509 CA failed, and there is no active X509
// authority. We will be unable to store the JWT authority, so we
// don't try to rotate the JWT key in this case.

func (r *Rotator) rotateJWTKey(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// if there is no current keypair set, generate one

// if there is no next keypair set and the current is within the
// preparation threshold, generate one.

func (r *Rotator) rotateWITKey(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// if there is no current keypair set, generate one

// if there is no next keypair set and the current is within the
// preparation threshold, generate one.

func (r *Rotator) rotateX509CA(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// if there is no current keypair set, generate one

// if there is no next keypair set and the current is within the
// preparation threshold, generate one.

func (r *Rotator) pruneBundleEvery(ctx context.Context, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rotator) pruneCAJournalsEvery(ctx context.Context, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rotator) failedRotationResult() uint64 { _ = "STUB: not implemented"; return 0 }
