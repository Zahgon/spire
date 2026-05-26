package attestor

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/common"
)

type attestor struct {
	c *Config
}

type Attestor interface {
	Attest(ctx context.Context, pid int) ([]*common.Selector, error)
}

func New(config *Config) Attestor { _ = "STUB: not implemented"; return *new(Attestor) }

func newAttestor(config *Config) *attestor { _ = "STUB: not implemented"; return nil }

type Config struct {
	Catalog catalog.Catalog
	Log     logrus.FieldLogger
	Metrics telemetry.Metrics

	// Test hook called when selectors are obtained from a workload attestor plugin
	selectorHook func([]*common.Selector)
}

// Attest invokes all workload attestor plugins against the provided PID. If an error
// is encountered, it is logged and selectors from the failing plugin are discarded.
func (wla *attestor) Attest(ctx context.Context, pid int) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect the results

// The agent health check currently exercises the Workload API. Since this
// can happen with some frequency, it has a tendency to fill up logs with
// hard-to-filter details if we're not careful (e.g. issue #1537). Only log
// if it is not the agent itself.

// invokeAttestor invokes attestation against the supplied plugin. Should be called from a goroutine.
func (wla *attestor) invokeAttestor(ctx context.Context, a workloadattestor.WorkloadAttestor, pid int) (_ []*common.Selector, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
