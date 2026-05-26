//go:build windows

package spiretest

import (
	"net"
	"testing"

	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"google.golang.org/grpc"
)

func StartWorkloadAPI(t *testing.T, server workload.SpiffeWorkloadAPIServer) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func StartWorkloadAPIOnNamedPipe(t *testing.T, pipeName string, server workload.SpiffeWorkloadAPIServer) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func StartGRPCServer(t *testing.T, registerFn func(s *grpc.Server)) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func StartGRPCOnNamedPipeServer(t *testing.T, pipeName string, registerFn func(s *grpc.Server)) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func ServeGRPCServerOnNamedPipe(t *testing.T, server *grpc.Server, pipeName string) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func ServeGRPCServerOnRandPipeName(t *testing.T, server *grpc.Server) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func GetRandNamedPipeAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func randUint64() uint64 { _ = "STUB: not implemented"; return 0 }

func ServeGRPCServerOnListener(t *testing.T, server *grpc.Server, listener net.Listener) {
	_ = "STUB: not implemented"
	// The Windows-specific implementation handles the race condition during cleanup
	// with named pipes. Named pipes on Windows require special handling during
	// shutdown to avoid deadlocks.
	return
}

// Close the listener first to unblock server.Serve()
// This is necessary on Windows to prevent deadlock during named pipe cleanup

// Wait for Serve to return. We expect an error since we closed the listener,
// so we don't need to check it.
