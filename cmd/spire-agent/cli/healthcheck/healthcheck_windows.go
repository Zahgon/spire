//go:build windows

package healthcheck

import (
	"flag"
	"net"
)

// healthCheckCommandOS has windows specific implementation
// that complements healthCheckCommand
type healthCheckCommandOS struct {
	namedPipeName string
}

func (c *healthCheckCommandOS) addOSFlags(flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *healthCheckCommandOS) getAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}
