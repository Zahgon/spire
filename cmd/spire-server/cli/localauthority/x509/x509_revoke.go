package x509

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewX509ActivateCommand creates a new "x509 revoke" subcommand for "localauthority" command.
func NewX509RevokeCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewX509ActivateCommandWithEnv creates a new "x509 revoke" subcommand for "localauthority" command
// using the environment specified
func NewX509RevokeCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type x509RevokeCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *x509RevokeCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*x509RevokeCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *x509RevokeCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority x509 revoke` CLI command
func (c *x509RevokeCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *x509RevokeCommand) validate() error { _ = "STUB: not implemented"; return nil }

func prettyPrintX509Revoke(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
