package itclient

import (
	"context"
	"crypto"
	"crypto/x509"
	"flag"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	agent "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	bundle "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	debug "github.com/spiffe/spire-api-sdk/proto/spire/api/server/debug/v1"
	entry "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	svid "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/server/trustdomain/v1"
	"google.golang.org/grpc"
)

var (
	tdFlag               = flag.String("trustDomain", "domain.test", "server trust domain")
	socketPathFlag       = flag.String("socketPath", "unix:///tmp/spire-agent/public/api.sock", "agent socket path")
	serverAddrFlag       = flag.String("serverAddr", "spire-server:8081", "server addr")
	serverSocketPathFlag = flag.String("serverSocketPath", "unix:///tmp/spire-server/private/api.sock", "server socket path")
	expectErrorsFlag     = flag.Bool("expectErrors", false, "client is used to validate permission errors")
)

type Client struct {
	ExpectErrors bool
	Td           spiffeid.TrustDomain

	connection *grpc.ClientConn
	source     *workloadapi.X509Source
}

func New(ctx context.Context) *Client { _ = "STUB: not implemented"; return nil }

// Create X509Source

// Create connection

func NewInsecure() *Client { _ = "STUB: not implemented"; return nil }

//nolint: gosec // this is intentional for the integration test

func NewWithCert(cert *x509.Certificate, key crypto.Signer) *Client {
	_ = "STUB: not implemented"
	return nil
}

//nolint: gosec // this is intentional for the integration test

func (c *Client) Release() { _ = "STUB: not implemented"; return }

func (c *Client) BundleClient() bundle.BundleClient {
	_ = "STUB: not implemented"
	return *new(bundle.BundleClient)
}

func (c *Client) EntryClient() entry.EntryClient {
	_ = "STUB: not implemented"
	return *new(entry.EntryClient)
}

func (c *Client) SVIDClient() svid.SVIDClient {
	_ = "STUB: not implemented"
	return *new(svid.SVIDClient)
}

func (c *Client) AgentClient() agent.AgentClient {
	_ = "STUB: not implemented"
	return *new(agent.AgentClient)
}

func (c *Client) DebugClient() debug.DebugClient {
	_ = "STUB: not implemented"
	return *new(debug.DebugClient)
}

func (c *Client) TrustDomainClient() trustdomain.TrustDomainClient {
	_ = "STUB: not implemented"
	return *new(trustdomain.TrustDomainClient)
}

// Open a client ON THE SPIRE-SERVER container
// Used for creating join tokens
type LocalServerClient struct {
	connection *grpc.ClientConn
}

func (c *LocalServerClient) AgentClient() agent.AgentClient {
	_ = "STUB: not implemented"
	return *new(agent.AgentClient)
}

func (c *LocalServerClient) BundleClient() bundle.BundleClient {
	_ = "STUB: not implemented"
	return *new(bundle.BundleClient)
}

func (c *LocalServerClient) EntryClient() entry.EntryClient {
	_ = "STUB: not implemented"
	return *new(entry.EntryClient)
}

func (c *LocalServerClient) Release() { _ = "STUB: not implemented"; return }

func NewLocalServerClient() *LocalServerClient { _ = "STUB: not implemented"; return nil }

type logger struct{}

func (l *logger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }
