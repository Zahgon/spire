package server

import (
	"context"
	_ "net/http/pprof" //nolint: gosec // import registers routes on DefaultServeMux

	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	bundle_client "github.com/spiffe/spire/pkg/server/bundle/client"
	"github.com/spiffe/spire/pkg/server/bundle/pubmanager"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/ca/manager"
	"github.com/spiffe/spire/pkg/server/ca/rotator"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/pkg/server/credtemplate"
	"github.com/spiffe/spire/pkg/server/credvalidator"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/endpoints"
	"github.com/spiffe/spire/pkg/server/hostservice/agentstore"
	"github.com/spiffe/spire/pkg/server/hostservice/identityprovider"
	"github.com/spiffe/spire/pkg/server/node"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher"
	"github.com/spiffe/spire/pkg/server/registration"
	"github.com/spiffe/spire/pkg/server/svid"
)

const (
	invalidTrustDomainAttestedNode = "An attested node with trust domain '%v' has been detected, " +
		"which does not match the configured trust domain of '%v'. Agents may need to be reconfigured to use new trust domain"
	invalidTrustDomainRegistrationEntry = "a registration entry with trust domain '%v' has been detected, " +
		"which does not match the configured trust domain of '%v'. If you want to change the trust domain, " +
		"please delete all existing registration entries"
	invalidSpiffeIDRegistrationEntry = "registration entry with id %v is malformed because invalid SPIFFE ID: %v"
	invalidSpiffeIDAttestedNode      = "could not parse SPIFFE ID, from attested node"

	pageSize = 1
)

type Server struct {
	config Config
}

// Run the server
// This method initializes the server, including its plugins,
// and then blocks until it's shut down or an error is encountered.
func (s *Server) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) ValidateConfig(ctx context.Context) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) run(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	// Log configuration values that are useful for debugging
	return nil
}

// create the data directory if needed

// Create the identity provider host service. It will not be functional
// until the call to SetDeps() below. There is some tricky initialization
// stuff going on since the identity provider host service requires plugins
// to do its job. RPC's from plugins to the identity provider before
// SetDeps() has been called will fail with a PreCondition status.

// Create the agent store host service. It will not be functional
// until the call to SetDeps() below.

// CA manager needs to be initialized before the rotator, otherwise the
// server CA plugin won't be able to sign CSRs

// Set the identity provider dependencies

// Return the server identity itself

// Set the agent store dependencies

// Wait for the server to start listening before proceeding with health
// checks.

func (s *Server) setupProfiling(ctx context.Context) (stop func()) {
	_ = "STUB: not implemented"
	return nil
}

// kick off a goroutine to serve the pprof endpoints and one to
// gracefully shut down the server when profiling is being torn down

func (s *Server) loadCatalog(ctx context.Context, metrics telemetry.Metrics, identityProvider *identityprovider.IdentityProvider, agentStore *agentstore.AgentStore,
	healthChecker health.Checker,
) (*catalog.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newCredBuilder(cat catalog.Catalog) (*credtemplate.Builder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newCredValidator() (*credvalidator.Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newCA(metrics telemetry.Metrics, credBuilder *credtemplate.Builder, credValidator *credvalidator.Validator, healthChecker health.Checker) *ca.CA {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) newCAManager(ctx context.Context, cat catalog.Catalog, metrics telemetry.Metrics, serverCA *ca.CA, credBuilder *credtemplate.Builder, credValidator *credvalidator.Validator) (*manager.Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newCASync(ctx context.Context, healthChecker health.Checker, caManager *manager.Manager) (*rotator.Rotator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newRegistrationManager(cat catalog.Catalog, metrics telemetry.Metrics) *registration.Manager {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) newNodeManager(cat catalog.Catalog, metrics telemetry.Metrics) *node.Manager {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) newSVIDRotator(ctx context.Context, serverCA ca.ServerCA, metrics telemetry.Metrics) (*svid.Rotator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newEndpointsServer(ctx context.Context, catalog catalog.Catalog, svidObserver svid.Observer, serverCA ca.ServerCA, metrics telemetry.Metrics, authorityManager manager.AuthorityManager, authPolicyEngine *authpolicy.Engine, bundleManager *bundle_client.Manager) (endpoints.Server, error) {
	_ = "STUB: not implemented"
	return *new(endpoints.Server), nil
}

func (s *Server) newBundleManager(cat catalog.Catalog, metrics telemetry.Metrics) *bundle_client.Manager {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) newBundlePublishingManager(bundlePublishers []bundlepublisher.BundlePublisher, ds datastore.DataStore) (*pubmanager.Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) validateTrustDomain(ctx context.Context, ds datastore.DataStore) error {
	_ = "STUB: not implemented"
	return nil
}

// Get only first page with a single element

// Get only first page with a single element

// CheckHealth is used as a top-level health check for the Server.
func (s *Server) CheckHealth() health.State {
	_ = "STUB: not implemented"
	return *

	// The API is served only after the server CA has been
	// signed by upstream. Hence, both live and ready checks
	// are determined by whether the bundles are received or not.
	// TODO: Better live check for server.
	new(health.State)
}

func (s *Server) tryGetBundle() error { _ = "STUB: not implemented"; return nil }

// Currently using the ability to fetch a bundle as the health check. This
// **could** be problematic if the Upstream CA signing process is lengthy.
// As currently coded however, the API isn't served until after
// the server CA has been signed by upstream.

type serverHealthDetails struct {
	GetBundleErr string `json:"get_bundle_err,omitempty"`
}

func errString(err error) string { _ = "STUB: not implemented"; return "" }
