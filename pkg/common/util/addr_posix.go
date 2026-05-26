//go:build !windows

package util

import (
	"net"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
)

func NewGRPCClient(target string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetWorkloadAPIClientOption(addr net.Addr) (workloadapi.ClientOption, error) {
	_ = "STUB: not implemented"
	return *new(workloadapi.ClientOption), nil
}
