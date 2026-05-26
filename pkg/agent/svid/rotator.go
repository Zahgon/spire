package svid

import (
	"context"
	"crypto"
	"crypto/x509"
	"sync"

	"github.com/andres-erbsen/clock"
	"github.com/imkira/go-observer"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/spire/pkg/agent/client"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/common/backoff"
	"google.golang.org/grpc"
)

type Rotator interface {
	Run(ctx context.Context) error
	Reattest(ctx context.Context) error
	// NotifyTaintedAuthorities processes new tainted authorities. If the current SVID is compromised,
	// it is marked to force rotation.
	NotifyTaintedAuthorities([]*x509.Certificate) error
	IsTainted() bool

	State() State
	Subscribe() observer.Stream
	GetRotationMtx() *sync.RWMutex
	SetRotationFinishedHook(func())
}

type Client interface {
	RenewSVID(ctx context.Context, csr []byte) (*client.X509SVID, error)
	Release()
}

type rotator struct {
	c      *RotatorConfig
	client Client

	state observer.Property
	clk   clock.Clock

	// backoff calculator for rotation check interval, backing off if error is returned on
	// rotation attempt
	backoff backoff.BackOff

	// Mutex used to protect access to c.BundleStream.
	bsm *sync.RWMutex

	// Mutex used to prevent rotations when a new connection is being created
	rotMtx *sync.RWMutex

	hooks struct {
		// Hook that will be called when the SVID rotation finishes
		rotationFinishedHook func()

		// Hook that is called when the rotator starts running
		runRotatorSignal chan struct{}
	}
	tainted bool
}

type State struct {
	SVID         []*x509.Certificate
	Key          crypto.Signer
	Reattestable bool
}

// Run runs the rotator. It monitors the server SVID for expiration and rotates
// as necessary. It also watches for changes to the trust bundle.
func (r *rotator) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *rotator) runRotation(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Since our X509 cert has expired, and we weren't able to carry out a rotation request, we're probably unrecoverable without re-attesting.

// Just log the error and wait for next rotation

func (r *rotator) processBundleUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rotator) State() State { _ = "STUB: not implemented"; return *new(State) }

func (r *rotator) Subscribe() observer.Stream {
	_ = "STUB: not implemented"
	return *new(observer.Stream)
}

func (r *rotator) IsTainted() bool { _ = "STUB: not implemented"; return false }

func (r *rotator) setTainted(tainted bool) { _ = "STUB: not implemented"; return }

func (r *rotator) NotifyTaintedAuthorities(taintedAuthorities []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rotator) GetRotationMtx() *sync.RWMutex { _ = "STUB: not implemented"; return nil }

func (r *rotator) SetRotationFinishedHook(f func()) { _ = "STUB: not implemented"; return }

func (r *rotator) Reattest(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *rotator) rotateSVIDIfNeeded(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// reattest goes through the full attestation process with the server and gets a new SVID.
func (r *rotator) reattest(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// Get the mtx before starting the reattestation
// In this way, the client do not create new connections until the new SVID is received

// We must release the client because its underlying connection is tied to an
// expired SVID, so next time the client is used, it will get a new connection with
// the most up-to-date SVID.

// rotateSVID asks SPIRE's server for a new agent's SVID.
func (r *rotator) rotateSVID(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Get the mtx before starting the rotation
// In this way, the client do not create new connections until the new SVID is received

// We must release the client because its underlying connection is tied to an
// expired SVID, so next time the client is used, it will get a new connection with
// the most up-to-date SVID.

func (r *rotator) getBundle() (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rotator) generateKey(ctx context.Context) (keymanager.Key, error) {
	_ = "STUB: not implemented"
	return *new(keymanager.Key), nil
}

func (r *rotator) serverConn(bundle *spiffebundle.Bundle) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rotationError(state State) string { _ = "STUB: not implemented"; return "" }
