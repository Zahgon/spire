//go:build windows

package namedpipe

import (
	"net"
)

type Addr struct {
	serverName string
	pipeName   string
}

func (p *Addr) PipeName() string { _ = "STUB: not implemented"; return "" }

func (p *Addr) Network() string { _ = "STUB: not implemented"; return "" }

func (p *Addr) String() string { _ = "STUB: not implemented"; return "" }

// AddrFromName returns a named pipe in the local
// computer with the specified pipe name
func AddrFromName(pipeName string) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func GetPipeName(addr string) string { _ = "STUB: not implemented"; return "" }
