package telemetry

import (
	"context"
)

type statsdRunner struct {
	loadedSinks []Sink
}

func newStatsdRunner(c *MetricsConfig) (sinkRunner, error) {
	_ = "STUB: not implemented"
	return *new(sinkRunner), nil
}

func (s *statsdRunner) isConfigured() bool { _ = "STUB: not implemented"; return false }

func (s *statsdRunner) sinks() []Sink { _ = "STUB: not implemented"; return nil }

func (s *statsdRunner) run(context.Context) error {
	_ = "STUB: not implemented"
	// Nothing to do here
	return nil
}

func (s *statsdRunner) requiresTypePrefix() bool { _ = "STUB: not implemented"; return false }
