package bundle

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

const (
	deleteBundleRestrict   = "restrict"
	deleteBundleDissociate = "dissociate"
	deleteBundleDelete     = "delete"
)

// NewDeleteCommand creates a new "delete" subcommand for "bundle" command.
func NewDeleteCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newDeleteCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type deleteCommand struct {
	env *commoncli.Env
	// SPIFFE ID of the trust domain bundle
	id string
	// Deletion mode
	mode string
	// Command printer
	printer cliprinter.Printer
}

func (c *deleteCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *deleteCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *deleteCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *deleteCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintDelete(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteModeFromFlag(mode string) (bundlev1.BatchDeleteFederatedBundleRequest_Mode, error) {
	_ = "STUB: not implemented"
	return *new(bundlev1.BatchDeleteFederatedBundleRequest_Mode), nil
}
