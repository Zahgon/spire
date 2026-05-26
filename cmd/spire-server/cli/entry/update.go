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

// NewUpdateCommand creates a new "update" subcommand for "entry" command.
func NewUpdateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newUpdateCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type updateCommand struct {
	// Path to an optional data file. If set, other
	// opts will be ignored.
	path string

	// Registration entry id to update
	entryID string

	// Type and value are delimited by a colon (:)
	// ex. "unix:uid:1000" or "spiffe_id:spiffe://example.org/foo"
	selectors StringsFlag

	// Workload parent spiffeID
	parentID string

	// Workload spiffeID
	spiffeID string

	// whether the entry is for a downstream SPIRE server
	downstream bool

	// TTL for x509 SVIDs issued to this workload
	x509SvidTTL int

	// TTL for JWT SVIDs issued to this workload
	jwtSvidTTL int

	// List of SPIFFE IDs of trust domains the registration entry is federated with
	federatesWith StringsFlag

	// whether the registration entry is for an "admin" workload
	admin bool

	// Expiry of entry
	entryExpiry int64

	// DNSNames entries for SVIDs based on this entry
	dnsNames StringsFlag

	// storeSVID determines if the issued SVID must be stored through an SVIDStore plugin
	storeSVID bool

	// additionalAttributesSet indicates whether any of the additional attributes if the requestration entry were set
	additionalAttributesSet bool

	// disableX509SVIDPrefetch tells the agent not to prefetch and cache X509 SVID for
	// the given entry
	disableX509SVIDPrefetch bool

	// Entry hint, used to disambiguate entries with the same SPIFFE ID
	hint string

	printer cliprinter.Printer

	env *commoncli.Env
}

func (*updateCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*updateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *updateCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *updateCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient serverutil.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// validate performs basic validation, even on fields that we
// have defaults defined for
func (c *updateCommand) validate() (err error) {
	_ = "STUB: not implemented"
	// If a path is set, we have all we need
	return nil
}

// parseConfig builds a registration entry from the given config
func (c *updateCommand) parseConfig() ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateEntries(ctx context.Context, c entryv1.EntryClient, entries []*types.Entry) (resp *entryv1.BatchUpdateEntryResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Entry API does not include in the results the entries that
// failed to be updated, so we populate them from the request data.

func prettyPrintUpdate(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Print entries that succeeded to be updated

// Print entries that failed to be updated
