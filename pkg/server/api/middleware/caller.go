package middleware

import (
	"context"

	"google.golang.org/grpc/peer"
)

func callerContextFromContext(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func tcpCallerContextFromPeer(ctx context.Context, p *peer.Peer) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// No TLS information. Return an unauthenticated TCP caller.

// The connection state unfortunately does not have VerifiedChains set
// because SPIFFE TLS does custom verification, i.e., Go's TLS stack only
// sets VerifiedChains if it is the one to verify the chain of trust.

// No certificates. Return an unauthenticated TCP caller.
