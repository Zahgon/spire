package jwt

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewJWTTaintCommand creates a new "jwt taint" subcommand for "localauthority" command.
func NewJWTTaintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewJWTTaintCommandWithEnv creates a new "jwt taint" subcommand for "localauthority" command
// using the environment specified
func NewJWTTaintCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newJWTTaintCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type jwtTaintCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *jwtTaintCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*jwtTaintCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *jwtTaintCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority jwt taint` CLI command
func (c *jwtTaintCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintJWTTaint(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *jwtTaintCommand) validate() error { _ = "STUB: not implemented"; return nil }
