package logger

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type getCommand struct {
	env     *commoncli.Env
	printer cliprinter.Printer
}

// Returns a cli.command that gets the logger information using
// the default cli environment.
func NewGetCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Returns a cli.command that gets the root logger information.
func NewGetCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

// The name of the command.
func (*getCommand) Name() string { _ = "STUB: not implemented"; return "" }

// The help presented description of the command.
func (*getCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Adds additional flags specific to the command.
func (c *getCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

// The routine that executes the command
func (c *getCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Formatting for the logger under pretty printing of output.
func (c *getCommand) prettyPrintLogger(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
