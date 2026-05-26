package api

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewValidateJWTCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newValidateJWTCommand(env *commoncli.Env, clientMaker workloadClientMaker) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type validateJWTCommand struct {
	audience string
	svid     string
	env      *commoncli.Env
	printer  cliprinter.Printer
}

func (*validateJWTCommand) name() string { _ = "STUB: not implemented"; return "" }

func (*validateJWTCommand) synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *validateJWTCommand) appendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *validateJWTCommand) run(ctx context.Context, _ *commoncli.Env, client *workloadClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *validateJWTCommand) validateJWTSVID(ctx context.Context, client *workloadClient) (*workload.ValidateJWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prettyPrintValidate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
