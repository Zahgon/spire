package x509certificate

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromCommonProto(pb *common.Certificate) (*X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FromCommonProtos(pbs []*common.Certificate) ([]*X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProto(x509Authority *X509Authority) (*common.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProtos(x509Authorities []*X509Authority) ([]*common.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonFromPluginProtos(pbs []*plugintypes.X509Certificate) ([]*common.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
