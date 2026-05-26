package bundle

import (
	"crypto/x509"
	"io"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

const (
	headerFmt = `****************************************
* %s
****************************************
`
)

// loadParamData loads the data from a parameter. If the parameter is empty then
// data is ready from "in", otherwise the parameter is used as a filename to
// read file contents.
func loadParamData(in io.Reader, fn string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// printX509Authorities print provided certificates into writer
func printX509Authorities(out io.Writer, certs []*types.X509Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// printCACertsPEM encodes DER certificates to PEM format and print using writer
func printCACertsPEM(out io.Writer, caCerts []byte) error { _ = "STUB: not implemented"; return nil }

// printBundle marshals and prints the bundle using the provided writer
func printBundle(out io.Writer, bundle *types.Bundle) error { _ = "STUB: not implemented"; return nil }

// bundleFromProto converts a bundle from the given *types.Bundle to *spiffebundle.Bundle
func bundleFromProto(bundleProto *types.Bundle) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// x509CertificatesFromProto converts X.509 certificates from the given []*types.X509Certificate to []*x509.Certificate
func x509CertificatesFromProto(proto []*types.X509Certificate) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printBundleWithFormat(out io.Writer, bundle *types.Bundle, format string, header bool) error {
	_ = "STUB: not implemented"
	return nil
}

// validateFormat validates that the provided format is a valid format.
// If no format is provided, the default format is returned
func validateFormat(format string) (string, error) { _ = "STUB: not implemented"; return "", nil }
