package catalog

import (
	"context"
	"net"
	"sync"
)

type pipeAddr struct{}

func (pipeAddr) Network() string { _ = "STUB: not implemented"; return "" }
func (pipeAddr) String() string  { _ = "STUB: not implemented"; return "" }

type pipeNet struct {
	accept    chan net.Conn
	closed    chan struct{}
	closeOnce sync.Once
}

func newPipeNet() *pipeNet { _ = "STUB: not implemented"; return nil }

func (n *pipeNet) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (n *pipeNet) Accept() (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

func (n *pipeNet) DialContext(ctx context.Context, _ string) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (n *pipeNet) Close() error { _ = "STUB: not implemented"; return nil }
