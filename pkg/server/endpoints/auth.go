package endpoints

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"github.com/spiffe/spire/pkg/server/svid"
)

var (
	misconfigLogMtx   sync.Mutex
	misconfigLogTimes = make(map[spiffeid.TrustDomain]time.Time)
	misconfigClk      = clock.New()
)

const misconfigLogEvery = time.Minute

// shouldLogFederationMisconfiguration returns true if the last time a misconfiguration
// was logged was more than misconfigLogEvery ago.
func shouldLogFederationMisconfiguration(td spiffeid.TrustDomain) bool {
	_ = "STUB: not implemented"
	return false
}

// bundleGetter fetches the bundle for the given trust domain and parse it as x509 certificates.
func (e *Endpoints) bundleGetter(ctx context.Context, td spiffeid.TrustDomain) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// serverSpiffeVerificationFunc returns a function that is used for peer certificate verification on TLS connections.
// The returned function will verify that the peer certificate is valid, and apply a custom authorization with matchMemberOrOneOf.
// If the peer certificate is not provided, the function will not make any verification and return nil.
func (e *Endpoints) serverSpiffeVerificationFunc(bundleSource x509bundle.Source) func(_ [][]byte, _ [][]*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// matchMemberOrOneOf is a custom spiffeid.Matcher which will validate that the peerSpiffeID belongs to the server
// trust domain or if it is included in the admin_ids configuration permissive list.
func matchMemberOrOneOf(trustDomain spiffeid.TrustDomain, adminIds ...spiffeid.ID) spiffeid.Matcher {
	_ = "STUB: not implemented"
	return *new(spiffeid.Matcher)
}

type x509SVIDSource struct {
	getter func() svid.State
}

func newX509SVIDSource(getter func() svid.State) x509svid.Source {
	_ = "STUB: not implemented"
	return *new(x509svid.Source)
}

func (xs *x509SVIDSource) GetX509SVID() (*x509svid.SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bundleSource struct {
	getter func(spiffeid.TrustDomain) ([]*x509.Certificate, error)
}

func newBundleSource(getter func(spiffeid.TrustDomain) ([]*x509.Certificate, error)) x509bundle.Source {
	_ = "STUB: not implemented"
	return *new(x509bundle.Source)
}

func (bs *bundleSource) GetX509BundleForTrustDomain(trustDomain spiffeid.TrustDomain) (*x509bundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
