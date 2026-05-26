package x509

import (
	"bytes"
	"context"
	"crypto"
	"flag"
	"time"

	"github.com/mitchellh/cli"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type generateKeyFunc func() (crypto.Signer, error)

func NewMintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newMintCommand(env *commoncli.Env, generateKey generateKeyFunc) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type mintCommand struct {
	generateKey generateKeyFunc

	spiffeID string
	ttl      time.Duration
	dnsNames commoncli.StringsFlag
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

// ttlToSeconds returns the number of seconds in a duration, rounded up to
// the nearest second
func ttlToSeconds(ttl time.Duration) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

type mintResult struct {
	X509SVID   [][]byte `json:"x509_svid"`
	PrivateKey []byte   `json:"private_key"`
	RootCAs    [][]byte `json:"root_cas"`
}

func (c *mintCommand) prettyPrintMint(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func convertSVIDResultToPEM(privateKey []byte, svidCertChain, rootCAs [][]byte) (*bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
