//go:build windows

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
	// Access control is already handled by the security
	// descriptor associated with the named pipe.
	// Nothing else is needed to be done here.
	return nil
}

func prepareLocalAddr(net.Addr) error {
	_ = "STUB: not implemented"
	// Nothing to do in this platform
	return nil
}
