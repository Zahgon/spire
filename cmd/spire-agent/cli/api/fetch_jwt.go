package api

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewFetchJWTCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newFetchJWTCommandWithEnv(env *commoncli.Env, clientMaker workloadClientMaker) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type fetchJWTCommand struct {
	audience commoncli.CommaStringsFlag
	spiffeID string
	printer  cliprinter.Printer
	env      *commoncli.Env
}

func (c *fetchJWTCommand) name() string { _ = "STUB: not implemented"; return "" }

func (c *fetchJWTCommand) synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *fetchJWTCommand) run(ctx context.Context, _ *commoncli.Env, client *workloadClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *fetchJWTCommand) appendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *fetchJWTCommand) fetchJWTSVID(ctx context.Context, client *workloadClient) (*workload.JWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *fetchJWTCommand) fetchJWTBundles(ctx context.Context, client *workloadClient) (*workload.JWTBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printPrettyResult(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
