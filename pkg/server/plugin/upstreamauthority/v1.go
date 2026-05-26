package upstreamauthority

import (
	"context"
	"crypto/x509"
	"time"

	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	"github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/pkg/common/coretypes/x509certificate"
	"github.com/spiffe/spire/pkg/common/plugin"
	"github.com/spiffe/spire/proto/spire/common"
)

type V1 struct {
	plugin.Facade
	upstreamauthorityv1.UpstreamAuthorityPluginClient
}

// MintX509CA provides the V1 implementation of the UpstreamAuthority
// interface method of the same name.
func (v1 *V1) MintX509CA(ctx context.Context, csr []byte, preferredTTL time.Duration) (_ []*x509.Certificate, _ []*x509certificate.X509Authority, _ UpstreamX509AuthorityStream, err error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(UpstreamX509AuthorityStream), nil
}

// Only cancel the context if the function fails. Otherwise, the
// returned stream will be in charge of cancellation.

// TODO: may we add a new type to get upstream authority with metadata?

// PublishJWTKey provides the V1 implementation of the UpstreamAuthority
// interface method of the same name.
func (v1 *V1) PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) (_ []*common.PublicKey, _ UpstreamJWTAuthorityStream, err error) {
	_ = "STUB: not implemented"
	return nil, *new(UpstreamJWTAuthorityStream), nil
}

// Only cancel the context if the function fails. Otherwise, the
// returned stream will be in charge of cancellation.

func (v1 *V1) SubscribeToLocalBundle(ctx context.Context) (_ []*x509certificate.X509Authority, _ []*common.PublicKey, _ LocalBundleUpdateStream, err error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(LocalBundleUpdateStream), nil
}

// Only cancel the context if the function fails. Otherwise, the
// returned stream will be in charge of cancellation.

func (v1 *V1) parseMintX509CAFirstResponse(resp *upstreamauthorityv1.MintX509CAResponse) ([]*x509.Certificate, []*x509certificate.X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (v1 *V1) parseMintX509CABundleUpdate(resp *upstreamauthorityv1.MintX509CAResponse) ([]*x509certificate.X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v1 *V1) parseX509Authorities(rawX509Authorities []*types.X509Certificate) ([]*x509certificate.X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v1 *V1) streamError(err error) error { _ = "STUB: not implemented"; return nil }

func (v1 *V1) toCommonProtos(pbs []*types.JWTKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type v1UpstreamX509AuthorityStream struct {
	v1     *V1
	stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeClient
	cancel context.CancelFunc
}

func (s *v1UpstreamX509AuthorityStream) RecvUpstreamX509Authorities() ([]*x509certificate.X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is expected if the plugin does not support streaming
// authority updates.

func (s *v1UpstreamX509AuthorityStream) Close() { _ = "STUB: not implemented"; return }

type v1UpstreamJWTAuthorityStream struct {
	v1     *V1
	stream upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeClient
	cancel context.CancelFunc
}

func (s *v1UpstreamJWTAuthorityStream) RecvUpstreamJWTAuthorities() ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is expected if the plugin does not support streaming
// authority updates.

func (s *v1UpstreamJWTAuthorityStream) Close() { _ = "STUB: not implemented"; return }

type v1LocalBundleStream struct {
	v1     *V1
	stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleClient
	cancel context.CancelFunc
}

func (s *v1LocalBundleStream) RecvLocalBundleUpdate() ([]*x509certificate.X509Authority, []*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// This is expected if the plugin does not support streaming
// authority updates.

func (s *v1LocalBundleStream) Close() { _ = "STUB: not implemented"; return }
