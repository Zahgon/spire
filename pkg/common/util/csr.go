package util

import (
	"crypto/x509"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

func MakeCSR(privateKey any, spiffeID spiffeid.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeCSRWithoutURISAN(privateKey any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeCSR(privateKey any, template *x509.CertificateRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
