package endpoints

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/spiffe/spire/pkg/server/endpoints/bundle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	debugv1_pb "github.com/spiffe/spire-api-sdk/proto/spire/api/server/debug/v1"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	localauthorityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/localauthority/v1"
	loggerv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/logger/v1"
	svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	trustdomainv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/trustdomain/v1"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/svid"
)

const (
	// This is the maximum amount of time an agent connection may exist before
	// the server sends a hangup request. This enables agents to more dynamically
	// route to the server in the case of a change in DNS membership.
	defaultMaxConnectionAge = 3 * time.Minute

	// This is the default amount of time between two reloads of the in-memory
	// entry cache.
	defaultCacheReloadInterval = 5 * time.Second

	// This is the default amount of time between full refreshes of the in-memory
	// entry cache.
	defaultFullCacheReloadInterval = 24 * time.Hour

	// This is the default amount of time events live before they are pruned
	defaultPruneEventsOlderThan = 12 * time.Hour

	// This is the default amount of time to wait for an event before giving up
	defaultEventTimeout = 15 * time.Minute

	// This is the time to wait for graceful termination of the gRPC server
	// before forcefully terminating.
	gracefulStopTimeout = 10 * time.Second
)

// Server manages gRPC and HTTP endpoint lifecycle
type Server interface {
	// ListenAndServe starts all endpoint servers and blocks until the context
	// is canceled or any of the servers fails to run. If the context is
	// canceled, the function returns nil. Otherwise, the error from the failed
	// server is returned.
	ListenAndServe(ctx context.Context) error

	// WaitForListening blocks until the server starts listening.
	WaitForListening()
}

type Endpoints struct {
	TCPAddr                      *net.TCPAddr
	LocalAddr                    net.Addr
	SVIDObserver                 svid.Observer
	TrustDomain                  spiffeid.TrustDomain
	DataStore                    datastore.DataStore
	BundleCache                  *bundle.Cache
	APIServers                   APIServers
	BundleEndpointServer         Server
	Log                          logrus.FieldLogger
	Metrics                      telemetry.Metrics
	RateLimit                    RateLimitConfig
	NodeCacheRebuildTask         func(context.Context) error
	EntryFetcherCacheRebuildTask func(context.Context) error
	EntryFetcherPruneEventsTask  func(context.Context) error
	CertificateReloadTask        func(context.Context) error
	AuditLogEnabled              bool
	ProxyProtocolTrustedCIDRs    []string
	AuthPolicyEngine             *authpolicy.Engine
	AdminIDs                     []spiffeid.ID
	TLSPolicy                    tlspolicy.Policy
	MaxAttestedNodeInfoStaleness time.Duration
	nodeCache                    api.AttestedNodeCache

	hooks struct {
		// test hook used to indicate that is listening
		listening chan struct{}
	}
}

type APIServers struct {
	AgentServer          agentv1.AgentServer
	BundleServer         bundlev1.BundleServer
	DebugServer          debugv1_pb.DebugServer
	EntryServer          entryv1.EntryServer
	HealthServer         grpc_health_v1.HealthServer
	LoggerServer         loggerv1.LoggerServer
	SVIDServer           svidv1.SVIDServer
	TrustDomainServer    trustdomainv1.TrustDomainServer
	LocalAUthorityServer localauthorityv1.LocalAuthorityServer
}

// RateLimitConfig holds rate limiting configurations.
type RateLimitConfig struct {
	// Attestation, if true, rate limits attestation
	Attestation bool

	// Signing, if true, rate limits JWT and X509 signing requests
	Signing bool
}

// New creates new endpoints struct
func New(ctx context.Context, c Config) (*Endpoints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cacheRebuildTask will take care of rebuilding the node cache

// ListenAndServe starts all endpoint servers and blocks until the context
// is canceled or any of the servers fails to run. If the context is
// canceled, the function returns nil. Otherwise, the error from the failed
// server is returned.
func (e *Endpoints) ListenAndServe(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// TCP and UDS

// UDS only

func (e *Endpoints) createTCPServer(ctx context.Context, unaryInterceptor grpc.UnaryServerInterceptor, streamInterceptor grpc.StreamServerInterceptor) *grpc.Server {
	_ = "STUB: not implemented"
	return nil
}

// Disable session ticket resumption so that VerifyPeerCertificate is
// called on every connection, ensuring the peer certificate chain is
// always validated against the current trust bundle.

func (e *Endpoints) createUDSServer(unaryInterceptor grpc.UnaryServerInterceptor, streamInterceptor grpc.StreamServerInterceptor) *grpc.Server {
	_ = "STUB: not implemented"
	return nil
}

// runTCPServer will start the server and block until it exits, or we are dying.
func (e *Endpoints) runTCPServer(ctx context.Context, server *grpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip use of tomb here so we don't pollute a clean shutdown with errors

// wrapListenerWithProxyProtocol wraps a net.Listener with PROXY protocol
// support, restricting header acceptance to the given trusted CIDRs.
func wrapListenerWithProxyProtocol(l net.Listener, trustedCIDRs []string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// runLocalAccess will start a grpc server to be accessed locally
// and block until it exits, or we are dying.
func (e *Endpoints) runLocalAccess(ctx context.Context, server *grpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip use of tomb here so we don't pollute a clean shutdown with errors

// handleShutdown is a helper function for gracefully terminating the grpc server.
// if the server does not terminate within the GratefulStopWait deadline, the server
// will be forcibly stopped.
func (e *Endpoints) handleShutdown(server *grpc.Server, errChan <-chan error, log *logrus.Entry) {
	_ = "STUB: not implemented"
	return
}

// getTLSConfig returns a TLS Config hook for the gRPC server
func (e *Endpoints) getTLSConfig(ctx context.Context) func(*tls.ClientHelloInfo) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil
}

// provided client certificates will be validated using the custom VerifyPeerCertificate hook

func (e *Endpoints) makeInterceptors() (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor), *new(grpc.StreamServerInterceptor)
}

func (e *Endpoints) triggerListeningHook() { _ = "STUB: not implemented"; return }

func (e *Endpoints) WaitForListening() { _ = "STUB: not implemented"; return }
