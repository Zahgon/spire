package agent

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type countCommand struct {
	// Type and value are delimited by a colon (:)
	// ex. "unix:uid:1000" or "spiffe_id:spiffe://example.org/foo"
	selectors commoncli.StringsFlag

	// Match used when filtering by selectors
	matchSelectorsOn string

	// Filters agents to those that are banned.
	banned commoncli.BoolFlag

	// Filters agents by those that expire before this value.
	expiresBefore string

	// Filters agents to those matching the attestation type.
	attestationType string

	// Filters agents that can re-attest.
	canReattest commoncli.BoolFlag

	env *commoncli.Env

	printer cliprinter.Printer
}

// NewCountCommand creates a new "count" subcommand for "agent" command.
func NewCountCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewCountCommandWithEnv creates a new "count" subcommand for "agent" command
// using the environment specified.
func NewCountCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*countCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*countCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run counts attested agents
func (c *countCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse the time string into a time.Time object

// 0: all, 1: can't reattest, 2: can reattest

// 0: all, 1: no-banned, 2: banned

func (c *countCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintCount(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
