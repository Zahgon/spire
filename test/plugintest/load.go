package plugintest

import (
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type Plugin interface {
	catalog.Configurer
	io.Closer
}

// Load loads a built-in plugin for testing with the given options. The plugin
// facade can be nil. If one of the Configure* options is given, the plugin
// will also be configured. The plugin will unload when the test is over. The
// function returns a plugin interface that can be closed to unload the
// built-in before the test is finished or used to reconfigure the plugin, but
// can otherwise be ignored.
func Load(t *testing.T, builtIn catalog.BuiltIn, pluginFacade catalog.Facade, options ...Option) Plugin {
	_ = "STUB: not implemented"
	return *new(Plugin)
}

func testLogger(t *testing.T) logrus.FieldLogger {
	_ = "STUB: not implemented"
	return *new(logrus.FieldLogger)
}

type logHook struct{ t *testing.T }

func (h logHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (h logHook) Fire(e *logrus.Entry) error { _ = "STUB: not implemented"; return nil }
