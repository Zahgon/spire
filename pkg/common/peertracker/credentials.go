package peertracker

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

type grpcCredentials struct{}

func NewCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *grpcCredentials) ClientHandshake(_ context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *grpcCredentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *grpcCredentials) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c *grpcCredentials) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *grpcCredentials) OverrideServerName(_ string) error { _ = "STUB: not implemented"; return nil }

func WatcherFromContext(ctx context.Context) (Watcher, bool) {
	_ = "STUB: not implemented"
	return *new(Watcher), false
}

func AuthInfoFromContext(ctx context.Context) (AuthInfo, bool) {
	_ = "STUB: not implemented"
	return *new(AuthInfo), false
}
