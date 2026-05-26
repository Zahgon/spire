package x509util

import (
	//nolint: gosec // usage of SHA1 is according to RFC 5280

	"github.com/spiffe/spire/pkg/common/util"
)

var x509utilsha256skid = util.FIPS140Only()

// GetSubjectKeyID calculates a subject key identifier by doing a hash
// over the ASN.1 encoding of the public key.
func GetSubjectKeyID(pubKey any) ([]byte, error) {
	_ = "STUB: not implemented"
	// Borrowed with love from cfssl under the BSD 2-Clause license.
	return nil, nil
}

// Borrowed with love from Go std lib crypto/x509 under the BSD 3-Clause license.

// SubjectKeyId generated using method 1 in RFC 7093, Section 2:
//    1) The keyIdentifier is composed of the leftmost 160-bits of the
//    SHA-256 hash of the value of the BIT STRING subjectPublicKey
//    (excluding the tag, length, and number of unused bits).

// SubjectKeyId generated using method 1 in RFC 5280, Section 4.2.1.2:
//   (1) The keyIdentifier is composed of the 160-bit SHA-1 hash of the
//   value of the BIT STRING subjectPublicKey (excluding the tag,
//   length, and number of unused bits).
//nolint: gosec // usage of SHA1 is according to RFC 5280

// SubjectKeyIDToString parse Subject Key ID into string
func SubjectKeyIDToString(ski []byte) string { _ = "STUB: not implemented"; return "" }

// Append leading 0 in cases where hexadecimal representation is odd number of characters
// in order to be more consistent with other tooling that displays certificate serial numbers.
