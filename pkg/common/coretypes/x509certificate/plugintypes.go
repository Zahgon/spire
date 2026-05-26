package x509certificate

import (
	"crypto/x509"

	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromPluginProto(pb *plugintypes.X509Certificate) (*X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FromPluginProtos(pbs []*plugintypes.X509Certificate) ([]*X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProto(x509Authority *X509Authority) (*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProtos(x509Authorities []*X509Authority) ([]*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCommonProtos(pbs []*common.Certificate) ([]*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCertificates(x509Certificates []*x509.Certificate) ([]*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCertificate(x509Certificate *x509.Certificate) (*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProto(pb *apitypes.X509Certificate) (*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProtos(pbs []*apitypes.X509Certificate) ([]*plugintypes.X509Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
