package federation

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewRefreshCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newRefreshCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type refreshCommand struct {
	id      string
	env     *commoncli.Env
	printer cliprinter.Printer
}

func (c *refreshCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *refreshCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *refreshCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *refreshCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintRefresh(env *commoncli.Env, _ ...any) error { _ = "STUB: not implemented"; return nil }
