package fakemetrics

import (
	"sync"
	"time"

	"github.com/spiffe/spire/pkg/common/telemetry"
)

type MetricType int

const (
	SetGaugeType MetricType = iota
	SetGaugeWithLabelsType
	EmitKeyType
	IncrCounterType
	IncrCounterWithLabelsType
	AddSampleType
	AddSampleWithLabelsType
	MeasureSinceType
	MeasureSinceWithLabelsType
)

type FakeMetrics struct {
	metrics []MetricItem
	mu      sync.Mutex
}

type MetricItem struct {
	Type   MetricType
	Key    []string
	Val    float64
	Labels []telemetry.Label
	Start  time.Time
}

func New() *FakeMetrics { _ = "STUB: not implemented"; return nil }

func (m *FakeMetrics) Reset() { _ = "STUB: not implemented"; return }

// AllMetrics return all collected metrics
func (m *FakeMetrics) AllMetrics() []MetricItem { _ = "STUB: not implemented"; return nil }

func (m *FakeMetrics) SetGauge(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *FakeMetrics) SetPrecisionGauge(key []string, val float64) {
	_ = "STUB: not implemented"
	return
}

func (m *FakeMetrics) SetGaugeWithLabels(key []string, val float32, labels []telemetry.Label) {
	_ = "STUB: not implemented"
	return
}

func (m *FakeMetrics) SetPrecisionGaugeWithLabels(key []string, val float64, labels []telemetry.Label) {
	_ = "STUB: not implemented"
	return
}

func (m *FakeMetrics) EmitKey(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *FakeMetrics) IncrCounter(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *FakeMetrics) IncrCounterWithLabels(key []string, val float32, labels []telemetry.Label) {
	_ = "STUB: not implemented"
	return
}

func (m *FakeMetrics) AddSample(key []string, val float32) { _ = "STUB: not implemented"; return }

func (m *FakeMetrics) AddSampleWithLabels(key []string, val float32, labels []telemetry.Label) {
	_ = "STUB: not implemented"
	return
}

func (m *FakeMetrics) MeasureSince(key []string, _ time.Time) { _ = "STUB: not implemented"; return }

// TODO: record `start` when it is convenient to thread a clock through all the telemetry helpers

func (m *FakeMetrics) MeasureSinceWithLabels(key []string, _ time.Time, labels []telemetry.Label) {
	_ = "STUB: not implemented"
	return
}

// TODO: record `start` when it is convenient to thread a clock through all the telemetry helpers
