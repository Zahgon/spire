package peertracker

import (
	"net"
)

func CallerFromNamedPipeConn(conn net.Conn) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}
