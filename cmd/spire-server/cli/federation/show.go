package federation

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newShowCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type showCommand struct {
	// Trust domain name of the federation relationship to show
	trustDomain string
	env         *commoncli.Env
	printer     cliprinter.Printer
}

func (c *showCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *showCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *showCommand) prettyPrintShow(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
