package jwt

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewJWTActivateCommand creates a new "jwt revoke" subcommand for "localauthority" command.
func NewJWTRevokeCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewJWTActivateCommandWithEnv creates a new "jwt revoke" subcommand for "localauthority" command
// using the environment specified
func NewJWTRevokeCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type jwtRevokeCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *jwtRevokeCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*jwtRevokeCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *jwtRevokeCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority jwt revoke` CLI command
func (c *jwtRevokeCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *jwtRevokeCommand) validate() error { _ = "STUB: not implemented"; return nil }

func prettyPrintJWTRevoke(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
