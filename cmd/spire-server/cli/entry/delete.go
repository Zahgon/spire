package entry

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewDeleteCommand creates a new "delete" subcommand for "entry" command.
func NewDeleteCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newDeleteCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type deleteCommand struct {
	// ID of the record to delete
	entryID string
	file    string
	env     *commoncli.Env
	printer cliprinter.Printer
}

func (*deleteCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*deleteCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *deleteCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func parseEntryDeleteJSON(path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *deleteCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Perform basic validation.
func (c *deleteCommand) validate() error { _ = "STUB: not implemented"; return nil }

func (c *deleteCommand) prettyPrintDelete(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
