package client

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"google.golang.org/grpc"
)

const (
	defaultDialTimeout      = 30 * time.Second
	roundRobinServiceConfig = `{ "loadBalancingConfig": [ { "round_robin": {} } ] }`
)

type ServerClientConfig struct {
	// Address is the SPIRE server address
	Address string

	TrustDomain spiffeid.TrustDomain

	// GetBundle is a required callback that returns the current trust bundle
	// for used to authenticate the server certificate.
	GetBundle func() []*x509.Certificate

	// GetAgentCertificate is an optional callback used to return the agent
	// certificate to present to the server during the TLS handshake.
	GetAgentCertificate func() *tls.Certificate

	// TLSPolicy determines the post-quantum-safe policy to apply to all TLS connections.
	TLSPolicy tlspolicy.Policy

	// dialOpts are optional gRPC dial options
	dialOpts []grpc.DialOption
}

func NewServerGRPCClient(config ServerClientConfig) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bundleSource struct {
	td     spiffeid.TrustDomain
	getter func() []*x509.Certificate
}

func newBundleSource(td spiffeid.TrustDomain, getter func() []*x509.Certificate) x509bundle.Source {
	_ = "STUB: not implemented"
	return *new(x509bundle.Source)
}

func (s *bundleSource) GetX509BundleForTrustDomain(trustDomain spiffeid.TrustDomain) (*x509bundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type x509SVIDSource struct {
	getter func() *tls.Certificate
}

func newX509SVIDSource(getter func() *tls.Certificate) x509svid.Source {
	_ = "STUB: not implemented"
	return *new(x509svid.Source)
}

func (s *x509SVIDSource) GetX509SVID() (*x509svid.SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
