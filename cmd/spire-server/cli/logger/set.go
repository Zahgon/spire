package logger

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type setCommand struct {
	env      *commoncli.Env
	newLevel string
	printer  cliprinter.Printer
}

// Returns a cli.command that sets the log level using the default
// cli environment.
func NewSetCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Returns a cli.command that sets the log level.
func NewSetCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

// The name of the command.
func (*setCommand) Name() string { _ = "STUB: not implemented"; return "" }

// The help presented description of the command.
func (*setCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Adds additional flags specific to the command.
func (c *setCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

// The routine that executes the command
func (c *setCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *setCommand) prettyPrintLogger(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
