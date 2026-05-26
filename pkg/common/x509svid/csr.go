package x509svid

import (
	"crypto/x509"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

func ParseAndValidateCSR(csrDER []byte, td spiffeid.TrustDomain) (csr *x509.CertificateRequest, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateCSR(csr *x509.CertificateRequest, td spiffeid.TrustDomain) error {
	_ = "STUB: not implemented"
	return nil
}
