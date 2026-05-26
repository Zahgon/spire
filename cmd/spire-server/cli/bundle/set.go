package bundle

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewSetCommand creates a new "set" subcommand for "bundle" command.
func NewSetCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newSetCommand(env *common_cli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type setCommand struct {
	env *common_cli.Env
	// SPIFFE ID of the trust bundle
	id string
	// Path to the bundle on disk (optional). If empty, reads from stdin.
	path         string
	bundleFormat string
	printer      cliprinter.Printer
}

func (c *setCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *setCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *setCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *setCommand) Run(ctx context.Context, env *common_cli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintSet(env *common_cli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
