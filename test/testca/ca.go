package testca

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/spiffe/go-spiffe/v2/bundle/jwtbundle"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/jwtsvid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
)

type CA struct {
	tb     testing.TB
	td     spiffeid.TrustDomain
	parent *CA
	cert   *x509.Certificate
	key    crypto.Signer
	jwtKey crypto.Signer
	jwtKid string
}

type CertificateOption interface {
	apply(*x509.Certificate)
}

type certificateOption func(*x509.Certificate)

func (co certificateOption) apply(c *x509.Certificate) { _ = "STUB: not implemented"; return }

func New(tb testing.TB, td spiffeid.TrustDomain) *CA { _ = "STUB: not implemented"; return nil }

func (ca *CA) ChildCA(options ...CertificateOption) *CA { _ = "STUB: not implemented"; return nil }

func (ca *CA) CreateX509SVID(id spiffeid.ID, options ...CertificateOption) *x509svid.SVID {
	_ = "STUB: not implemented"
	return nil
}

func (ca *CA) CreateX509Certificate(options ...CertificateOption) ([]*x509.Certificate, crypto.Signer) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer)
}

func (ca *CA) CreateJWTSVID(id spiffeid.ID, audience []string) *jwtsvid.SVID {
	_ = "STUB: not implemented"
	return nil
}

func (ca *CA) X509Authorities() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func (ca *CA) JWTAuthorities() map[string]crypto.PublicKey { _ = "STUB: not implemented"; return nil }

func (ca *CA) Bundle() *spiffebundle.Bundle { _ = "STUB: not implemented"; return nil }

func (ca *CA) X509Bundle() *x509bundle.Bundle { _ = "STUB: not implemented"; return nil }

func (ca *CA) JWTBundle() *jwtbundle.Bundle { _ = "STUB: not implemented"; return nil }

func (ca *CA) GetSubjectKeyID() string { _ = "STUB: not implemented"; return "" }

func (ca *CA) GetUpstreamAuthorityID() string { _ = "STUB: not implemented"; return "" }

func (ca *CA) chain(includeRoot bool) []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func CreateCACertificate(tb testing.TB, parent *x509.Certificate, parentKey crypto.Signer, options ...CertificateOption) (*x509.Certificate, crypto.Signer) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer)
}

func CreateX509Certificate(tb testing.TB, parent *x509.Certificate, parentKey crypto.Signer, options ...CertificateOption) (*x509.Certificate, crypto.Signer) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer)
}

func CreateX509SVID(tb testing.TB, parent *x509.Certificate, parentKey crypto.Signer, id spiffeid.ID, options ...CertificateOption) (*x509.Certificate, crypto.Signer) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer)
}

func CreateCertificate(tb testing.TB, tmpl, parent *x509.Certificate, publicKey, privateKey any) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func newSerial(tb testing.TB) *big.Int { _ = "STUB: not implemented"; return nil }

func WithSerial(serial *big.Int) CertificateOption {
	_ = "STUB: not implemented"
	return *new(CertificateOption)
}

func WithKeyUsage(keyUsage x509.KeyUsage) CertificateOption {
	_ = "STUB: not implemented"
	return *new(CertificateOption)
}

func WithLifetime(notBefore, notAfter time.Time) CertificateOption {
	_ = "STUB: not implemented"
	return *new(CertificateOption)
}

func WithID(id spiffeid.ID) CertificateOption {
	_ = "STUB: not implemented"
	return *new(CertificateOption)
}

func WithSubject(subject pkix.Name) CertificateOption {
	_ = "STUB: not implemented"
	return *new(CertificateOption)
}

func applyOptions(c *x509.Certificate, options ...CertificateOption) {
	_ = "STUB: not implemented"
	return
}

// newKeyID returns a random id useful for identifying keys
func newKeyID(tb testing.TB) string { _ = "STUB: not implemented"; return "" }

func keyIDFromBytes(choices []byte) string { _ = "STUB: not implemented"; return "" }
