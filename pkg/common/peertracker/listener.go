package peertracker

import (
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

var _ net.Listener = &Listener{}

type ListenerFactory struct {
	Log               logrus.FieldLogger
	NewTracker        func(log logrus.FieldLogger) (PeerTracker, error)
	ListenerFactoryOS // OS specific
}

type Listener struct {
	l       net.Listener
	log     logrus.FieldLogger
	Tracker PeerTracker
}

func newNoopLogger() *logrus.Logger { _ = "STUB: not implemented"; return nil }

func (l *Listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Support future Listener types

func (l *Listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *Listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type closeOnIsAliveErr struct {
	Watcher
	conn io.Closer
}

func (w closeOnIsAliveErr) IsAlive() error { _ = "STUB: not implemented"; return nil }
