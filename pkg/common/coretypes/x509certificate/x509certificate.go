package x509certificate

import (
	"crypto/x509"
)

// TODO: may we call it Authority?
// TODO: may we add subjectKeyID?
type X509Authority struct {
	Certificate *x509.Certificate
	Tainted     bool
}

func fromProtoFields(asn1 []byte, tainted bool) (*X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toProtoFields(x509Authority *X509Authority) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func validateX509Certificate(cert *x509.Certificate) error { _ = "STUB: not implemented"; return nil }
