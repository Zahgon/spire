//go:build !windows

package spireplugin

import (
	"net"
)

func (p *Plugin) getWorkloadAPIAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}
