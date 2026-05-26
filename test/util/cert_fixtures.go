package util

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"path"
)

var (
	svidPath        = path.Join(ProjectRoot(), "test/fixture/certs/svid.pem")
	svidKeyPath     = path.Join(ProjectRoot(), "test/fixture/certs/svid_key.pem")
	caPath          = path.Join(ProjectRoot(), "test/fixture/certs/ca.pem")
	caKeyPath       = path.Join(ProjectRoot(), "test/fixture/certs/ca_key.pem")
	bundlePath      = path.Join(ProjectRoot(), "test/fixture/certs/bundle.der")
	largeBundlePath = path.Join(ProjectRoot(), "test/fixture/certs/large_bundle.der")
)

// LoadCAFixture reads, parses, and returns the pre-defined CA fixture and key
func LoadCAFixture() (ca *x509.Certificate, key *ecdsa.PrivateKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// LoadSVIDFixture reads, parses, and returns the pre-defined SVID fixture and key
func LoadSVIDFixture() (svid *x509.Certificate, key *ecdsa.PrivateKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func LoadBundleFixture() ([]*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadLargeBundleFixture() ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadCertAndKey reads and parses both a certificate and a private key at once
func LoadCertAndKey(crtPath, keyPath string) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// LoadCert reads and parses an X.509 certificate at the specified path
func LoadCert(path string) (*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadKey reads and parses the ECDSA private key at the specified path
func LoadKey(path string) (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadPEM reads and parses the PEM structure at the specified path
func LoadPEM(path string) (*pem.Block, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadBundle(path string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
