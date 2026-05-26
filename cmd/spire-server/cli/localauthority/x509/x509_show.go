package x509

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewShowCommand creates a new "x509 show" subcommand for "localauthority" command.
func NewX509ShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewX509ShowCommandWithEnv creates a new "x509 show" subcommand for "localauthority" command
// using the environment specified
func NewX509ShowCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type x509ShowCommand struct {
	printer cliprinter.Printer

	env *commoncli.Env
}

func (c *x509ShowCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*x509ShowCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *x509ShowCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority x509 show` CLI command
func (c *x509ShowCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintX509Show(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
