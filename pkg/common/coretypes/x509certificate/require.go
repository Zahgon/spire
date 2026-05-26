package x509certificate

import (
	"crypto/x509"

	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func RequireFromCommonProto(pb *common.Certificate) *X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func RequireFromCommonProtos(pbs []*common.Certificate) []*X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonProto(x509Certificate *X509Authority) *common.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonProtos(x509Certificates []*X509Authority) []*common.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonFromPluginProtos(pbs []*plugintypes.X509Certificate) []*common.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireFromPluginProto(pb *plugintypes.X509Certificate) *X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func RequireFromPluginProtos(pbs []*plugintypes.X509Certificate) []*X509Authority {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginProto(x509Certificate *X509Authority) *plugintypes.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginProtos(x509Certificates []*X509Authority) []*plugintypes.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCommonProtos(pbs []*common.Certificate) []*plugintypes.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCertificates(x509Certificates []*x509.Certificate) []*plugintypes.X509Certificate {
	_ = "STUB: not implemented"
	return nil
}

func panicOnError(err error) { _ = "STUB: not implemented"; return }
