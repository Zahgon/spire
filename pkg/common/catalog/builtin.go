package catalog

import (
	"context"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"google.golang.org/grpc"
)

type BuiltIn struct {
	Name     string
	Plugin   pluginsdk.PluginServer
	Services []pluginsdk.ServiceServer
}

func MakeBuiltIn(name string, pluginServer pluginsdk.PluginServer, serviceServers ...pluginsdk.ServiceServer) BuiltIn {
	_ = "STUB: not implemented"
	return *new(BuiltIn)
}

type BuiltInConfig struct {
	// Log is the logger to be wired to the external plugin.
	Log logrus.FieldLogger

	// HostServices are the host service servers provided to the plugin.
	HostServices []pluginsdk.ServiceServer
}

func LoadBuiltIn(ctx context.Context, builtIn BuiltIn, config BuiltInConfig) (_ Plugin, err error) {
	_ = "STUB: not implemented"
	return *new(Plugin), nil
}

func loadBuiltIn(ctx context.Context, builtIn BuiltIn, config BuiltInConfig) (_ *pluginImpl, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBuiltInServer(log logrus.FieldLogger) (*grpc.Server, io.Closer) {
	_ = "STUB: not implemented"
	return nil, *new(io.Closer)
}

type builtinDialer struct {
	pluginName   string
	log          logrus.FieldLogger
	hostServices []pluginsdk.ServiceServer
	conn         *pipeConn
}

func (d *builtinDialer) DialHost(context.Context) (grpc.ClientConnInterface, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientConnInterface), nil
}

func (d *builtinDialer) Close() error { _ = "STUB: not implemented"; return nil }

type pipeConn struct {
	grpc.ClientConnInterface
	io.Closer
}

func startPipeServer(server *grpc.Server, log logrus.FieldLogger) (_ *pipeConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dial the server

type drainHandlers struct {
	wg sync.WaitGroup
}

func (d *drainHandlers) Wait() { _ = "STUB: not implemented"; return }

func (d *drainHandlers) UnaryServerInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (d *drainHandlers) StreamServerInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}
