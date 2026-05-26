package spireplugin

import (
	"context"
	"crypto/x509"
	"net"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/logger"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"google.golang.org/grpc"
)

// newServerClient creates a new spire-server client
func newServerClient(serverID spiffeid.ID, serverAddr string, workloadAPIAddr net.Addr, log hclog.Logger, tlsPolicy tlspolicy.Policy) *serverClient {
	_ = "STUB: not implemented"
	return nil
}

type serverClient struct {
	serverID        spiffeid.ID
	conn            *grpc.ClientConn
	serverAddr      string
	workloadAPIAddr net.Addr
	log             logger.Logger
	tlsPolicy       tlspolicy.Policy

	mtx    sync.RWMutex
	source *workloadapi.X509Source

	bundleClient bundlev1.BundleClient
	svidClient   svidv1.SVIDClient
}

// start initializes spire-server endpoints client, it uses X509 source to keep an active connection
func (c *serverClient) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close active connection

// Update connection and source

// release releases the connection to SPIRE server and cleans clients
func (c *serverClient) release() { _ = "STUB: not implemented"; return }

// newDownstreamX509CA requests new downstream CAs to server
func (c *serverClient) newDownstreamX509CA(ctx context.Context, csr []byte, preferredTTL int32) ([]*x509.Certificate, []*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// parse authorities to verify that are valid X509 certificates

// parse cert chains to verify that are valid X509 certificates

// newDownstreamX509CA publishes a JWT key to the server
func (c *serverClient) publishJWTAuthority(ctx context.Context, key *types.JWTKey) ([]*types.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBundle gets the bundle for the trust domain of the server
func (c *serverClient) getBundle(ctx context.Context) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type logAdapter struct {
	log hclog.Logger
}

func (l *logAdapter) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logAdapter) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logAdapter) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (l *logAdapter) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }
