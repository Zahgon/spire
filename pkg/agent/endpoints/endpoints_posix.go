//go:build !windows

package endpoints

import (
	"net"
)

func (e *Endpoints) createUDSListener() (net.Listener, error) {
	_ = "STUB: not implemented"
	// Remove uds if already exists
	return *new(net.Listener), nil
}

func (e *Endpoints) createListener() (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
