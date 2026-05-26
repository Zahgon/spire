//go:build windows

package run

import (
	"flag"
	"net"

	"github.com/spiffe/spire/pkg/agent"
)

func (c *agentConfig) addOSFlags(flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *agentConfig) setPlatformDefaults() { _ = "STUB: not implemented"; return }

func (c *agentConfig) getAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *agentConfig) getAdminAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *agentConfig) hasAdminAddr() bool { _ = "STUB: not implemented"; return false }

// validateOS performs windows specific validations of the agent config
func (c *agentConfig) validateOS() error { _ = "STUB: not implemented"; return nil }

func prepareEndpoints(*agent.Config) error {
	_ = "STUB: not implemented"
	// Nothing to do in this platform
	return nil
}
