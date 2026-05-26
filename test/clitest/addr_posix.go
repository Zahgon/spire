//go:build !windows

package clitest

import (
	"net"
)

func GetAddr(addr net.Addr) string { _ = "STUB: not implemented"; return "" }
