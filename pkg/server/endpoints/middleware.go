package endpoints

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/api/bundle/v1"
	"github.com/spiffe/spire/pkg/server/api/middleware"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	"github.com/spiffe/spire/pkg/server/ca/manager"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/test/clock"
)

func Middleware(log logrus.FieldLogger, metrics telemetry.Metrics, ds datastore.DataStore, nodeCache api.AttestedNodeCache, maxAttestedNodeInfoStaleness time.Duration, clk clock.Clock, rlConf RateLimitConfig, policyEngine *authpolicy.Engine, auditLogEnabled bool, adminIDs []spiffeid.ID) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

// Add audit log with local tracking enabled

func EntryFetcher(ds datastore.DataStore) middleware.EntryFetcher {
	_ = "STUB: not implemented"
	return *new(middleware.EntryFetcher)
}

func UpstreamPublisher(jwtKeyPublisher manager.JwtKeyPublisher) bundle.UpstreamPublisher {
	_ = "STUB: not implemented"
	return *new(bundle.UpstreamPublisher)
}

func AgentAuthorizer(ds datastore.DataStore, nodeCache api.AttestedNodeCache, maxAttestedNodeInfoStaleness time.Duration, clk clock.Clock) middleware.AgentAuthorizer {
	_ = "STUB: not implemented"
	return *new(middleware.AgentAuthorizer)
}

// AttestedNode not found in local cache, will fetch from the datastore

// Cached AttestedNode is stale, will attempt to refresh from the database

// Attested node was not found in the cache, will fetch from the datastore

// AgentSVID matches the current serial number, access granted.

// Could not validate the agent using the cache attested node information
// so we'll try fetching the up to date data from the datastore.

// AgentSVID matches the current serial number, access granted

// AgentSVID matches the new serial number, access granted
// Also update the attested node agent serial number from 'new' to 'current'

func RateLimits(config RateLimitConfig) map[string]api.RateLimiter {
	_ = "STUB: not implemented"
	return nil
}
