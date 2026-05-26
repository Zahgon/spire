//go:build windows

package util

import (
	"context"
	"flag"
	"net"
)

type adapterOS struct {
	namedPipeName string
}

func (a *Adapter) addOSFlags(flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func dialer(ctx context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	// This is an ugly workaround to circumvent grpc-go needing us to provide the resolver in the address
	// in order to bypass DNS lookup, which is not relevant in the case of CLI invocation.
	return *new(net.Conn), nil
}

func (a *Adapter) getGRPCAddr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// When grpc-go deprecated grpc.DialContext() in favor of grpc.NewClient(),
// they made a breaking change to always use the DNS resolver, even when overriding the context dialer.
// This is problematic for clients that do not use DNS for address resolution and don't set a resolver in the address.
// As a workaround, use the passthrough resolver to prevent using the DNS resolver.
// More context can be found in this issue: https://github.com/grpc/grpc-go/issues/1786#issuecomment-2114124036
