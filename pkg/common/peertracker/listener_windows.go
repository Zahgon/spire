//go:build windows

package peertracker

import (
	"net"

	"github.com/Microsoft/go-winio"
)

type ListenerFactoryOS struct {
	NewPipeListener func(pipe string, pipeConfig *winio.PipeConfig) (net.Listener, error)
}

func (lf *ListenerFactory) ListenPipe(pipe string, pipeConfig *winio.PipeConfig) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lf *ListenerFactory) listenPipe(pipe string, pipeConfig *winio.PipeConfig) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
