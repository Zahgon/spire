package x509

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewX509PrepareCommand creates a new "x509 prepare" subcommand for "localauthority" command.
func NewX509PrepareCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewX509PrepareCommandWithEnv creates a new "x509 prepare" subcommand for "localauthority" command
// using the environment specified
func NewX509PrepareCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type x509PrepareCommand struct {
	printer cliprinter.Printer
	env     *commoncli.Env
}

func (c *x509PrepareCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*x509PrepareCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *x509PrepareCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority x509 prepare` CLI command
func (c *x509PrepareCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintX509Prepare(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
