//go:build !windows

package endpoints

import (
	"net"

	"github.com/spiffe/spire/pkg/common/peertracker"
)

func (e *Endpoints) listen() (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (e *Endpoints) listenWithAuditLog() (*peertracker.Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Endpoints) restrictLocalAddr() error {
	_ = "STUB: not implemented"
	// Restrict access to the UDS to processes running as the same user or
	// group as the server.
	return nil
}

func prepareLocalAddr(localAddr net.Addr) error { _ = "STUB: not implemented"; return nil }
