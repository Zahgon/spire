package agent

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type evictCommand struct {
	env *commoncli.Env
	// SPIFFE ID of the agent being evicted
	spiffeID string
	printer  cliprinter.Printer
}

// NewEvictCommand creates a new "evict" subcommand for "agent" command.
func NewEvictCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewEvictCommandWithEnv creates a new "evict" subcommand for "agent" command
// using the environment specified
func NewEvictCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*evictCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*evictCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run evicts an agent given its SPIFFE ID
func (c *evictCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *evictCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintEvictResult(env *commoncli.Env, _ ...any) error {
	_ = "STUB: not implemented"
	return nil
}
