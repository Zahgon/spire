package upstreamauthority

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewRevokeCommand creates a new "upstreamauthority revoke" subcommand for "upstreamauthority" command.
func NewRevokeCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewRevokeCommandWithEnv creates a new "upstreamauthority revoke" subcommand for "upstreamauthority" command
// using the environment specified
func NewRevokeCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newRevokeCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type upstreamauthorityRevokeCommand struct {
	subjectKeyID string
	printer      cliprinter.Printer
	env          *commoncli.Env
}

func (c *upstreamauthorityRevokeCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*upstreamauthorityRevokeCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *upstreamauthorityRevokeCommand) AppendFlags(f *flag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

// Run executes all logic associated with a single invocation of the
// `spire-server upstreamauthority revoke` CLI command
func (c *upstreamauthorityRevokeCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintRevoke(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *upstreamauthorityRevokeCommand) validate() error { _ = "STUB: not implemented"; return nil }
