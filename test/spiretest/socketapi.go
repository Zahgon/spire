package spiretest

import (
	"net"
	"testing"

	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"google.golang.org/grpc"
)

func StartWorkloadAPIOnUDSSocket(t *testing.T, socketPath string, server workload.SpiffeWorkloadAPIServer) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func StartGRPCUDSSocketServer(t *testing.T, socketPath string, registerFn func(s *grpc.Server)) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func ServeGRPCServerOnTempUDSSocket(t *testing.T, server *grpc.Server) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func ServeGRPCServerOnUDSSocket(t *testing.T, server *grpc.Server, socketPath string) net.Addr {
	_ = "STUB: not implemented"
	// ensure the directory holding the socket exists
	return *new(net.Addr)
}
