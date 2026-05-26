package telemetry

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-metrics"
	"github.com/sirupsen/logrus"
)

const (
	inmemInterval  = 1 * time.Second
	inmemRetention = 1 * time.Hour
)

type inmemRunner struct {
	log        logrus.FieldLogger
	w          io.Writer
	loadedSink *metrics.InmemSink
}

func newInmemRunner(c *MetricsConfig) (sinkRunner, error) {
	_ = "STUB: not implemented"
	return *new(sinkRunner), nil
}

// Don't enable If the InMem block is not present.

func (i *inmemRunner) isConfigured() bool { _ = "STUB: not implemented"; return false }

func (i *inmemRunner) sinks() []Sink { _ = "STUB: not implemented"; return nil }

func (i *inmemRunner) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (i *inmemRunner) requiresTypePrefix() bool { _ = "STUB: not implemented"; return false }
