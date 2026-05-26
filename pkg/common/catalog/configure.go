package catalog

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
)

type CoreConfig struct {
	TrustDomain spiffeid.TrustDomain
}

func (c CoreConfig) v1() *configv1.CoreConfiguration { _ = "STUB: not implemented"; return nil }

type Configurer interface {
	Configure(ctx context.Context, coreConfig CoreConfig, configuration string) error
	Validate(ctx context.Context, coreConfig CoreConfig, configuration string) (*configv1.ValidateResponse, error)
}

type ConfigurerFunc func(ctx context.Context, coreConfig CoreConfig, configuration string) error
type ValidatorFunc func(ctx context.Context, coreConfig CoreConfig, configuration string) (*configv1.ValidateResponse, error)

func (fn ConfigurerFunc) Configure(ctx context.Context, coreConfig CoreConfig, configuration string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn ValidatorFunc) Validate(ctx context.Context, coreConfig CoreConfig, configuration string) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConfigurePlugin(ctx context.Context, coreConfig CoreConfig, configurer Configurer, dataSource DataSource, lastHash string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ReconfigureTask(log logrus.FieldLogger, reconfigurer Reconfigurer) func(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type Reconfigurer interface {
	Reconfigure(ctx context.Context)
}

type Reconfigurers []Reconfigurer

func (rs Reconfigurers) Reconfigure(ctx context.Context) { _ = "STUB: not implemented"; return }

type Reconfigurable struct {
	Log        logrus.FieldLogger
	CoreConfig CoreConfig
	Configurer Configurer
	DataSource DataSource
	LastHash   string
}

func (r *Reconfigurable) Reconfigure(ctx context.Context) { _ = "STUB: not implemented"; return }

func configurePlugin(ctx context.Context, pluginLog logrus.FieldLogger, coreConfig CoreConfig, configurer Configurer, dataSource DataSource) (Reconfigurer, error) {
	_ = "STUB: not implemented"
	return *new(Reconfigurer), nil
}

// The plugin doesn't support configuration and no data source was configured. Nothing to do.

// The plugin does not support configuration but a data source was configured. This is a failure.

// The plugin supports configuration but no data source was configured. Default to an empty, fixed configuration.

// The plugin supports configuration and there was a data source.

type configurerRepo struct {
	configurer Configurer
}

func (repo *configurerRepo) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *configurerRepo) Versions() []Version { _ = "STUB: not implemented"; return nil }

func (repo *configurerRepo) Clear() {
	_ = "STUB: not implemented"
	// This function is only for conforming to the Repo interface and isn't
	// expected to be called, but just in case, we'll do the right thing
	// and clear out the configurer that has been bound.
	return
}

type configurerV1Version struct{}

func (configurerV1Version) New() Facade      { _ = "STUB: not implemented"; return *new(Facade) }
func (configurerV1Version) Deprecated() bool { _ = "STUB: not implemented"; return false }

type configurerV1 struct {
	configv1.ConfigServiceClient
}

var _ Configurer = (*configurerV1)(nil)

func (v1 *configurerV1) InitInfo(PluginInfo) { _ = "STUB: not implemented"; return }

func (v1 *configurerV1) InitLog(logrus.FieldLogger) { _ = "STUB: not implemented"; return }

func (v1 *configurerV1) Configure(ctx context.Context, coreConfig CoreConfig, hclConfiguration string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v1 *configurerV1) Validate(ctx context.Context, coreConfig CoreConfig, hclConfiguration string) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hashData(data string) string { _ = "STUB: not implemented"; return "" }
