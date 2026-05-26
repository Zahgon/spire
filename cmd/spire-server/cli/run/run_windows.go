//go:build windows

package run

import (
	"flag"
	"net"
)

func (c *serverConfig) addOSFlags(flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *serverConfig) getAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *serverConfig) setDefaultsIfNeeded() { _ = "STUB: not implemented"; return }

// validateOS performs OS specific validations of the server config
func (c *Config) validateOS() error { _ = "STUB: not implemented"; return nil }
