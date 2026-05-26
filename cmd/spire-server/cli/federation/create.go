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

const (
	profileHTTPSWeb    = "https_web"
	profileHTTPSSPIFFE = "https_spiffe"
)

// NewCreateCommand creates a new "create" subcommand for "federation" command.
func NewCreateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newCreateCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type createCommand struct {
	path                    string
	config                  *federationRelationshipConfig
	env                     *commoncli.Env
	printer                 cliprinter.Printer
	federationRelationships []*types.FederationRelationship
}

func (*createCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*createCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *createCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *createCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *createCommand) prettyPrintCreate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Process results

// The trust domain API does not include in the results the relationships that
// failed to be created, so we populate them from the request data.

// Print federation relationships that succeeded to be created

// Print federation relationships that failed to be created
