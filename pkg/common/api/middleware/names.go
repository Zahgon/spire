package middleware

import (
	"context"
	"strings"
	"sync"

	"github.com/spiffe/spire/pkg/common/api"
)

const (
	serverAPIPrefix = "spire.api.server."

	WorkloadAPIServiceName             = "SpiffeWorkloadAPI"
	WorkloadAPIServiceShortName        = "WorkloadAPI"
	EnvoySDSv3ServiceName              = "envoy.service.secret.v3.SecretDiscoveryService"
	EnvoySDSv3ServiceShortName         = "SDS.v3"
	HealthServiceName                  = "grpc.health.v1.Health"
	HealthServiceShortName             = "Health"
	LoggerServiceName                  = "logger.v1.Logger"
	LoggerServiceShortName             = "Logger"
	DebugServiceName                   = "spire.agent.debug.v1.Debug"
	DebugServiceShortName              = "Debug"
	DelegatedIdentityServiceName       = "spire.api.agent.delegatedidentity.v1.DelegatedIdentity"
	DelegatedIdentityServiceShortName  = "DelegatedIdentity"
	ServerReflectionServiceName        = "grpc.reflection.v1.ServerReflection"
	ServerReflectionV1AlphaServiceName = "grpc.reflection.v1alpha.ServerReflection"
	SubscribeToX509SVIDsMethodName     = "SubscribeToX509SVIDs"
	SubscribeToX509SVIDsMetricKey      = "subscribe_to_x509_svids"
)

var (
	serviceReplacer = strings.NewReplacer(
		serverAPIPrefix, "",
		WorkloadAPIServiceName, WorkloadAPIServiceShortName,
		EnvoySDSv3ServiceName, EnvoySDSv3ServiceShortName,
		HealthServiceName, HealthServiceShortName,
		LoggerServiceName, LoggerServiceShortName,
		DebugServiceName, DebugServiceShortName,
		DelegatedIdentityServiceName, DelegatedIdentityServiceShortName,
	)

	// methodMetricKeyReplacer allows adding replacement for method names that
	// are not parsed correctly by metricKey func. Since changes to metricKey would
	// be breaking, add a direct replacement here for the required metric key.
	methodMetricKeyReplacer = strings.NewReplacer(
		SubscribeToX509SVIDsMethodName, SubscribeToX509SVIDsMetricKey,
	)

	// namesCache caches parsed names
	namesCache sync.Map
)

// withNames returns a context and the names parsed out of the given full
// method. If the given context already has the parsed names, then those names
// are returned. Otherwise, a global cache is checked for the names, keyed by
// the full method. If present, the cached names are returned. Otherwise, the
// full method is parsed and the names cached and returned along with an
// embellished context.
func withNames(ctx context.Context, fullMethod string) (context.Context, api.Names) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(api.Names)
}

// makeNames parses a gRPC full method name into individual parts.  It expects
// the input to be well-formed since it gets its input from gRPC generated
// names. It will not panic if given bad input, but will not provide meaningful
// names.
func makeNames(fullMethod string) (names api.Names) {
	_ = "STUB: not implemented"
	// Strip the leading slash. It should always be present in practice.
	return *new(api.Names)
}

// Parse the slash separated service and method name. The separating slash
// should always be present in practice.

// metricKey converts an RPC service or method name into one appropriate for
// metrics use. It converts PascalCase into snake_case, also converting any
// non-alphanumeric rune into an underscore.
func metricKey(s string) string { _ = "STUB: not implemented"; return "" }

// Add an underscore if the current rune:
// - is uppercase
// - not the first rune
// - is followed or preceded by a lowercase rune
// - was not preceded by an underscore in the output
