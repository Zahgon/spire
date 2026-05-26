package svid

import (
	"crypto/x509"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/client"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/common/rotationutil"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
)

const DefaultRotatorInterval = 5 * time.Second

type RotatorConfig struct {
	SVIDKeyManager keymanager.SVIDKeyManager
	Log            logrus.FieldLogger
	Metrics        telemetry.Metrics
	TrustDomain    spiffeid.TrustDomain
	ServerAddr     string
	NodeAttestor   nodeattestor.NodeAttestor
	Reattestable   bool

	// Initial SVID and key
	SVID    []*x509.Certificate
	SVIDKey keymanager.Key

	BundleStream *cache.BundleStream

	// How long to wait between expiry checks
	Interval time.Duration

	// Clk is the clock that the rotator will use to create a ticker
	Clk clock.Clock

	RotationStrategy *rotationutil.RotationStrategy

	// TLSPolicy determines the post-quantum-safe policy for TLS connections.
	TLSPolicy tlspolicy.Policy
}

func NewRotator(c *RotatorConfig) (Rotator, client.Client) {
	_ = "STUB: not implemented"
	return *new(Rotator), *new(client.Client)
}

func newRotator(c *RotatorConfig) (*rotator, client.Client) {
	_ = "STUB: not implemented"
	return nil, *new(client.Client)
}
