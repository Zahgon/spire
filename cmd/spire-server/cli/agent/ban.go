package agent

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type banCommand struct {
	env *commoncli.Env
	// SPIFFE ID of agent being banned
	spiffeID string
	printer  cliprinter.Printer
}

// NewBanCommand creates a new "ban" subcommand for "agent" command.
func NewBanCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewBanCommandWithEnv creates a new "ban" subcommand for "agent" command
// using the environment specified
func NewBanCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*banCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*banCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run ban an agent given its SPIFFE ID
func (c *banCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *banCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintBanResult(env *commoncli.Env, _ ...any) error {
	_ = "STUB: not implemented"
	return nil
}
