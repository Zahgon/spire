package middleware

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

type EntryFetcher interface {
	// FetchEntries fetches the downstream entries matching the given SPIFFE ID.
	FetchEntries(ctx context.Context, id spiffeid.ID) ([]*types.Entry, error)
}

// EntryFetcherFunc implements EntryFetcher with a function
type EntryFetcherFunc func(ctx context.Context, id spiffeid.ID) ([]*types.Entry, error)

// FetchEntries fetches the downstream entries matching the given SPIFFE ID.
func (fn EntryFetcherFunc) FetchEntries(ctx context.Context, id spiffeid.ID) ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type callerEntriesKey struct{}

// WithCallerEntries returns the caller entries retrieved using the given
// fetcher. If the context already has the caller entries, they are returned
// without re-fetching. This reduces entry fetching in the face of multiple
// authorizers.
func WithCallerEntries(ctx context.Context, entryFetcher EntryFetcher) (context.Context, []*types.Entry, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}
