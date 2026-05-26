package agent

import (
	"context"
	"crypto/x509"
	_ "net/http/pprof" //nolint: gosec // import registers routes on DefaultServeMux
	"time"

	admin_api "github.com/spiffe/spire/pkg/agent/api"
	node_attestor "github.com/spiffe/spire/pkg/agent/attestor/node"
	workload_attestor "github.com/spiffe/spire/pkg/agent/attestor/workload"
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/endpoints"
	"github.com/spiffe/spire/pkg/agent/manager"
	"github.com/spiffe/spire/pkg/agent/manager/storecache"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/agent/storage"
	"github.com/spiffe/spire/pkg/agent/svid/store"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/util"
	_ "golang.org/x/net/trace" // registers handlers on the DefaultServeMux
)

const (
	bootstrapBackoffInterval         = 5 * time.Second
	bootstrapBackoffMaxElapsedTime   = 1 * time.Minute
	startHealthChecksTimeout         = 8 * time.Second
	rebootstrapBackoffMaxElapsedTime = 24 * time.Hour
)

type Agent struct {
	c       *Config
	started bool
}

// Run the agent
// This method initializes the agent, including its plugins,
// and then blocks on the main event loop.
func (a *Agent) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *Agent) setupProfiling(ctx context.Context) (stop func()) {
	_ = "STUB: not implemented"
	return nil
}

// kick off a goroutine to serve the pprof endpoints and one to
// gracefully shut down the server when profiling is being torn down

func (a *Agent) attest(ctx context.Context, sto storage.Storage, cat catalog.Catalog, metrics telemetry.Metrics, na nodeattestor.NodeAttestor, bootstrapTrustBundle []*x509.Certificate, insecureBootstrap bool) (*node_attestor.AttestationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Agent) newManager(ctx context.Context, sto storage.Storage, cat catalog.Catalog, metrics telemetry.Metrics, as *node_attestor.AttestationResult, cache *storecache.Cache, na nodeattestor.NodeAttestor) (manager.Manager, error) {
	_ = "STUB: not implemented"
	return *new(manager.Manager), nil
}

func (a *Agent) newSVIDStoreCache(metrics telemetry.Metrics) *storecache.Cache {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) newSVIDStoreService(cache *storecache.Cache, cat catalog.Catalog, metrics telemetry.Metrics) *store.SVIDStoreService {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) newEndpoints(metrics telemetry.Metrics, mgr manager.Manager, attestor workload_attestor.Attestor) endpoints.Server {
	_ = "STUB: not implemented"
	return *new(endpoints.Server)
}

func (a *Agent) newAdminEndpoints(metrics telemetry.Metrics, mgr manager.Manager, attestor workload_attestor.Attestor, authorizedDelegates []string) admin_api.Server {
	_ = "STUB: not implemented"
	return *new(admin_api.Server)
}

// CheckHealth is used as a top-level health check for the agent.
func (a *Agent) CheckHealth() health.State { _ = "STUB: not implemented"; return *new(health.State) }

// Both liveness and readiness checks are done by
// agents ability to create new Workload API client
// for the X509SVID service.
// TODO: Better live check for agent.

func (a *Agent) checkWorkloadAPI() error { _ = "STUB: not implemented"; return nil }

// Only an unavailable status fails the health check.

type agentHealthDetails struct {
	WorkloadAPIErr string `json:"make_new_x509_err,omitempty"`
}

func errString(suppress bool, err error) string { _ = "STUB: not implemented"; return "" }

func (a *Agent) startHealthChecks(readyForHealthChecks chan struct{}, taskRunner *util.TaskRunner, healthChecker health.ServableChecker) {
	_ = "STUB: not implemented"
	return
}

// Endpoints are ready for health checks, proceed with health checks.

// Timeout waiting for endpoints to start listening.
