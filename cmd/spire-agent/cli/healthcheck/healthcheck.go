package healthcheck

import (
	"github.com/mitchellh/cli"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
)

func NewHealthCheckCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newHealthCheckCommand(env *common_cli.Env) *healthCheckCommand {
	_ = "STUB: not implemented"
	return nil
}

type healthCheckCommand struct {
	healthCheckCommandOS // os specific

	env *common_cli.Env

	shallow bool
	verbose bool
}

func (c *healthCheckCommand) Help() string {
	_ = "STUB: not implemented"
	// ignoring parsing errors since "-h" is always supported by the flags package
	return ""
}

func (c *healthCheckCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *healthCheckCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Ignore error since a failure to write to stderr cannot very well be
// reported

func (c *healthCheckCommand) parseFlags(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *healthCheckCommand) run() error { _ = "STUB: not implemented"; return nil }

// Ignore error since a failure to write to stderr cannot very well
// be reported
