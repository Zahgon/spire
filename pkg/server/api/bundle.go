package api

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func BundleToProto(b *common.Bundle) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CertificatesToProto(rootCas []*common.Certificate) []*types.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}

func PublicKeysToJWTKeys(keys []*common.PublicKey) []*types.JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func PublicKeysToWITKeys(keys []*common.PublicKey) []*types.WITKey {
	_ = "STUB: not implemented"
	return nil
}

func ProtoToBundle(b *types.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProtoToBundleMask(mask *types.BundleMask) *common.BundleMask {
	_ = "STUB: not implemented"
	return nil
}

func ParseX509Authorities(certs []*types.X509Certificate) ([]*common.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseJWTAuthorities(keys []*types.JWTKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HashByte(b []byte) string { _ = "STUB: not implemented"; return "" }

func FieldsFromBundleProto(proto *types.Bundle, inputMask *types.BundleMask) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func FieldsFromJwtAuthoritiesProto(jwtAuthorities []*types.JWTKey) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func FieldsFromX509AuthoritiesProto(x509Authorities []*types.X509Certificate) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}
