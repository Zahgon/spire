//go:build !windows

package peertracker

import (
	"net"
)

func getCallerInfoFromNamedPipeConn(net.Conn) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}
