package entry

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
)

const defaultEntryPageSize = 500

// Config defines the service configuration.
type Config struct {
	TrustDomain   spiffeid.TrustDomain
	EntryFetcher  api.AuthorizedEntryFetcher
	DataStore     datastore.DataStore
	EntryPageSize int
}

// Service defines the v1 entry service.
type Service struct {
	entryv1.UnsafeEntryServer

	td            spiffeid.TrustDomain
	ds            datastore.DataStore
	ef            api.AuthorizedEntryFetcher
	entryPageSize int
}

// New creates a new v1 entry service.
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// RegisterService registers the entry service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// CountEntries returns the total number of entries.
func (s *Service) CountEntries(ctx context.Context, req *entryv1.CountEntriesRequest) (*entryv1.CountEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListEntries returns the optionally filtered and/or paginated list of entries.
func (s *Service) ListEntries(ctx context.Context, req *entryv1.ListEntriesRequest) (*entryv1.ListEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEntry returns the registration entry associated with the given SpiffeID
func (s *Service) GetEntry(ctx context.Context, req *entryv1.GetEntryRequest) (*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BatchCreateEntry adds one or more entries to the server.
func (s *Service) BatchCreateEntry(ctx context.Context, req *entryv1.BatchCreateEntryRequest) (*entryv1.BatchCreateEntryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) createEntry(ctx context.Context, e *types.Entry, outputMask *types.EntryMask) *entryv1.BatchCreateEntryResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// BatchUpdateEntry updates one or more entries in the server.
func (s *Service) BatchUpdateEntry(ctx context.Context, req *entryv1.BatchUpdateEntryRequest) (*entryv1.BatchUpdateEntryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BatchDeleteEntry removes one or more entries from the server.
func (s *Service) BatchDeleteEntry(ctx context.Context, req *entryv1.BatchDeleteEntryRequest) (*entryv1.BatchDeleteEntryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) deleteEntry(ctx context.Context, id string) *entryv1.BatchDeleteEntryResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthorizedEntries returns the list of entries authorized for the caller ID in the context.
func (s *Service) GetAuthorizedEntries(ctx context.Context, req *entryv1.GetAuthorizedEntriesRequest) (*entryv1.GetAuthorizedEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SyncAuthorizedEntries returns the list of entries authorized for the caller ID in the context.
func (s *Service) SyncAuthorizedEntries(stream entryv1.Entry_SyncAuthorizedEntriesServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Emit "success" auditing if we succeed.

func SyncAuthorizedEntries(stream entryv1.Entry_SyncAuthorizedEntriesServer, entries []api.ReadOnlyEntry, entryPageSize int) (err error) {
	_ = "STUB: not implemented"
	// Receive the initial request with the output mask.
	return nil
}

// There is no reason we couldn't support filtering by ID on the initial
// response but there doesn't seem to be a reason to. For now, fail if
// the initial request has IDs set.

// The revision number should probably have never been included in the
// entry mask. In any case, it is required to allow the caller to determine
// if it needs to ask for the full entry, so disallow masking here.

// Apply output mask to entries. The output mask field will be
// intentionally ignored on subsequent requests.

// If the number of entries is less than or equal to the entry page size,
// then just send the full list back. Otherwise, we'll send a sparse list
// and then stream back full entries as requested.

// Prepopulate the entry page used in the response with empty entry structs.
// These will be reused for each sparse entry response.

// Now wait for the client to request IDs that they need the full copy of.
// Each request is treated independently. Entries are paged back fully
// before the next request is received, using the More field as a flag to
// signal to the caller when all requested entries have been streamed back.

// EOF is normal and happens when the server processes the
// CloseSend sent by the client. If the client closes the stream
// before that point, then Canceled is expected. Either way, these
// conditions are normal and not an error.

// Sort the entries by ID for efficient lookups. This is done
// lazily since we only need these lookups if full copies are
// being requested.

// Sort the requested IDs for efficient lookups into the sorted entry
// list. Agents SHOULD already send the list sorted, but we need to
// make sure they are sorted for correctness of the search loop below.
// The go stdlib sorting algorithm performs well on pre-sorted data.

// Page back the requested entries. The slice for the entries in the response
// is reused to reduce memory pressure. Since both the entries and
// requested IDs are sorted, we can reduce the amount of entries we
// need to search as we iteratively move through the requested IDs.

// Adding the entry just found will exceed our page size.
// Ship the pageful of entries first and signal that there
// is more to follow.

// The response is either empty or contains a partial page. Either way
// we need to send what we have and signal there is no more to follow.

// fetchEntries fetches authorized entries using caller ID from context
func (s *Service) fetchEntries(ctx context.Context, log logrus.FieldLogger) ([]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyMask(e *types.Entry, mask *types.EntryMask) { _ = "STUB: not implemented"; return }

func (s *Service) updateEntry(ctx context.Context, e *types.Entry, inputMask *types.EntryMask, outputMask *types.EntryMask) *entryv1.BatchUpdateEntryResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

func fieldsFromEntryProto(ctx context.Context, proto *types.Entry, inputMask *types.EntryMask) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func fieldsFromListEntryFilter(ctx context.Context, td spiffeid.TrustDomain, filter *entryv1.ListEntriesRequest_Filter) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func fieldsFromCountEntryFilter(ctx context.Context, td spiffeid.TrustDomain, filter *entryv1.CountEntriesRequest_Filter) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func sortEntriesByID(entries []api.ReadOnlyEntry) { _ = "STUB: not implemented"; return }
