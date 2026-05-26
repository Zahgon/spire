//go:build !windows

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

// validateOS performs posix specific validations of the agent config
func (c *agentConfig) validateOS() error { _ = "STUB: not implemented"; return nil }

func prepareEndpoints(c *agent.Config) error {
	_ = "STUB: not implemented"
	// Create uds dir and parents if not exists
	return nil
}

// Set umask before starting up the agent

// Create uds dir and parents if not exists
