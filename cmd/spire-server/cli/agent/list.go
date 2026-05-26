package agent

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type listCommand struct {
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

// NewListCommand creates a new "list" subcommand for "agent" command.
func NewListCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewListCommandWithEnv creates a new "list" subcommand for "agent" command
// using the environment specified
func NewListCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*listCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*listCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run lists attested agents
func (c *listCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse the time string into a time.Time object

// 0: all, 1: can't reattest, 2: can reattest

// 0: all, 1: no-banned, 2: banned

// comfortably under the (4 MB/theoretical maximum size of 1 agent in MB)

func (c *listCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func prettyPrintAgents(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func printAgents(env *commoncli.Env, agents ...*types.Agent) error {
	_ = "STUB: not implemented"
	return nil
}

// Banned agents will have an empty serial number

func parseToSelectorMatch(match string) (types.SelectorMatch_MatchBehavior, error) {
	_ = "STUB: not implemented"
	return *new(types.SelectorMatch_MatchBehavior), nil
}
