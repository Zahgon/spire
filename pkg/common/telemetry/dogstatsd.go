package telemetry

import (
	"context"
)

type dogStatsdRunner struct {
	loadedSinks []Sink
}

func newDogStatsdRunner(c *MetricsConfig) (sinkRunner, error) {
	_ = "STUB: not implemented"
	return *new(sinkRunner), nil
}

func (d *dogStatsdRunner) isConfigured() bool { _ = "STUB: not implemented"; return false }

func (d *dogStatsdRunner) sinks() []Sink { _ = "STUB: not implemented"; return nil }

func (d *dogStatsdRunner) run(context.Context) error {
	_ = "STUB: not implemented"
	// Nothing to do here
	return nil
}

func (d *dogStatsdRunner) requiresTypePrefix() bool { _ = "STUB: not implemented"; return false }
