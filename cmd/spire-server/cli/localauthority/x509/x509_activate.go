package x509

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewX509ActivateCommand creates a new "x509 activate" subcommand for "localauthority" command.
func NewX509ActivateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewX509ActivateCommandWithEnv creates a new "x509 activate" subcommand for "localauthority" command
// using the environment specified
func NewX509ActivateCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type x509ActivateCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *x509ActivateCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*x509ActivateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *x509ActivateCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority x509 activate` CLI command
func (c *x509ActivateCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *x509ActivateCommand) validate() error { _ = "STUB: not implemented"; return nil }

func prettyPrintX509Activate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
