//go:build !windows

package main

import (
	"net"
)

func (c *Config) getWorkloadAPIAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *Config) getServerAPITargetName() string { _ = "STUB: not implemented"; return "" }

// validateOS performs os specific validations of the configuration
func (c *Config) validateOS() (err error) { _ = "STUB: not implemented"; return nil }

func listenLocal(c *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
