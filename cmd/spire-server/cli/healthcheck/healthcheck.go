package healthcheck

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
)

func NewHealthCheckCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newHealthCheckCommand(env *common_cli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type healthCheckCommand struct {
	shallow bool
	verbose bool
}

func (c *healthCheckCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *healthCheckCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *healthCheckCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *healthCheckCommand) Run(ctx context.Context, env *common_cli.Env, client util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *healthCheckCommand) run(ctx context.Context, env *common_cli.Env, client util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore error since a failure to write to stderr cannot very well
// be reported
