package util

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"math/big"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/test/clock"
)

// NewCSRTemplate returns a default CSR template with the specified SPIFFE ID.
func NewCSRTemplate(spiffeID string) ([]byte, crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.PublicKey), nil
}

func NewCSRTemplateWithKey(spiffeID string, key crypto.Signer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSVIDTemplate returns a default SVID template with the specified SPIFFE ID. Must
// be signed before it's valid.
func NewSVIDTemplate(clk clock.Clock, spiffeID string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCATemplate returns a default CA template with the specified trust domain. Must
// be signed before it's valid.
func NewCATemplate(clk clock.Clock, trustDomain spiffeid.TrustDomain) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelfSign creates a new self-signed certificate with the provided template.
func SelfSign(req *x509.Certificate) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil,

		// Sign creates a new certificate based on the provided template and signed using parent
		// certificate and signerPrivateKey.
		nil, nil
}

func Sign(req, parent *x509.Certificate, signerPrivateKey any) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Returns an SVID template with many default values set. Should be overwritten prior to
// generating a new test SVID
func defaultSVIDTemplate(clk clock.Clock) *x509.Certificate { _ = "STUB: not implemented"; return nil }

// Returns an CA template with many default values set.
func defaultCATemplate(clk clock.Clock) *x509.Certificate { _ = "STUB: not implemented"; return nil }

// Create an x509 extension with the URI SAN of the given SPIFFE ID, and set it onto
// the referenced certificate
func addSpiffeExtension(spiffeID string, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// Creates a random certificate serial number
func randomSerial() *big.Int { _ = "STUB: not implemented"; return nil }
