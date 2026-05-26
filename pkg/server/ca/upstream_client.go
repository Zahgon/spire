package ca

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/spiffe/spire/pkg/common/coretypes/x509certificate"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority"
	"github.com/spiffe/spire/proto/spire/common"
)

// BundleUpdater is the interface used by the UpstreamClient to append bundle
// updates.
type BundleUpdater interface {
	SyncX509Roots(ctx context.Context, roots []*x509certificate.X509Authority) error
	AppendJWTKeys(ctx context.Context, keys []*common.PublicKey) ([]*common.PublicKey, error)
	LogError(err error, msg string)
}

// ValidateX509CAFunc is used by the upstream client to validate an X509CA
// newly minted by an upstream authority before it accepts it.
type ValidateX509CAFunc = func(x509CA, x509Roots []*x509.Certificate) error

// UpstreamClientConfig is the configuration for an UpstreamClient. Each field
// is required.
type UpstreamClientConfig struct {
	UpstreamAuthority upstreamauthority.UpstreamAuthority
	BundleUpdater     BundleUpdater
}

// UpstreamClient is used to interact with and stream updates from the
// UpstreamAuthority plugin.
type UpstreamClient struct {
	c UpstreamClientConfig

	mintX509CAMtx                   sync.Mutex
	mintX509CAStream                *streamState
	publishJWTKeyMtx                sync.Mutex
	publishJWTKeyStream             *streamState
	subscribeToLocalBundleStreamMtx sync.Mutex
	subscribeToLocalBundleStream    *streamState
}

// NewUpstreamClient returns a new UpstreamAuthority plugin client.
func NewUpstreamClient(config UpstreamClientConfig) *UpstreamClient {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the client, stopping any open streams against the
// UpstreamAuthority plugin.
func (u *UpstreamClient) Close() error { _ = "STUB: not implemented"; return nil }

// MintX509CA mints an X.509CA using the UpstreamAuthority. It maintains an
// open stream to the UpstreamAuthority plugin to receive and append X.509 root
// updates to the bundle. The stream remains open until another call to
// MintX509CA happens or the client is closed.
func (u *UpstreamClient) MintX509CA(ctx context.Context, csr []byte, ttl time.Duration, validateX509CA ValidateX509CAFunc) (_ []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublishJWTKey publishes the JWT key to the UpstreamAuthority. It maintains
// an open stream to the UpstreamAuthority plugin to receive and append JWT key
// updates to the bundle. The stream remains open until another call to
// PublishJWTKey happens or the client is closed.
func (u *UpstreamClient) PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) (_ []*common.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *UpstreamClient) SubscribeToLocalBundle(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpstreamClient) runMintX509CAStream(ctx context.Context, csr []byte, ttl time.Duration, validateX509CA ValidateX509CAFunc, firstResultCh chan<- mintX509CAResult) {
	_ = "STUB: not implemented"
	return
}

// Extract all root certificates

// Before we append the roots and return the response, we must first
// validate that the minted intermediate can sign a valid, conformant
// X509-SVID chain of trust using the provided callback.

// This is normal if the plugin does not support streaming
// bundle updates.

// This is normal. This client cancels this stream when opening
// a new stream.

func (u *UpstreamClient) runPublishJWTKeyStream(ctx context.Context, jwtKey *common.PublicKey, firstResultCh chan<- publishJWTKeyResult) {
	_ = "STUB: not implemented"
	return
}

// This is normal if the plugin does not support streaming
// bundle updates.

// This is normal. This client cancels this stream when opening
// a new stream.

func (u *UpstreamClient) runSubscribeToLocalBundleStream(ctx context.Context, firstResultCh chan<- bundleUpdatesResult) {
	_ = "STUB: not implemented"
	return
}

// This is normal if the plugin does not support streaming
// bundle updates.

// This is normal. This client cancels this stream when opening
// a new stream.

type mintX509CAResult struct {
	x509CA []*x509.Certificate
	err    error
}

type publishJWTKeyResult struct {
	jwtKeys []*common.PublicKey
	err     error
}

type bundleUpdatesResult struct {
	x509CA  []*x509.Certificate
	jwtKeys []*common.PublicKey
	err     error
}

// streamState manages the state for open streams to the plugin that are
// receiving bundle updates. It is protected by the respective mutexes in
// the UpstreamClient.
type streamState struct {
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	stopOnce *sync.Once
	stopped  chan struct{}
}

func newStreamState() *streamState { _ = "STUB: not implemented"; return nil }

func (s *streamState) Stop() { _ = "STUB: not implemented"; return }

func (s *streamState) Start(fn func(context.Context)) { _ = "STUB: not implemented"; return }

func (s *streamState) WaitUntilStopped(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamState) stop() { _ = "STUB: not implemented"; return }
