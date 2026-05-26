//go:build !windows

package peertracker

import "net"

type ListenerFactoryOS struct {
	NewUnixListener func(network string, laddr *net.UnixAddr) (*net.UnixListener, error)
}

func (lf *ListenerFactory) ListenUnix(network string, laddr *net.UnixAddr) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lf *ListenerFactory) listenUnix(network string, laddr *net.UnixAddr) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
