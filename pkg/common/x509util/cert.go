package x509util

import (
	"crypto"
	"crypto/x509"
)

const (
	unknownAuthorityErr = "x509: certificate signed by unknown authority"
)

func CreateCertificate(template, parent *x509.Certificate, publicKey, privateKey any) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CertificateMatchesPrivateKey(certificate *x509.Certificate, privateKey crypto.PrivateKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DedupeCertificates(bundles ...[]*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

// Retain ordering for easier testing

func DERFromCertificates(certs []*x509.Certificate) (derBytes []byte) {
	_ = "STUB: not implemented"
	return nil
}

// RawCertsToCertificates parses certificates from the given slice of ASN.1 DER data
func RawCertsToCertificates(rawCerts [][]byte) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RawCertsFromCertificates parses ASN.1 DER data from given slice of X.509 Certificates
func RawCertsFromCertificates(certs []*x509.Certificate) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// IsUnknownAuthorityError returns true if the Server returned an unknown authority error when verifying
// presented SVID
func IsUnknownAuthorityError(err error) bool { _ = "STUB: not implemented"; return false }

// Since it is an rpc error we are unable to use errors.As since it is not possible to unwrap

// IsSignedByRoot checks if the provided certificate chain is signed by one of the specified root CAs.
func IsSignedByRoot(chain []*x509.Certificate, rootCAs []*x509.Certificate) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Verify certificate chain, using tainted authorities as root
