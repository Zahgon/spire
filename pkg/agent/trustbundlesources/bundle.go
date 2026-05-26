package trustbundlesources

import (
	"crypto/x509"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/agent/storage"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

type Bundle struct {
	config             *Config
	use                int
	connectionAttempts int
	startTime          time.Time
	log                logrus.FieldLogger
	metrics            telemetry.Metrics
	storage            storage.Storage
	lastBundle         []*x509.Certificate
}

func New(config *Config, log logrus.FieldLogger) Bundle {
	_ = "STUB: not implemented"
	return *new(Bundle)
}

func (b *Bundle) SetMetrics(metrics telemetry.Metrics) { _ = "STUB: not implemented"; return }

func (b *Bundle) SetStorage(sto storage.Storage) error { _ = "STUB: not implemented"; return nil }

func (b *Bundle) SetUse(use int) error { _ = "STUB: not implemented"; return nil }

func (b *Bundle) SetSuccessIfRunning() error { _ = "STUB: not implemented"; return nil }

func (b *Bundle) SetSuccess() error { _ = "STUB: not implemented"; return nil }

func (b *Bundle) SetForceRebootstrap() error { _ = "STUB: not implemented"; return nil }

func (b *Bundle) GetStartTime() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (b *Bundle) IsBootstrap() bool { _ = "STUB: not implemented"; return false }

func (b *Bundle) IsRebootstrap() bool { _ = "STUB: not implemented"; return false }

func (b *Bundle) GetBundle() ([]*x509.Certificate, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// If InsecureBootstrap is configured, the bundle is not required

func (b *Bundle) GetInsecureBootstrap() bool { _ = "STUB: not implemented"; return false }

func (b *Bundle) updateMetrics() { _ = "STUB: not implemented"; return }

func parseTrustBundle(bundleBytes []byte, trustBundleContentType string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func downloadTrustBundle(trustBundleURL string, trustBundleUnixSocket string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Download the trust bundle URL from the user specified URL
// We use gosec -- the annotation below will disable a security check that URLs are not tainted
/* #nosec G107 */

func loadTrustBundle(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
