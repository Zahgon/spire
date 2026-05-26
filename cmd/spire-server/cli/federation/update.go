package federation

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewUpdateCommand creates a new "update" subcommand for "federation" command.
func NewUpdateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newUpdateCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type updateCommand struct {
	path                    string
	config                  *federationRelationshipConfig
	env                     *commoncli.Env
	printer                 cliprinter.Printer
	federationRelationships []*types.FederationRelationship
}

func (*updateCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*updateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *updateCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *updateCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *updateCommand) prettyPrintUpdate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Process results

// The trust domain API does not include in the results the relationships that
// failed to be updated, so we populate them from the request data.

// Print federation relationships that succeeded to be updated

// Print federation relationships that failed to be updated
