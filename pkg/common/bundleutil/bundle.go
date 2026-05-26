package bundleutil

import (
	"crypto"
	"crypto/x509"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func CommonBundleFromProto(b *types.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SPIFFEBundleToProto(b *spiffebundle.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* TODO: Parse WIT authorities once go-spiffe adds support */

func SPIFFEBundleFromProto(b *common.Bundle) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* TODO: Set WIT authorities once go-spiffe adds support
witSigningKeys, err := WITSigningKeysFromBundleProto(b)
if err != nil {
	return nil, err
}
*/

func BundleProtoFromRootCA(trustDomainID string, rootCA *x509.Certificate) *common.Bundle {
	_ = "STUB: not implemented"
	return nil
}

func BundleProtoFromRootCAs(trustDomainID string, rootCAs []*x509.Certificate) *common.Bundle {
	_ = "STUB: not implemented"
	return nil
}

func RootCAsFromBundleProto(b *common.Bundle) (out []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func JWTSigningKeysFromBundleProto(b *common.Bundle) (map[string]crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WITSigningKeysFromBundleProto(b *common.Bundle) (map[string]crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MergeBundles(a, b *common.Bundle) (*common.Bundle, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// PruneBundle removes the bundle RootCAs and JWT keys that expired before a given time
// It returns an error if pruning results in a bundle with no CAs or keys
func PruneBundle(bundle *common.Bundle, expiration time.Time, log logrus.FieldLogger) (*common.Bundle, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Zero value is a valid time, but probably unintended

// Creates new bundle with non expired certs only

// if any cert in the chain has expired, throw the whole chain out

// FindX509Authorities search for all X.509 authorities with provided subjectKeyIDs
func FindX509Authorities(bundle *spiffebundle.Bundle, subjectKeyIDs []string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getX509Authority(bundle *spiffebundle.Bundle, subjectKeyID string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneBundle(b *common.Bundle) *common.Bundle { _ = "STUB: not implemented"; return nil }
