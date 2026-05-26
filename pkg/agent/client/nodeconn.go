package client

import (
	"sync"

	"google.golang.org/grpc"
)

type nodeConn struct {
	conn     *grpc.ClientConn
	refcount int32
	mu       sync.RWMutex
}

func newNodeConn(conn *grpc.ClientConn) *nodeConn { _ = "STUB: not implemented"; return nil }

func (c *nodeConn) AddRef() { _ = "STUB: not implemented"; return }

func (c *nodeConn) Release() { _ = "STUB: not implemented"; return }

func (c *nodeConn) Conn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }
