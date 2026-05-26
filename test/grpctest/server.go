package grpctest

import (
	"context"
	"sync"
	"testing"

	"github.com/spiffe/spire/pkg/common/api/middleware"
	"google.golang.org/grpc"
)

type ServerOption = func(*serverConfig)

type Server struct {
	dialTarget  string
	dialOptions []grpc.DialOption
	stop        func()
}

func (s *Server) NewGRPCClient(tb testing.TB, extraOptions ...grpc.DialOption) grpc.ClientConnInterface {
	_ = "STUB: not implemented"
	return *new(grpc.ClientConnInterface)
}

func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func StartServer(tb testing.TB, registerFn func(s grpc.ServiceRegistrar), opts ...ServerOption) *Server {
	_ = "STUB: not implemented"
	return nil
}

// Add the drain interceptors first so that they ensure all other handlers
// down the chain are complete before allowing the server to stop.

// Now add the context override so loggers or other things attached are
// available to subsequent interceptors.

// Now append the custom interceptors

// When grpc-go deprecated grpc.DialContext() in favor of grpc.NewClient(),
// they made a breaking change to always use the DNS resolver, even when overriding the context dialer.
// This is problematic for tests that rely on the grpc-go bufconn transport.
// grpc-go mentions that bufconn was only designed for internal testing of grpc-go, but we are relying on it in our tests.
// As a workaround, use the passthrough resolver to prevent using the DNS resolver,
// since the address is anyway being thrown away by the dialer method.
// More context can be found in this issue: https://github.com/grpc/grpc-go/issues/1786#issuecomment-2114124036

// Clean up when the test is closed.

// In case the test does not explicitly stop, do it on test cleanup.

type serverConfig struct {
	net                string
	addr               string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	contextOverride    func(context.Context) context.Context
}

func OverUDS() ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Middleware(ms ...middleware.Middleware) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func OverrideContext(fn func(context.Context) context.Context) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func unaryContextOverride(fn func(ctx context.Context) context.Context) func(context.Context, any, *grpc.UnaryServerInfo, grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return nil
}

func streamContextOverride(fn func(ctx context.Context) context.Context) func(any, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

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

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
