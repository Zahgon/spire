package plugintest

import (
	"testing"

	"github.com/spiffe/spire/pkg/common/catalog"
)

type config struct {
	builtInConfig catalog.BuiltInConfig

	serviceFacades []catalog.Facade

	loadErr *error

	doConfigure  bool
	configureErr *error
	coreConfig   catalog.CoreConfig
	plainConfig  *string
	jsonConfig   any
}

func (conf *config) makeConfigData(t *testing.T) string { _ = "STUB: not implemented"; return "" }
