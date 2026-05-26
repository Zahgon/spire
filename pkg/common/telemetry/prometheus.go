package telemetry

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/sirupsen/logrus"
)

type prometheusRunner struct {
	c      *PrometheusConfig
	log    logrus.FieldLogger
	server *http.Server
	sink   Sink
}

func newPrometheusRunner(c *MetricsConfig) (sinkRunner, error) {
	_ = "STUB: not implemented"
	return *new(sinkRunner), nil
}

func (p *prometheusRunner) isConfigured() bool { _ = "STUB: not implemented"; return false }

func (p *prometheusRunner) sinks() []Sink { _ = "STUB: not implemented"; return nil }

func (p *prometheusRunner) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *prometheusRunner) requiresTypePrefix() bool { _ = "STUB: not implemented"; return false }

func (p *prometheusRunner) newTLSConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// easier to return the tls config rather than assigning it to the server directly from maintenance perspective
