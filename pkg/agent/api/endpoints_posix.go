//go:build !windows

package api

import (
	"net"
)

func (e *Endpoints) createListener() (net.Listener, error) {
	_ = "STUB: not implemented"
	// Remove uds if already exists
	return *new(net.Listener), nil
}
