package telemetry

import (
	"sync"
	"time"
)

// Latency is used to track timing between two specific events. It
// is a generic version of CallCounter and can be used to measure latency between any two events.
//
// Example:
//
//	func Foo() {
//	    latency := StartLatencyMetric(metrics, "foo")
//	    call.AddLabel("food", "burgers")
//	    // do something
//	    latency.Measure()
//	    // do other things
//	}
//
// Instances of this struct should only be created directly by this package
// and its subpackages, which define the specific metrics that are emitted.
// It is left exported for testing purposes.
type Latency struct {
	metrics Metrics
	key     []string
	labels  []Label
	start   time.Time
	mu      sync.Mutex
}

// StartLatencyMetric starts a "call", which when finished via Done() will emit timing
// and error related metrics.
func StartLatencyMetric(metrics Metrics, key string, keyn ...string) *Latency {
	_ = "STUB: not implemented"
	return nil
}

// AddLabel adds a label to be emitted with the call counter. It is safe to call
// from multiple goroutines.
func (l *Latency) AddLabel(name, value string) { _ = "STUB: not implemented"; return }

// Measure emits a latency metric based on l.start along with labels configured.
func (l *Latency) Measure() { _ = "STUB: not implemented"; return }
