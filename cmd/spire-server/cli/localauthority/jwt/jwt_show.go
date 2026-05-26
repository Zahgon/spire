package jwt

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewJWTShowCommand creates a new "jwt show" subcommand for "localauthority" command.
func NewJWTShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewJWTShowCommandWithEnv creates a new "jwt show" subcommand for "localauthority" command
// using the environment specified
func NewJWTShowCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type jwtShowCommand struct {
	printer cliprinter.Printer

	env *commoncli.Env
}

func (c *jwtShowCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*jwtShowCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *jwtShowCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority jwt show` CLI command
func (c *jwtShowCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintJWTShow(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
