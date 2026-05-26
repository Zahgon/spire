//go:build !windows

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

func StartGRPCServer(t *testing.T, registerFn func(s *grpc.Server)) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func ServeGRPCServerOnListener(t *testing.T, server *grpc.Server, listener net.Listener) {
	_ = "STUB: not implemented"
	return
}
