package telemetry

import (
	"context"
	"time"

	"github.com/hashicorp/go-metrics"
)

const timerGranularity = time.Millisecond

// Label is a label/tag for a metric
type Label = metrics.Label

// Sink is an interface for emitting metrics
type Sink = metrics.MetricSink

// Metrics is an interface for all metrics plugins and services
type Metrics interface {
	// A Gauge should retain the last value it is set to
	SetGauge(key []string, val float32)
	SetGaugeWithLabels(key []string, val float32, labels []Label)
	SetPrecisionGauge(key []string, val float64)
	SetPrecisionGaugeWithLabels(key []string, val float64, labels []Label)

	// Should emit a Key/Value pair for each call
	EmitKey(key []string, val float32)

	// Counters should accumulate values
	IncrCounter(key []string, val float32)
	IncrCounterWithLabels(key []string, val float32, labels []Label)

	// Samples are for timing information, where quantiles are used
	AddSample(key []string, val float32)
	AddSampleWithLabels(key []string, val float32, labels []Label)

	// A convenience function for measuring elapsed time with a single line
	MeasureSince(key []string, start time.Time)
	MeasureSinceWithLabels(key []string, start time.Time, labels []Label)
}

type MetricsImpl struct {
	*metrics.Metrics

	c       *MetricsConfig
	runners []sinkRunner
	// Each instance of metrics.Metrics in the slice corresponds to one metrics sink type
	metricsSinks           []*metrics.Metrics
	enableTrustDomainLabel bool
}

var _ Metrics = (*MetricsImpl)(nil)

// NewMetrics returns a Metric implementation
func NewMetrics(c *MetricsConfig) (*MetricsImpl, error) { _ = "STUB: not implemented"; return nil, nil }

// ListenAndServe starts the metrics process
func (m *MetricsImpl) ListenAndServe(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MetricsImpl) SetGauge(key []string, val float32) { _ = "STUB: not implemented"; return }

// SetGaugeWithLabels delegates to embedded metrics, sanitizing labels
func (m *MetricsImpl) SetGaugeWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *MetricsImpl) SetPrecisionGauge(key []string, val float64) {
	_ = "STUB: not implemented"
	return
}

// SetPrecisionGaugeWithLabels delegates to embedded metrics, sanitizing labels
func (m *MetricsImpl) SetPrecisionGaugeWithLabels(key []string, val float64, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *MetricsImpl) EmitKey(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *MetricsImpl) IncrCounter(key []string, val float32) { _ = "STUB: not implemented"; return }

// IncrCounterWithLabels delegates to embedded metrics, sanitizing labels
func (m *MetricsImpl) IncrCounterWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *MetricsImpl) AddSample(key []string, val float32) { _ = "STUB: not implemented"; return }

// AddSampleWithLabels delegates to embedded metrics, sanitizing labels
func (m *MetricsImpl) AddSampleWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (m *MetricsImpl) MeasureSince(key []string, start time.Time) {
	_ = "STUB: not implemented"
	return
}

// MeasureSinceWithLabels delegates to embedded metrics, sanitizing labels
func (m *MetricsImpl) MeasureSinceWithLabels(key []string, start time.Time, labels []Label) {
	_ = "STUB: not implemented"
	return
}
