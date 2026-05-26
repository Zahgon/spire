package spiretest

import (
	"crypto"
	"crypto/x509"
	"testing"

	"github.com/spiffe/spire/pkg/common/pemutil"
)

var (
	EC256Key, _ = pemutil.ParseSigner([]byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcyW+Ne33t4e7HVxn
5aWdL02CcurRNixGgu1vVqQzq3+hRANCAASSQSfkTYd3+u8JEMJUw2Pd143QAOKP
24lWY34SXQInPaja544bc67U0dG0YCNozyAtZxIHFjV+t2HGThM8qNYg
-----END PRIVATE KEY-----
`))
	DefaultKey = EC256Key
)

func SelfSignCertificate(tb testing.TB, tmpl *x509.Certificate) (*x509.Certificate, crypto.Signer) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer)
}

func SelfSignCertificateWithKey(tb testing.TB, tmpl *x509.Certificate, key crypto.Signer) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func CreateCertificate(tb testing.TB, tmpl, parent *x509.Certificate, publicKey, privateKey any) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}
