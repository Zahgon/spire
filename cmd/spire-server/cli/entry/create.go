package entry

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	serverutil "github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewCreateCommand creates a new "create" subcommand for "entry" command.
func NewCreateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newCreateCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type createCommand struct {
	// Path to an optional data file. If set, other
	// opts will be ignored.
	path string

	// Type and value are delimited by a colon (:)
	// ex. "unix:uid:1000" or "spiffe_id:spiffe://example.org/foo"
	selectors StringsFlag

	// Registration entry ID
	entryID string

	// Workload parent spiffeID
	parentID string

	// Workload spiffeID
	spiffeID string

	// Entry hint, used to disambiguate entries with the same SPIFFE ID
	hint string

	// TTL for x509 SVIDs issued to this workload
	x509SVIDTTL int

	// TTL for JWT SVIDs issued to this workload
	jwtSVIDTTL int

	// List of SPIFFE IDs of trust domains the registration entry is federated with
	federatesWith StringsFlag

	// whether the registration entry is for an "admin" workload
	admin bool

	// whether the entry is for a downstream SPIRE server
	downstream bool

	// whether the entry represents a node or group of nodes
	node bool

	// Expiry of entry
	entryExpiry int64

	// DNSNames entries for SVIDs based on this entry
	dnsNames StringsFlag

	// storeSVID determines if the issued SVID must be stored through an SVIDStore plugin
	storeSVID bool

	// disableX509SVIDPrefetch tells the agent not to prefetch and cache X509 SVID for
	// the given entry
	disableX509SVIDPrefetch bool

	printer cliprinter.Printer

	env *commoncli.Env
}

func (*createCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*createCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *createCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *createCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// validate performs basic validation, even on fields that we
// have defaults defined for.
func (c *createCommand) validate() (err error) {
	_ = "STUB: not implemented"
	// If a path is set, we have all we need
	return nil
}

// parseConfig builds a registration entry from the given config
func (c *createCommand) parseConfig() ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createEntries(ctx context.Context, c entryv1.EntryClient, entries []*types.Entry) (resp *entryv1.BatchCreateEntryResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Entry API does not include in the results the entries that
// failed to be created, so we populate them from the request data.

func getParentID(config *createCommand, td string) (*types.SPIFFEID, error) {
	_ = "STUB: not implemented"
	// If the node flag is set, then set the Parent ID to the server's expected SPIFFE ID
	return nil, nil
}

func prettyPrintCreate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}
