package health

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
)

type checkState struct {
	// err is the error returned from a failed health check
	err error

	// details contains more contextual detail about a
	// failing health check.
	details State

	// checkTime is the time of the last health check
	checkTime time.Time

	// contiguousFailures the number of failures that occurred in a row
	contiguousFailures int64

	// timeOfFirstFailure the time of the initial transitional failure for
	// any given health check
	timeOfFirstFailure time.Time
}

type checkerSubsystem struct {
	state     checkState
	checkable Checkable
}

func newCache(log logrus.FieldLogger, clock clock.Clock) *cache {
	_ = "STUB: not implemented"
	return nil
}

type cache struct {
	checkerSubsystems map[string]*checkerSubsystem

	mtx sync.RWMutex
	clk clock.Clock

	log   logrus.FieldLogger
	hooks struct {
		statusUpdated chan struct{}
	}
	startupComplete chan struct{}
}

func (c *cache) addCheck(name string, checkable Checkable) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache) getCheckerSubsystems() map[string]*checkerSubsystem {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache) getStatuses() map[string]checkState { _ = "STUB: not implemented"; return nil }

func (c *cache) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *cache) startRunner(ctx context.Context) { _ = "STUB: not implemented"; return }

// Run health check in a tighter loop until we get an initial ready + live state

// Wait until initial ready + live state is achieved, then periodically check health at a longer interval

func (c *cache) setStatus(name string, prevState checkState, state checkState) {
	_ = "STUB: not implemented"
	return
}

// We are sure that checker exists in this place, to be able to check
// status of a subsystem we must call the checker inside this map

func (c *cache) embellishState(name string, prevState, state *checkState) {
	_ = "STUB: not implemented"
	return
}

// All fine continue

// State start to fail, add log and set failures tracking

// Error still happening, carry the time of first failure from the previous state

// Current state has no error, notify about error recovering

func verifyStatus(check Checkable) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}
