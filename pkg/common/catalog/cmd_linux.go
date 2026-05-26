package catalog

import (
	"os/exec"
)

func pluginCmd(name string, arg ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

// This is insurance that a plugin process does not outlive SPIRE on linux.
