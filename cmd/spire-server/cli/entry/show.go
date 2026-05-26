package entry

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

const listEntriesRequestPageSize = 500

// NewShowCommand creates a new "show" subcommand for "entry" command.
func NewShowCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newShowCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type showCommand struct {
	// Type and value are delimited by a colon (:)
	// ex. "unix:uid:1000" or "spiffe_id:spiffe://example.org/foo"
	selectors StringsFlag

	// ID of the entry to be shown
	entryID string

	// Workload parent spiffeID
	parentID string

	// Workload spiffeID
	spiffeID string

	// Entry hint
	hint string

	// List of SPIFFE IDs of trust domains the registration entry is federated with
	federatesWith StringsFlag

	// whether the entry is for a downstream SPIRE server
	downstream bool

	// Match used when filtering by federates with
	matchFederatesWithOn string

	// Match used when filtering by selectors
	matchSelectorsOn string

	printer cliprinter.Printer

	env *commoncli.Env
}

func (c *showCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*showCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server entry show` CLI command
func (c *showCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// validate ensures that the values in showCommand are valid
func (c *showCommand) validate() error {
	_ = "STUB: not implemented"
	// If entryID is given, it should be the only constraint
	return nil
}

func (c *showCommand) fetchEntries(ctx context.Context, client entryv1.EntryClient) (*entryv1.ListEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If an Entry ID was specified, look it up directly

// fetchByEntryID uses the configured EntryID to fetch the appropriate registration entry
func (c *showCommand) fetchByEntryID(ctx context.Context, id string, client entryv1.EntryClient) (*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printEntries(entries []*types.Entry, env *commoncli.Env) { _ = "STUB: not implemented"; return }

func parseToSelectorMatch(match string) (types.SelectorMatch_MatchBehavior, error) {
	_ = "STUB: not implemented"
	return *new(types.SelectorMatch_MatchBehavior), nil
}

func parseToFederatesWithMatch(match string) (types.FederatesWithMatch_MatchBehavior, error) {
	_ = "STUB: not implemented"
	return *new(types.FederatesWithMatch_MatchBehavior), nil
}

func prettyPrintShow(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
