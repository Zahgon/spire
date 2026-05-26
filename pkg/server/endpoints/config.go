package endpoints

import (
	"context"
	"net"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"github.com/spiffe/spire/pkg/server/api"
	loggerv1 "github.com/spiffe/spire/pkg/server/api/logger/v1"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	bundle_client "github.com/spiffe/spire/pkg/server/bundle/client"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/ca/manager"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/pkg/server/endpoints/bundle"
	"github.com/spiffe/spire/pkg/server/svid"
)

// Config is a configuration for endpoints
type Config struct {
	// TPCAddr is the address to bind the TCP listener to.
	TCPAddr *net.TCPAddr

	// LocalAddr is the local address to bind the listener to.
	LocalAddr net.Addr

	// The svid rotator used to obtain the latest server credentials
	SVIDObserver svid.Observer

	// The server's configured trust domain. Used for validation, server SVID, etc.
	TrustDomain spiffeid.TrustDomain

	// Plugin catalog
	Catalog catalog.Catalog

	// Server CA for signing SVIDs
	ServerCA ca.ServerCA

	// Bundle endpoint configuration
	BundleEndpoint bundle.EndpointConfig

	// Authority manager
	AuthorityManager manager.AuthorityManager

	// Makes policy decisions
	AuthPolicyEngine *authpolicy.Engine

	// The logger for the endpoints subsystem
	Log logrus.FieldLogger

	// The root logger for the entire process
	RootLog loggerv1.Logger

	// The default (original config) log level
	LaunchLogLevel logrus.Level

	Metrics telemetry.Metrics

	// RateLimit holds rate limiting configurations.
	RateLimit RateLimitConfig

	Uptime func() time.Duration

	Clock clock.Clock

	// CacheReloadInterval controls how often the in-memory entry cache reloads
	CacheReloadInterval time.Duration

	// CacheReloadInterval controls how often the in-memory events based cache full reloads
	FullCacheReloadInterval time.Duration

	// EventsBasedCache enabled event driven cache reloads
	EventsBasedCache bool

	// PruneEventsOlderThan controls how long events can live before they are pruned
	PruneEventsOlderThan time.Duration

	// EventTimeout controls how long to wait for an event before giving up
	EventTimeout time.Duration

	AuditLogEnabled bool

	// ProxyProtocolTrustedCIDRs is a list of trusted CIDRs for PROXY protocol.
	// When non-empty, PROXY protocol is enabled and only connections from
	// these CIDRs are allowed to send PROXY headers.
	ProxyProtocolTrustedCIDRs []string

	// AdminIDs are a list of fixed IDs that when presented by a caller in an
	// X509-SVID, are granted admin rights.
	AdminIDs []spiffeid.ID

	BundleManager *bundle_client.Manager

	// TLSPolicy determines the post-quantum-safe policy used for all TLS
	// connections.
	TLSPolicy tlspolicy.Policy

	MaxAttestedNodeInfoStaleness time.Duration

	AgentSpiffeIdAsSelector bool
}

func (c *Config) maybeMakeBundleEndpointServer() (Server, func(context.Context) error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// Start watching for file changes

func (c *Config) makeAPIServers(entryFetcher api.AuthorizedEntryFetcher) APIServers {
	_ = "STUB: not implemented"
	return *new(APIServers)
}
