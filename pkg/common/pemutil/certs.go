package pemutil

import (
	"bytes"
	"crypto/x509"
)

func ParseCertificate(pemBytes []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadCertificate(path string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCertificates(pemBytes []byte) (certs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadCertificates(path string) (certs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeCertificates(certs []*x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func EncodeCertificate(cert *x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func certFromObject(object any) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certsFromBlocks(blocks []Block) (certs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeCertificate(buf *bytes.Buffer, cert *x509.Certificate) {
	_ = "STUB: not implemented"
	// encoding to a memory buffer should not error out
	return
}
