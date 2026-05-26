package jwt

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewJWTPrepareCommand creates a new "jwt prepare" subcommand for "localauthority" command.
func NewJWTPrepareCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewJWTPrepareCommandWithEnv creates a new "jwt prepare" subcommand for "localauthority" command
// using the environment specified
func NewJWTPrepareCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type jwtPrepareCommand struct {
	printer cliprinter.Printer
	env     *commoncli.Env
}

func (c *jwtPrepareCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*jwtPrepareCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *jwtPrepareCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority jwt prepare` CLI command
func (c *jwtPrepareCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintJWTPrepare(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
