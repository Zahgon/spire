package plugin

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/catalog"
	"google.golang.org/grpc/codes"
)

// PrefixMessage prefixes the given message with plugin information. The prefix
// is only applied if it is not already applied.
func PrefixMessage(pluginInfo catalog.PluginInfo, message string) string {
	_ = "STUB: not implemented"
	return ""
}

// Facade is embedded by plugin interface facade implementations as a
// convenient way to embed PluginInfo but also provide a set of convenient
// functions for embellishing and generating errors that have the plugin
// name prefixed.
type Facade struct {
	catalog.PluginInfo
	Log logrus.FieldLogger
}

// FixedFacade is a helper that creates a facade from fixed information, i.e.
// not the product of a loaded plugin.
func FixedFacade(pluginName, pluginType string, log logrus.FieldLogger) Facade {
	_ = "STUB: not implemented"
	return *new(Facade)
}

// InitInfo partially satisfies the catalog.Facade interface
func (f *Facade) InitInfo(pluginInfo catalog.PluginInfo) { _ = "STUB: not implemented"; return }

// InitLog partially satisfies the catalog.Facade interface
func (f *Facade) InitLog(log logrus.FieldLogger) {
	_ = "STUB: not implemented"

	// WrapErr wraps a given error such that it will be prefixed with the plugin
	// name. This method should be used by facade implementations to wrap errors
	// that come out of plugin implementations.
	return
}

func (f *Facade) WrapErr(err error) error { _ = "STUB: not implemented"; return nil }

// Embellish the gRPC status with the prefix, if necessary.

// Care must be taken to preserve any status details. Therefore, the
// proto is embellished directly and a new status created from that
// proto.

// Embellish the normal error with the prefix, if necessary. This is a
// defensive measure since plugins go over gRPC.

// Error creates a gRPC status with the given code and message. The message
// will be prefixed with the plugin name.
func (f *Facade) Error(code codes.Code, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Errorf creates a gRPC status with the given code and
// formatted message. The message will be prefixed with the plugin name.
func (f *Facade) Errorf(code codes.Code, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func prefixMessage(pluginInfo catalog.PluginInfo, message string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func messagePrefix(pluginInfo catalog.PluginInfo) string { _ = "STUB: not implemented"; return "" }

type facadeError struct {
	wrapped error
	message string
}

func (e *facadeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *facadeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type pluginInfo struct {
	pluginName string
	pluginType string
}

func (info pluginInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (info pluginInfo) Type() string { _ = "STUB: not implemented"; return "" }
