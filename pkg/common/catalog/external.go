package catalog

import (
	"context"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"google.golang.org/grpc"
)

type externalConfig struct {
	// Name of the plugin
	Name string

	// Type is the plugin type (e.g. KeyManager)
	Type string

	// Path is the path on disk to the plugin.
	Path string

	// Args are the command line arguments to supply to the plugin
	Args []string

	// Checksum is the hex-encoded SHA256 hash of the plugin binary.
	Checksum string

	// Log is the logger to be wired to the external plugin.
	Log logrus.FieldLogger

	// HostServices are the host service servers provided to the plugin.
	HostServices []pluginsdk.ServiceServer
}

func loadExternal(ctx context.Context, config externalConfig) (*pluginImpl, error) {
	_ = "STUB: not implemented"
	// TODO: honor context cancellation... unfortunately go-plugin doesn't seem
	// to give us a mechanism for this, so we'd have to spin up some goroutine
	// to watch for cancellation and start killing clients and closing
	// connections and the like.
	return nil, nil
}

// Resolve path to an absolute path. We don't want to rely on PATH
// environment lookups for security reasons.

// Start the external plugin.

// TODO: Enable AutoMTLS if it is fixed to work with brokering.
// See https://github.com/hashicorp/go-plugin/issues/109

// Ensure the loaded plugin is killed if there is a failure.

// Create the GRPC client and ensure it is closed on error.

// Dispense the client, which invokes the GRPCClient method in the
// hcClientPlugin. The result of that method call is returned here, which
// is coerced back into the correct type.

// Purely defensive. This should never happen since we control what
// gets returned from hcClientPlugin.

// Plugin has been loaded and initialized. Ensure the plugin client is
// killed when the plugin is closed.

type hcClientPlugin struct {
	goplugin.NetRPCUnsupportedPlugin

	config externalConfig
}

var _ goplugin.GRPCPlugin = (*hcClientPlugin)(nil)

func (p *hcClientPlugin) GRPCServer(*goplugin.GRPCBroker, *grpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *hcClientPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	_ = "STUB: not implemented"
	// Manually start up the server via b.Accept since b.AcceptAndServe does
	// some logging we don't care for. Although b.AcceptAndServe is currently
	// the only way to feed the TLS config to the brokered connection, AutoMTLS
	// does not work yet anyway, so it is a moot point.
	return *new(any), nil
}

//nolint:gosec // G118: false complaint about cancel not being called

type hcPlugin struct {
	conn    grpc.ClientConnInterface
	closers closerGroup
}

func buildSecureConfig(checksum string) (*goplugin.SecureConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
