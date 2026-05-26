package api

import (
	"context"
	"crypto"
	"crypto/x509"
	"flag"
	"time"

	"github.com/mitchellh/cli"
	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

func NewFetchX509Command() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newFetchX509Command(env *commoncli.Env, clientMaker workloadClientMaker) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type fetchX509Command struct {
	silent    bool
	writePath string
	env       *commoncli.Env
	printer   cliprinter.Printer
	respTime  time.Duration
}

func (*fetchX509Command) name() string { _ = "STUB: not implemented"; return "" }

func (*fetchX509Command) synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *fetchX509Command) run(ctx context.Context, _ *commoncli.Env, client *workloadClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *fetchX509Command) appendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *fetchX509Command) fetchX509SVID(ctx context.Context, client *workloadClient) (*workload.X509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *fetchX509Command) writeResponse(svids []*X509SVID) error {
	_ = "STUB: not implemented"
	return nil
}

// sort and write the keys by trust domain so the output is consistent

// writeCerts takes a slice of data, which may contain multiple certificates,
// and encodes them as PEM blocks, writing them to filename
func (c *fetchX509Command) writeCerts(filename string, certs []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// writeKey takes a private key, formats as PEM, and writes it to filename
func (c *fetchX509Command) writeKey(filename string, privateKey crypto.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

// writeFile creates or truncates filename, and writes data to it
func (c *fetchX509Command) writeFile(filename string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *fetchX509Command) prettyPrintFetchX509(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

type X509SVID struct {
	SPIFFEID         string
	Hint             string
	Certificates     []*x509.Certificate
	PrivateKey       crypto.Signer
	Bundle           []*x509.Certificate
	FederatedBundles map[string][]*x509.Certificate
}

func parseAndValidateX509SVIDResponse(resp *workload.X509SVIDResponse) ([]*X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseX509SVIDResponse(resp *workload.X509SVIDResponse) ([]*X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseX509SVID(svid *workload.X509SVID, federatedBundles map[string][]*x509.Certificate) (*X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateX509SVIDs(svids []*X509SVID) error { _ = "STUB: not implemented"; return nil }

func validateX509SVID(svid *X509SVID) error { _ = "STUB: not implemented"; return nil }
