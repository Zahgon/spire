package catalog

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"google.golang.org/grpc"
)

const (
	deinitTimeout = 10 * time.Second
)

// Plugin is a loaded plugin.
type Plugin interface {
	// Closer is used to unload the plugin. Any facades initialized by the
	// call to bind are considered invalidated after the plugin is closed.
	io.Closer

	// Bind binds the given facades to the plugin. It also returns a Configurer
	// that can be used to configure the plugin. If the plugin does not support
	// a given facade, an error will be returned. This function is designed
	// only for use by unit-tests for built-in plugin implementations or fake
	// facade implementations that rely on built-ins.
	Bind(facades ...Facade) (Configurer, error)
}

type pluginImpl struct {
	closerGroup

	conn             grpc.ClientConnInterface
	info             PluginInfo
	log              logrus.FieldLogger
	grpcServiceNames []string
}

func newPlugin(ctx context.Context, conn grpc.ClientConnInterface, info PluginInfo, log logrus.FieldLogger, closers closerGroup, hostServices []pluginsdk.ServiceServer) (*pluginImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bind implements the Plugin interface method of the same name.
func (p *pluginImpl) Bind(facades ...Facade) (Configurer, error) {
	_ = "STUB: not implemented"
	return *new(Configurer), nil
}

func (p *pluginImpl) bindFacade(repo bindable, facade Facade) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *pluginImpl) initFacade(facade Facade) any { _ = "STUB: not implemented"; return *new(any) }

func (p *pluginImpl) bindRepos(pluginRepo bindablePluginRepo, serviceRepos []bindableServiceRepo) (Configurer, error) {
	_ = "STUB: not implemented"
	return *new(Configurer), nil
}

func (p *pluginImpl) makeConfigurer(grpcServiceNames map[string]struct{}) (Configurer, error) {
	_ = "STUB: not implemented"
	return *new(Configurer), nil
}

func (p *pluginImpl) bindRepo(repo bindableServiceRepo, grpcServiceNames map[string]struct{}) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Use the first matching version (in case the plugin implements
// more than one). The rest will be removed from the list of
// service names above so we can properly warn of unhandled
// services without false negatives.

func warnIfDeprecated(log logrus.FieldLogger, thisVersion, latestVersion Version) {
	_ = "STUB: not implemented"
	return
}

func grpcServiceNameSet(grpcServiceNames []string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func sortStringSet(set map[string]struct{}) []string { _ = "STUB: not implemented"; return nil }
