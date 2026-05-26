package tpmdevid

import (
	//nolint: gosec // SHA1 use is according to specification
	"crypto/x509"
)

func buildSelectorValues(leaf *x509.Certificate, chains [][]*x509.Certificate) []string {
	_ = "STUB: not implemented"
	return nil
}

// Used to avoid duplicating selectors.

// Iterate over all the certs in the chain (skip leaf at the 0 index)

// If the same fingerprint is generated, continue with the next certificate, because
// a selector should have been already created for it.

func Fingerprint(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

//nolint: gosec // SHA1 use is according to specification
