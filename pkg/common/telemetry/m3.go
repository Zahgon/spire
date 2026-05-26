package telemetry

import (
	"context"
	"io"
	"time"

	"github.com/uber-go/tally/v4"
)

var (
	// buckets for time durations, usually latency.
	// # 1879 The default tally buckets only go up to 5 seconds,
	// but agent timeout was increased to 30 seconds. We capture
	// this latency threshold.
	durationBuckets = tally.DurationBuckets{
		0 * time.Millisecond,
		10 * time.Millisecond,
		25 * time.Millisecond,
		50 * time.Millisecond,
		75 * time.Millisecond,
		100 * time.Millisecond,
		200 * time.Millisecond,
		300 * time.Millisecond,
		400 * time.Millisecond,
		500 * time.Millisecond,
		600 * time.Millisecond,
		800 * time.Millisecond,
		1 * time.Second,
		2 * time.Second,
		5 * time.Second,
		10 * time.Second,
		15 * time.Second,
		20 * time.Second,
		25 * time.Second,
		30 * time.Second,
	}

	// buckets for orders of magnitude of values, up to 100,000
	// given nature of SPIRE, we do not expect negative values
	exponentialValueBuckets = append(tally.ValueBuckets{0}, tally.MustMakeExponentialValueBuckets(1, 10, 5)...)
)

type m3Sink struct {
	closer io.Closer
	scope  tally.Scope
}

func newM3Sink(serviceName, address, env string) (*m3Sink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newM3TestSink(scope tally.Scope) *m3Sink { _ = "STUB: not implemented"; return nil }

func (m *m3Sink) SetGauge(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *m3Sink) SetPrecisionGauge(key []string, val float64) { _ = "STUB: not implemented"; return }

func (m *m3Sink) SetGaugeWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) SetPrecisionGaugeWithLabels(key []string, val float64, labels []Label) {
	_ = "STUB: not implemented"
	return
}

// Not implemented for m3
func (m *m3Sink) EmitKey([]string, float32) {
	_ = "STUB: not implemented"

	// Counters should accumulate values
	return
}

func (m *m3Sink) IncrCounter(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *m3Sink) IncrCounterWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

// Samples are for timing information, where quantiles are used
func (m *m3Sink) AddSample(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *m3Sink) AddSampleWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) subscopeWithLabels(labels []Label) tally.Scope {
	_ = "STUB: not implemented"
	return *new(tally.Scope)
}

// Flattens the key for formatting, removes spaces
func (m *m3Sink) flattenKey(parts []string) string {
	_ = "STUB: not implemented"
	// Ignore service name and type of metric as part of metric name,
	// i.e. prefer "foo_bar" to "service_counter_foo_bar"
	return ""
}

func (m *m3Sink) Shutdown() { _ = "STUB: not implemented"; return }

func labelsToTags(labels []Label) map[string]string { _ = "STUB: not implemented"; return nil }

func (m *m3Sink) setGauge(key []string, val float64, scope tally.Scope) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) getGauge(key []string, scope tally.Scope) tally.Gauge {
	_ = "STUB: not implemented"
	return *new(tally.Gauge)
}

func (m *m3Sink) incrCounter(key []string, val float32, scope tally.Scope) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) getCounter(key []string, scope tally.Scope) tally.Counter {
	_ = "STUB: not implemented"
	return *new(tally.Counter)
}

func (m *m3Sink) addSample(key []string, val float32, scope tally.Scope) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) addDurationSample(flattenedKey string, val float32, scope tally.Scope) {
	_ = "STUB: not implemented"
	return
}

func (m *m3Sink) addValueSample(flattenedKey string, val float32, scope tally.Scope) {
	_ = "STUB: not implemented"
	return
}

var _ Sink = (*m3Sink)(nil)

type m3Runner struct {
	loadedSinks []*m3Sink
}

func newM3Runner(c *MetricsConfig) (sinkRunner, error) {
	_ = "STUB: not implemented"
	return *new(sinkRunner), nil
}

func (r *m3Runner) isConfigured() bool { _ = "STUB: not implemented"; return false }

func (r *m3Runner) sinks() []Sink { _ = "STUB: not implemented"; return nil }

func (r *m3Runner) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *m3Runner) requiresTypePrefix() bool { _ = "STUB: not implemented"; return false }
