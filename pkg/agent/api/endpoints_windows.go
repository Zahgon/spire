//go:build windows

package api

import (
	"net"
)

func (e *Endpoints) createListener() (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
