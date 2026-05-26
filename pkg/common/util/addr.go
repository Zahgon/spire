package util

import (
	"net"
)

// GetUnixAddr returns a unix address with the designated
// path. Path is converted to an absolute path when constructing
// the returned unix domain socket address.
func GetUnixAddrWithAbsPath(path string) (*net.UnixAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetUnixAddr(name string) *net.UnixAddr { _ = "STUB: not implemented"; return nil }

// GetTargetName gets the fully qualified, self-contained name used
// for gRPC channel construction. Supported networks are unix and tcp.
// Unix paths must be absolute.
func GetTargetName(addr net.Addr) (string, error) { _ = "STUB: not implemented"; return "", nil }
