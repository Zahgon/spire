package util

import (
	"context"
	"crypto/x509"
	"flag"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	localauthorityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/localauthority/v1"
	loggerv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/logger/v1"
	svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	trustdomainv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/trustdomain/v1"
	api_types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	DefaultSocketPath    = "/tmp/spire-server/private/api.sock"
	DefaultNamedPipeName = "\\spire-server\\private\\api"
	FormatPEM            = "pem"
	FormatSPIFFE         = "spiffe"
)

func NewGRPCClient(addr string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServerClient interface {
	Release()
	NewAgentClient() agentv1.AgentClient
	NewBundleClient() bundlev1.BundleClient
	NewEntryClient() entryv1.EntryClient
	NewLoggerClient() loggerv1.LoggerClient
	NewSVIDClient() svidv1.SVIDClient
	NewTrustDomainClient() trustdomainv1.TrustDomainClient
	NewLocalAuthorityClient() localauthorityv1.LocalAuthorityClient
	NewHealthClient() grpc_health_v1.HealthClient
}

func NewServerClient(addr string) (ServerClient, error) {
	_ = "STUB: not implemented"
	return *new(ServerClient), nil
}

type serverClient struct {
	conn *grpc.ClientConn
}

func (c *serverClient) Release() { _ = "STUB: not implemented"; return }

func (c *serverClient) NewAgentClient() agentv1.AgentClient {
	_ = "STUB: not implemented"
	return *new(agentv1.AgentClient)
}

func (c *serverClient) NewBundleClient() bundlev1.BundleClient {
	_ = "STUB: not implemented"
	return *new(bundlev1.BundleClient)
}

func (c *serverClient) NewEntryClient() entryv1.EntryClient {
	_ = "STUB: not implemented"
	return *new(entryv1.EntryClient)
}

func (c *serverClient) NewLoggerClient() loggerv1.LoggerClient {
	_ = "STUB: not implemented"
	return *new(loggerv1.LoggerClient)
}

func (c *serverClient) NewSVIDClient() svidv1.SVIDClient {
	_ = "STUB: not implemented"
	return *new(svidv1.SVIDClient)
}

func (c *serverClient) NewTrustDomainClient() trustdomainv1.TrustDomainClient {
	_ = "STUB: not implemented"
	return *new(trustdomainv1.TrustDomainClient)
}

func (c *serverClient) NewHealthClient() grpc_health_v1.HealthClient {
	_ = "STUB: not implemented"
	return *new(grpc_health_v1.HealthClient)
}

func (c *serverClient) NewLocalAuthorityClient() localauthorityv1.LocalAuthorityClient {
	_ = "STUB: not implemented"
	return *new(localauthorityv1.LocalAuthorityClient)
}

// Pluralizer concatenates `singular` to `msg` when `val` is one, and
// `plural` on all other occasions. It is meant to facilitate friendlier
// CLI output.
func Pluralizer(msg string, singular string, plural string, val int) string {
	_ = "STUB: not implemented"
	return ""
}

// Command is a common interface for commands in this package. the adapter
// can adapter this interface to the Command interface from github.com/mitchellh/cli.
type Command interface {
	Name() string
	Synopsis() string
	AppendFlags(*flag.FlagSet)
	Run(context.Context, *common_cli.Env, ServerClient) error
}

type Adapter struct {
	env *common_cli.Env
	cmd Command

	flags *flag.FlagSet

	adapterOS // OS specific
}

// AdaptCommand converts a command into one conforming to the Command interface from github.com/mitchellh/cli
func AdaptCommand(env *common_cli.Env, cmd Command) *Adapter { _ = "STUB: not implemented"; return nil }

func (a *Adapter) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

func (a *Adapter) Help() string { _ = "STUB: not implemented"; return "" }

func (a *Adapter) Synopsis() string { _ = "STUB: not implemented"; return "" }

// parseSelector parses a CLI string from type:value into a selector type.
// Everything to the right of the first ":" is considered a selector value.
func ParseSelector(str string) (*api_types.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Strip the trailing delimiter

func ParseBundle(bundleBytes []byte, format, id string) (*api_types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleProtoFromX509Authorities creates a Bundle API type from a trustdomain and
// a list of root CAs.
func bundleProtoFromX509Authorities(trustDomain string, rootCAs []*x509.Certificate) *api_types.Bundle {
	_ = "STUB: not implemented"
	return nil
}

// protoFromSpiffeBundle converts a bundle from the given *spiffebundle.Bundle to *api_types.Bundle
func protoFromSpiffeBundle(bundle *spiffebundle.Bundle) (*api_types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// protoFromX509Certificates converts X.509 certificates from the given []*x509.Certificate to []*types.X509Certificate
func protoFromX509Certificates(certs []*x509.Certificate) []*api_types.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}
