package attestor

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/agent/storage"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"google.golang.org/grpc"
)

const (
	roundRobinServiceConfig = `{ "loadBalancingConfig": [ { "round_robin": {} } ] }`
)

type AttestationResult struct {
	SVID         []*x509.Certificate
	Key          keymanager.Key
	Bundle       *spiffebundle.Bundle
	Reattestable bool
}

type Attestor interface {
	Attest(ctx context.Context) (*AttestationResult, error)
}

type Config struct {
	Catalog              catalog.Catalog
	Metrics              telemetry.Metrics
	JoinToken            string
	TrustDomain          spiffeid.TrustDomain
	BootstrapTrustBundle []*x509.Certificate
	InsecureBootstrap    bool
	Storage              storage.Storage
	Log                  logrus.FieldLogger
	ServerAddress        string
	NodeAttestor         nodeattestor.NodeAttestor
	TLSPolicy            tlspolicy.Policy
}

type attestor struct {
	c *Config
}

func New(config *Config) Attestor { _ = "STUB: not implemented"; return *new(Attestor) }

func (a *attestor) Attest(ctx context.Context) (res *AttestationResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is a bizarre case where we have an SVID but were unable to
// load a bundle from the cache which suggests some tampering with the
// cache on disk.

// Load the current SVID and key. The returned SVID is nil to indicate a new SVID should be created.
func (a *attestor) loadSVID(ctx context.Context) ([]*x509.Certificate, keymanager.Key, bool, error) {
	_ = "STUB: not implemented"
	return nil, *new(keymanager.Key), false, nil
}

// Neither private key nor SVID were found.

// IsSVIDExpired returns true if the X.509 SVID provided is expired
func IsSVIDExpired(svid []*x509.Certificate, timeNow func() time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *attestor) loadBundle() (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *attestor) getBundle(ctx context.Context, conn *grpc.ClientConn) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *attestor) getSVID(ctx context.Context, conn *grpc.ClientConn, csr []byte, attestor nodeattestor.NodeAttestor) ([]*x509.Certificate, bool, error) {
	_ = "STUB: not implemented"
	// make sure all the streams are cancelled if something goes awry
	return nil, false, nil
}

// Read agent SVID from data dir. If an error is encountered, it will be logged and `nil`
// will be returned.
func (a *attestor) readSVIDFromDisk() ([]*x509.Certificate, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// newSVID obtains an agent svid for the given private key by performing node attesatation. The bundle is
// necessary in order to validate the SPIRE server we are attesting to. Returns the SVID and an updated bundle.
func (a *attestor) newSVID(ctx context.Context, key keymanager.Key, bundle *spiffebundle.Bundle) (_ []*x509.Certificate, _ *spiffebundle.Bundle, _ bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
}

func (a *attestor) serverConn(bundle *spiffebundle.Bundle) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We shouldn't get here since loadBundle() should fail if the bundle
// is empty, but just in case...

// Insecure bootstrapping. Do not verify the server chain but rather do a
// simple, soft verification that the server URI matches the expected
// SPIFFE ID. This is not a security feature but rather a check that we've
// reached what appears to be the right trust domain server.

//nolint: gosec // this is required in order to do non-hostname based verification
//nolint:gosec // we don't need SessionTicketsDisabled

// This is not really possible without a catastrophic bug
// creeping into the TLS stack.

type ServerStream struct {
	Client       agentv1.AgentClient
	Csr          []byte
	Log          logrus.FieldLogger
	SVID         []*x509.Certificate
	Reattestable bool
	stream       agentv1.Agent_AttestAgentClient
}

func (ss *ServerStream) SendAttestationData(ctx context.Context, attestationData nodeattestor.AttestationData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss *ServerStream) SendChallengeResponse(ctx context.Context, response []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss *ServerStream) sendRequest(ctx context.Context, req *agentv1.AttestAgentRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findKeyForSVID(keys []keymanager.Key, svid []*x509.Certificate) (keymanager.Key, bool) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), false
}

func getSVIDFromAttestAgentResponse(r *agentv1.AttestAgentResponse) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
