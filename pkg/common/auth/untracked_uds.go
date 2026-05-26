package auth

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

// UntrackedUDSCredentials returns credentials for UDS servers that rely solely
// on file permissions for access control. If the caller information (e.g. PID,
// UID, GID) is in any way used for further access control or authorization
// decisions, these credentials SHOULD NOT be used. The peertracker package
// should instead be used, which provides mitigation against PID reuse and
// related attacks.
func UntrackedUDSCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

type UntrackedUDSAuthInfo struct{}

func (UntrackedUDSAuthInfo) AuthType() string { _ = "STUB: not implemented"; return "" }

type untrackedUDSCredentials struct{}

func (c untrackedUDSCredentials) ClientHandshake(_ context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c untrackedUDSCredentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c untrackedUDSCredentials) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c untrackedUDSCredentials) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c untrackedUDSCredentials) OverrideServerName(_ string) error {
	_ = "STUB: not implemented"
	return nil
}
