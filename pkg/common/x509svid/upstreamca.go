package x509svid

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/x509util"
)

const (
	DefaultUpstreamCABackdate = time.Second * 10
	DefaultUpstreamCATTL      = time.Hour
)

type UpstreamCAOptions struct {
	Backdate time.Duration
	Clock    clock.Clock
}

type UpstreamCA struct {
	keypair     x509util.Keypair
	trustDomain spiffeid.TrustDomain
	options     UpstreamCAOptions
}

func NewUpstreamCA(keypair x509util.Keypair, trustDomain spiffeid.TrustDomain, options UpstreamCAOptions) *UpstreamCA {
	_ = "STUB: not implemented"
	return nil
}

func (ca *UpstreamCA) SignCSR(ctx context.Context, csrDER []byte, preferredTTL time.Duration) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the default TTL setting unless a preferred TTL is specified.
