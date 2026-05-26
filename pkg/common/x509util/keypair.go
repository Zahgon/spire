package x509util

import (
	"context"
	"crypto"
	"crypto/x509"
)

type Keypair interface {
	// GetCertificate returns the keypair certificate. It is called for each
	// signing request.
	GetCertificate(ctx context.Context) (*x509.Certificate, error)

	// CreateCertificate signs a certificate with the keypair.
	CreateCertificate(ctx context.Context, template *x509.Certificate, publicKey any) (certDER []byte, err error)
}

type MemoryKeypair struct {
	cert *x509.Certificate
	key  crypto.PrivateKey
}

func NewMemoryKeypair(cert *x509.Certificate, key crypto.PrivateKey) *MemoryKeypair {
	_ = "STUB: not implemented"
	return nil
}

func (m *MemoryKeypair) GetCertificate(_ context.Context) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MemoryKeypair) CreateCertificate(_ context.Context, template *x509.Certificate, publicKey any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
