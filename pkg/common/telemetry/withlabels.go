package telemetry

import "time"

type withLabels struct {
	metrics Metrics
	labels  []Label
}

var _ Metrics = (*withLabels)(nil)

func WithLabels(metrics Metrics, labels []Label) Metrics {
	_ = "STUB: not implemented"
	return *new(Metrics)
}

func (w *withLabels) SetGauge(key []string, val float32) { _ = "STUB: not implemented"; return }

func (w *withLabels) SetPrecisionGauge(key []string, val float64) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) SetGaugeWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) SetPrecisionGaugeWithLabels(key []string, val float64, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) EmitKey(key []string, val float32) { _ = "STUB: not implemented"; return }

func (w *withLabels) IncrCounter(key []string, val float32) { _ = "STUB: not implemented"; return }

func (w *withLabels) IncrCounterWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) AddSample(key []string, val float32) { _ = "STUB: not implemented"; return }

func (w *withLabels) AddSampleWithLabels(key []string, val float32, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) MeasureSince(key []string, start time.Time) { _ = "STUB: not implemented"; return }

func (w *withLabels) MeasureSinceWithLabels(key []string, start time.Time, labels []Label) {
	_ = "STUB: not implemented"
	return
}

func (w *withLabels) combineLabels(labels []Label) (combined []Label) {
	_ = "STUB: not implemented"
	return nil
}
