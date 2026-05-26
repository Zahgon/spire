package ca

import (
	"crypto"
	"crypto/x509"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/pemutil"
	"github.com/spiffe/spire/pkg/server/credvalidator"
)

var (
	validationPubkey, _ = pemutil.ParsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzLY1/SRlsMJExTnuvzBO292RjGjU
3L8jFRtmQl0CjBeHdxUlGK1OkNLDYh0b6AW4siWt+y+DcbUAWNb14e5zWg==
-----END PUBLIC KEY-----`))
)

type X509CAValidator struct {
	TrustDomain   spiffeid.TrustDomain
	CredValidator *credvalidator.Validator
	Signer        crypto.Signer
	Clock         clock.Clock
}

func (v *X509CAValidator) ValidateUpstreamX509CA(x509CA, upstreamRoots []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *X509CAValidator) ValidateSelfSignedX509CA(x509CA *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *X509CAValidator) validateX509CA(x509CA *x509.Certificate, x509Roots, upstreamChain []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}
