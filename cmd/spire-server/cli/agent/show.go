package agent

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type showCommand struct {
	env *commoncli.Env
	// SPIFFE ID of the agent being shown
	spiffeID string
	printer  cliprinter.Printer
}

// NewShowCommand creates a new "show" subcommand for "agent" command.
func NewShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewShowCommandWithEnv creates a new "show" subcommand for "agent" command
// using the environment specified
func NewShowCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*showCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*showCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run shows an agent given its SPIFFE ID
func (c *showCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *showCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintAgent(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
