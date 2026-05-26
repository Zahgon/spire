package telemetry

import (
	"time"
)

// Blackhole implements the Metrics interface, but throws away the metric data
// Useful for satisfying the Metrics interface when testing code which depends on it.
type Blackhole struct{}

var _ Metrics = Blackhole{}

func (Blackhole) SetGauge([]string, float32)                    { _ = "STUB: not implemented"; return }
func (Blackhole) SetGaugeWithLabels([]string, float32, []Label) { _ = "STUB: not implemented"; return }
func (Blackhole) SetPrecisionGauge([]string, float64)           { _ = "STUB: not implemented"; return }
func (Blackhole) SetPrecisionGaugeWithLabels([]string, float64, []Label) {
	_ = "STUB: not implemented"
	return
}
func (Blackhole) EmitKey([]string, float32)     { _ = "STUB: not implemented"; return }
func (Blackhole) IncrCounter([]string, float32) { _ = "STUB: not implemented"; return }
func (Blackhole) IncrCounterWithLabels([]string, float32, []Label) {
	_ = "STUB: not implemented"
	return
}
func (Blackhole) AddSample([]string, float32)                    { _ = "STUB: not implemented"; return }
func (Blackhole) AddSampleWithLabels([]string, float32, []Label) { _ = "STUB: not implemented"; return }
func (Blackhole) MeasureSince([]string, time.Time)               { _ = "STUB: not implemented"; return }
func (Blackhole) MeasureSinceWithLabels([]string, time.Time, []Label) {
	_ = "STUB: not implemented"
	return
}
