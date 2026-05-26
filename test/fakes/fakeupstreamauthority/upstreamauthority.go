package fakeupstreamauthority

import (
	"context"
	"crypto"
	"crypto/x509"
	"sync"
	"testing"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	"github.com/spiffe/spire/pkg/common/coretypes/x509certificate"
	"github.com/spiffe/spire/pkg/common/x509svid"
	"github.com/spiffe/spire/proto/spire/common"
	"github.com/spiffe/spire/test/clock"
	"github.com/spiffe/spire/test/testkey"
)

var (
	x509RootKey = testkey.MustEC256()
	x509IntKey  = testkey.MustEC256()
)

type Config struct {
	Clock                       clock.Clock
	TrustDomain                 spiffeid.TrustDomain
	UseIntermediate             bool
	DisallowPublishJWTKey       bool
	UseSubscribeToLocalBundle   bool
	KeyUsage                    x509.KeyUsage
	MutateMintX509CAResponse    func(*upstreamauthorityv1.MintX509CAResponse)
	MutatePublishJWTKeyResponse func(*upstreamauthorityv1.PublishJWTKeyResponse)
}

type UpstreamAuthority struct {
	upstreamauthorityv1.UnimplementedUpstreamAuthorityServer

	t      *testing.T
	config Config

	x509CAMtx        sync.RWMutex
	x509CA           *x509svid.UpstreamCA
	x509CASN         int64
	x509Root         *x509certificate.X509Authority
	x509Intermediate *x509.Certificate
	x509Roots        []*x509certificate.X509Authority

	jwtKeysMtx sync.RWMutex
	jwtKeys    []*common.PublicKey

	streamsMtx           sync.Mutex
	mintX509CAStreams    map[chan struct{}]struct{}
	publishJWTKeyStreams map[chan struct{}]struct{}
}

func New(t *testing.T, config Config) *UpstreamAuthority { _ = "STUB: not implemented"; return nil }

func (ua *UpstreamAuthority) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) PublishJWTKeyAndSubscribe(req *upstreamauthorityv1.PublishJWTKeyRequest, stream upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a first update on the stream, as required.

func (ua *UpstreamAuthority) RotateX509CA() { _ = "STUB: not implemented"; return }

func (ua *UpstreamAuthority) TaintAuthority(index int) error { _ = "STUB: not implemented"; return nil }

func (ua *UpstreamAuthority) X509Root() *x509certificate.X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) X509Roots() []*x509certificate.X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) X509Intermediate() *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) JWTKeys() []*common.PublicKey { _ = "STUB: not implemented"; return nil }

func (ua *UpstreamAuthority) AppendJWTKey(jwtKey *common.PublicKey) {
	_ = "STUB: not implemented"
	return
}

func (ua *UpstreamAuthority) TriggerX509RootsChanged() { _ = "STUB: not implemented"; return }

func (ua *UpstreamAuthority) TriggerJWTKeysChanged() { _ = "STUB: not implemented"; return }

func (ua *UpstreamAuthority) newMintX509CAStream() chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) removeMintX509CAStream(streamCh chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (ua *UpstreamAuthority) mintX509CA(ctx context.Context, csr []byte, preferredTTL time.Duration) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ua *UpstreamAuthority) sendMintX509CAResponse(stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer, resp *upstreamauthorityv1.MintX509CAResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) newPublishJWTKeyStream() chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) removePublishJWTKeyStream(streamCh chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (ua *UpstreamAuthority) sendPublishJWTKeyStream(stream upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer, resp *upstreamauthorityv1.PublishJWTKeyResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (ua *UpstreamAuthority) createRootCertificate() { _ = "STUB: not implemented"; return }

func (ua *UpstreamAuthority) createIntermediateCertificate() { _ = "STUB: not implemented"; return }

func (ua *UpstreamAuthority) nextX509CASN() int64 { _ = "STUB: not implemented"; return 0 }

func createCATemplate(now time.Time, cn string, sn int64, keyUsage x509.KeyUsage) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func createCertificate(t *testing.T, template, parent *x509.Certificate, publicKey crypto.PublicKey, privateKey crypto.PrivateKey) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}
