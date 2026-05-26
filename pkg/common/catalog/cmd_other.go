//go:build !linux

package catalog

import (
	"os/exec"
)

func pluginCmd(name string, arg ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }
