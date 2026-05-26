package peertracker

import (
	"net"
)

func CallerFromUDSConn(conn net.Conn) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}
