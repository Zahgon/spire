//go:build !windows

package common

import (
	"flag"
	"net"
)

type ConfigOS struct {
	socketPath string
	instance   string
}

func (c *ConfigOS) AddOSFlags(flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *ConfigOS) GetAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *ConfigOS) GetTargetName() (string, error) { _ = "STUB: not implemented"; return "", nil }
