package bundle

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"

	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

type countCommand struct {
	env     *commoncli.Env
	printer cliprinter.Printer
}

// NewCountCommand creates a new "count" subcommand for "bundle" command.
func NewCountCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewCountCommandWithEnv creates a new "count" subcommand for "bundle" command
// using the environment specified.
func NewCountCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*countCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*countCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run counts attested bundles
func (c *countCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *countCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintCount(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
