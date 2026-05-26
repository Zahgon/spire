package rpccontext

import (
	"context"
	"crypto/x509"
	"net"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

type callerAddrKey struct{}
type callerIDKey struct{}
type callerX509SVIDKey struct{}
type callerDownstreamEntriesKey struct{}
type callerAdminTagKey struct{}
type callerLocalTagKey struct{}
type callerAgentTagKey struct{}

// WithCallerAddr returns a context with the given address.
func WithCallerAddr(ctx context.Context, addr net.Addr) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerAddr returns the caller address.
func CallerAddr(ctx context.Context) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// WithCallerID returns a context with the given ID.
func WithCallerID(ctx context.Context, id spiffeid.ID) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerID returns the caller ID, if available.
func CallerID(ctx context.Context) (spiffeid.ID, bool) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), false
}

// WithCallerX509SVID returns a context with the given X509SVID.
func WithCallerX509SVID(ctx context.Context, x509SVID *x509.Certificate) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerX509SVID returns the caller X509SVID, if available.
func CallerX509SVID(ctx context.Context) (*x509.Certificate, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// WithCallerDownstreamEntries returns a context with the given entries.
func WithCallerDownstreamEntries(ctx context.Context, entries []*types.Entry) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerDownstreamEntries returns the downstream entries for the caller. If the caller is not
// a downstream caller, it returns false.
func CallerDownstreamEntries(ctx context.Context) ([]*types.Entry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// CallerIsDownstream returns true if the caller is a downstream caller.
func CallerIsDownstream(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// WithAdminCaller returns a context where the caller is tagged as an admin.
func WithAdminCaller(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerIsAdmin returns true if the caller is an admin.
func CallerIsAdmin(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// WithLocalCaller returns a context where the caller is tagged as local.
func WithLocalCaller(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerIsLocal returns true if the caller is local.
func CallerIsLocal(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// WithAgentCaller returns a context where the caller is tagged as an agent.
func WithAgentCaller(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerIsAgent returns true if the caller is an agent.
func CallerIsAgent(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
