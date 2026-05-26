package peertracker

import (
	"net"
)

type Conn struct {
	net.Conn
	Info AuthInfo
}

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }
