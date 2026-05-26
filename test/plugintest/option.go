package plugintest

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"github.com/spiffe/spire/pkg/common/catalog"
)

// Option is plugin test option
type Option interface {
	setOption(config *config)
}

type optionFunc func(conf *config)

func (fn optionFunc) setOption(conf *config) {
	_ = "STUB: not implemented"

	// Log sets the logger for the plugin.
	return
}

func Log(log logrus.FieldLogger) Option { _ = "STUB: not implemented"; return *new(Option) }

// Services sets the services also implemented by the plugin.
func Services(serviceFacades ...catalog.Facade) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// HostServices sets the host services the host will offer to the plugin.
func HostServices(hostServices ...pluginsdk.ServiceServer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// CoreConfig provides the core configuration passed to the plugin when
// configured.
func CoreConfig(coreConfig catalog.CoreConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Configure provides raw configuration to the plugin for configuration.
func Configure(plainConfig string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Configuref provides a formatted string to the plugin for configuration.
func Configuref(format string, args ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

// ConfigureJSON marshals the given object and passes the resulting JSON to
// the plugin for configuration.
func ConfigureJSON(jsonConfig any) Option { _ = "STUB: not implemented"; return *new(Option) }

// CaptureLoadError captures the error encountered during loading. If loading
// fails, and this option is not provided, the test will fail.
func CaptureLoadError(errp *error) Option { _ = "STUB: not implemented"; return *new(Option) }

// CaptureLoadError captures the error encountered during configuration. If
// configuration fails, and this option is not provided, the test will fail.
func CaptureConfigureError(errp *error) Option { _ = "STUB: not implemented"; return *new(Option) }
