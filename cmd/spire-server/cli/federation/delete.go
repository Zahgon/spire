package federation

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewDeleteCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newDeleteCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type deleteCommand struct {
	// SPIFFE ID of the trust domain to delete
	id      string
	env     *commoncli.Env
	printer cliprinter.Printer
}

func (c *deleteCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *deleteCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *deleteCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *deleteCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintDelete(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
