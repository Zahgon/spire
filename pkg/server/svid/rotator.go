package svid

import (
	"context"
	"crypto"
	"crypto/x509"
	"time"

	"github.com/imkira/go-observer"
)

var (
	defaultBundleVerificationTicker = 30 * time.Second
)

type Rotator struct {
	c *RotatorConfig

	state           observer.Property
	isSVIDTainted   bool
	taintedReceived chan bool
}

// State is the current SVID and key
type State struct {
	SVID []*x509.Certificate
	Key  crypto.Signer
}

// Start generates a new SVID and then starts the rotator.
func (r *Rotator) Initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Rotator) State() State { _ = "STUB: not implemented"; return *new(State) }

func (r *Rotator) Subscribe() observer.Stream {
	_ = "STUB: not implemented"
	return *new(observer.Stream)
}

func (r *Rotator) Interval() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (r *Rotator) triggerTaintedReceived(tainted bool) { _ = "STUB: not implemented"; return }

// Run starts a ticker which monitors the server SVID
// for expiration and rotates the SVID as necessary.
func (r *Rotator) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// shouldRotate returns a boolean informing the caller of whether the
// SVID should be rotated.
func (r *Rotator) shouldRotate() bool { _ = "STUB: not implemented"; return false }

func (r *Rotator) isX509AuthorityTainted(taintedAuthorities []*x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify certificate chain, using tainted authority as root

// rotateSVID cuts a new server SVID from the CA plugin and installs
// it on the endpoints struct. Also updates the CA certificates.
func (r *Rotator) rotateSVID(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// New SVID must not be tainted. Rotator is notified about tainted
// authorities only when the intermediate is already rotated.

func certHalfLife(cert *x509.Certificate) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
