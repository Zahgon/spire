package pemutil

import (
	"crypto/x509"
)

func ParseCertificateRequest(pemBytes []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadCertificateRequest(path string) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func csrFromObject(object any) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
