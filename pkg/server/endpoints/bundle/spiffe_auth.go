package bundle

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
)

func SPIFFEAuth(getter func() ([]*x509.Certificate, crypto.PrivateKey, error)) ServerAuth {
	_ = "STUB: not implemented"
	return *new(ServerAuth)
}

type spiffeAuth struct {
	getter func() ([]*x509.Certificate, crypto.PrivateKey, error)
}

func (s *spiffeAuth) GetTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

func (s *spiffeAuth) getCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
