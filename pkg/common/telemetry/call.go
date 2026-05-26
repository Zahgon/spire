package telemetry

import (
	"sync"
	"time"
)

// CallCounter is used to track timing and other information about a "call". It
// is intended to be scoped to a function with a defer and a named error value,
// if applicable, like so:
//
//	func Foo() (err error) {
//	    call := StartCall(metrics, "foo")
//	    defer call.Done(&err)
//
//	    call.AddLabel("food", "burgers")
//	}
//
// See `Done` doc for labels automatically added.
//
// Instances of this struct should only be created directly by this package
// and its subpackages, which define the specific metrics that are emitted.
// It is left exported for testing purposes.
type CallCounter struct {
	metrics Metrics
	key     []string
	labels  []Label
	start   time.Time
	done    bool
	mu      sync.Mutex
}

// StartCall starts a "call", which when finished via Done() will emit timing
// and error related metrics.
func StartCall(metrics Metrics, key string, keyn ...string) *CallCounter {
	_ = "STUB: not implemented"
	return nil
}

// AddLabel adds a label to be emitted with the call counter. It is safe to call
// from multiple goroutines.
func (c *CallCounter) AddLabel(name, value string) { _ = "STUB: not implemented"; return }

// Discard marks the call counter as done without emitting any metrics.
// This is used to silently drop metrics for calls that should not be
// tracked (e.g. the agent's own health check loopback calls).
func (c *CallCounter) Discard() {
	_ = "STUB: not implemented"

	// Done finishes the "call" and emits metrics. No other calls to the CallCounter
	// should be done during or after the call to Done. In other words, it is not
	// thread-safe and is intended to be the final call to the CallCounter struct.
	// Emits latency and counter metrics, including adding a Status label according
	// to gRPC code of the given error. If nil error, the code is OK (success).
	return
}

func (c *CallCounter) Done(errp *error) { _ = "STUB: not implemented"; return }
