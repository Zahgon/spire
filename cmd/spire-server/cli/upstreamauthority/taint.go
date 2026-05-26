package upstreamauthority

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewTaintCommand creates a new "upstreamauthority taint" subcommand for "upstreamauthority" command.
func NewTaintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewUpstreamauthorityTaintCommandWithEnv creates a new "upstreamauthority taint" subcommand for "upstreamauthority" command
// using the environment specified
func NewTaintCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newTaintCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type upstreamauthorityTaintCommand struct {
	subjectKeyID string
	printer      cliprinter.Printer
	env          *commoncli.Env
}

func (c *upstreamauthorityTaintCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*upstreamauthorityTaintCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *upstreamauthorityTaintCommand) AppendFlags(f *flag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

// Run executes all logic associated with a single invocation of the
// `spire-server upstreamauthority taint` CLI command
func (c *upstreamauthorityTaintCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintTaint(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *upstreamauthorityTaintCommand) validate() error { _ = "STUB: not implemented"; return nil }
