package wit

import (
	"context"
	"crypto"
	"flag"
	"time"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
)

func NewMintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newMintCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

// Test helper function, to have control over the workload key that is being generated
func newMintCommandWithKeyGenerator(env *commoncli.Env, workloadKeyGenerator func() (crypto.Signer, error)) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type mintCommand struct {
	spiffeID             string
	keyType              string
	signingAlgorithm     string
	ttl                  time.Duration
	write                string
	env                  *commoncli.Env
	printer              cliprinter.Printer
	workloadKeyGenerator func() (crypto.Signer, error)
}

func (c *mintCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (c *mintCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *mintCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

type mintResult struct {
	PrivateKey string         `json:"private_key"`
	Svid       *types.WITSVID `json:"svid"`
}

func (c *mintCommand) Run(ctx context.Context, env *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Print in stdout

func (c *mintCommand) validateToken(token string, env *commoncli.Env) error {
	_ = "STUB: not implemented"
	return nil
}

func getWITSVIDEndOfLife(token string) (time.Time, error) {
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

func validateKeyTypeAndSigningAlgorithm(keyType string, signingAlgorithm string) (keymanager.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(keymanager.KeyType), nil
}
