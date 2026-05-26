package jwt

import (
	"context"
	"flag"
	"time"

	"github.com/mitchellh/cli"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewMintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newMintCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type mintCommand struct {
	spiffeID string
	ttl      time.Duration
	audience commoncli.StringsFlag
	write    string
	env      *commoncli.Env
	printer  cliprinter.Printer
}

func (c *mintCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *mintCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *mintCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *mintCommand) Run(ctx context.Context, env *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Print in stdout

// Save in file

func (c *mintCommand) validateToken(token string, env *commoncli.Env) error {
	_ = "STUB: not implemented"
	return nil
}

func getJWTSVIDEndOfLife(token string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ttlToSeconds returns the number of seconds in a duration, rounded up to
// the nearest second
func ttlToSeconds(ttl time.Duration) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func prettyPrintMint(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
