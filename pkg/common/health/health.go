package health

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	readyCheckInitialInterval = time.Second
	readyCheckInterval        = time.Minute
)

// State is the health state of a subsystem.
type State struct {
	// Started is whether the subsystem is finished starting.
	// if undefined, it is treated as started=true.
	Started *bool

	// Live is whether the subsystem is live (i.e. in a good state
	// or in a state it can recover from while remaining alive). Global
	// liveness is only reported true if all subsystems report live.
	Live bool

	// Ready is whether the subsystem is ready (i.e. ready to perform
	// its function). Global readiness is only reported true if all subsystems
	// report ready.
	Ready bool

	// Subsystems can return whatever details they want here as long as it is
	// serializable via json.Marshal.
	// LiveDetails are opaque details related to the live check.
	LiveDetails any

	// ReadyDetails are opaque details related to the live check.
	ReadyDetails any
}

// Checkable is the interface implemented by subsystems that the checker uses
// to determine subsystem health.
type Checkable interface {
	CheckHealth() State
}

// Checker is responsible for running health checks and serving the healthcheck HTTP paths
type Checker interface {
	AddCheck(name string, checkable Checkable) error
}

type ServableChecker interface {
	Checker
	ListenAndServe(ctx context.Context) error
}

func NewChecker(config Config, log logrus.FieldLogger) ServableChecker {
	_ = "STUB: not implemented"
	return *new(ServableChecker)
}

// Start HTTP server if ListenerEnabled is true

type checker struct {
	config Config

	server *http.Server

	mutex sync.Mutex // Mutex protects non-threadsafe

	log   logrus.FieldLogger
	cache *cache
}

func (c *checker) AddCheck(name string, checkable Checkable) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *checker) ListenAndServe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// StartedState returns the global startup state.
func (c *checker) StartedState() bool { _ = "STUB: not implemented"; return false }

// LiveState returns the global live state and details.
func (c *checker) LiveState() (bool, any) { _ = "STUB: not implemented"; return false, *new(any) }

// ReadyState returns the global ready state and details.
func (c *checker) ReadyState() (bool, any) { _ = "STUB: not implemented"; return false, *new(any) }

func (c *checker) checkStates() (bool, bool, bool, any, any) {
	_ = "STUB: not implemented"
	return false, false, false, *new(any), *new(any)
}

func (c *checker) liveHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (c *checker) readyHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
