package entry

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
	selectors StringsFlag

	// Workload parent spiffeID
	parentID string

	// Workload spiffeID
	spiffeID string

	// Entry hint
	hint string

	// List of SPIFFE IDs of trust domains the registration entry is federated with
	federatesWith StringsFlag

	// Whether the entry is for a downstream SPIRE server
	downstream bool

	// Match used when filtering by federates with
	matchFederatesWithOn string

	// Match used when filtering by selectors
	matchSelectorsOn string

	printer cliprinter.Printer
	env     *commoncli.Env
}

// NewCountCommand creates a new "count" subcommand for "entry" command.
func NewCountCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewCountCommandWithEnv creates a new "count" subcommand for "entry" command
// using the environment specified.
func NewCountCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*countCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*countCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

// Run counts attested entries
func (c *countCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *countCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *countCommand) prettyPrintCount(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
