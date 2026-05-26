package token

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	prototypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewGenerateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newGenerateCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type generateCommand struct {
	// Optional SPIFFE ID to create with the token
	SpiffeID string

	// Token TTL in seconds
	TTL     int
	env     *commoncli.Env
	printer cliprinter.Printer
}

func (g *generateCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (g *generateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (g *generateCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func getID(spiffeID string) (*prototypes.SPIFFEID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *generateCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (g *generateCommand) prettyPrintGenerate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
