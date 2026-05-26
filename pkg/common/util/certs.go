package util

import (
	"crypto/x509"
)

// NewCertPool creates a new *x509.CertPool based on the certificates given
// as parameters.
func NewCertPool(certs ...*x509.Certificate) *x509.CertPool { _ = "STUB: not implemented"; return nil }

// LoadCertPool loads one or more certificates into an *x509.CertPool from
// a PEM file on disk.
func LoadCertPool(path string) (*x509.CertPool, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadCertificates loads one or more certificates into an []*x509.Certificate from
// a PEM file on disk.
func LoadCertificates(path string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
