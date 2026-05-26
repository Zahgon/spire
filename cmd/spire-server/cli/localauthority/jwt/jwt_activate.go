package jwt

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewJWTActivateCommand creates a new "jwt activate" subcommand for "localauthority" command.
func NewJWTActivateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewJWTActivateCommandWithEnv creates a new "jwt activate" subcommand for "localauthority" command
// using the environment specified
func NewJWTActivateCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type jwtActivateCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *jwtActivateCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*jwtActivateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *jwtActivateCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority jwt activate` CLI command
func (c *jwtActivateCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *jwtActivateCommand) validate() error { _ = "STUB: not implemented"; return nil }

func prettyPrintJWTActivate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
