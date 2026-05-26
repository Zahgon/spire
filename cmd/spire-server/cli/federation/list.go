package federation

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewListCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newListCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type listCommand struct {
	env     *commoncli.Env
	printer cliprinter.Printer
}

func (c *listCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *listCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *listCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *listCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintList(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
