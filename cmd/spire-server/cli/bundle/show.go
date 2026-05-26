package bundle

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewShowCommand creates a new "show" subcommand for "bundle" command.
func NewShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newShowCommand(env *common_cli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type showCommand struct {
	env          *common_cli.Env
	bundleFormat string
	printer      cliprinter.Printer
}

func (c *showCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *showCommand) Run(ctx context.Context, _ *common_cli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *showCommand) prettyPrintBundle(env *common_cli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
