package middleware

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/authpolicy"
)

func (m *authorizationMiddleware) opaAuth(ctx context.Context, req any, fullMethod string) (context.Context, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), false, nil
}

// Get SPIFFE ID

func (m *authorizationMiddleware) reconcileResult(ctx context.Context, res authpolicy.Result) (context.Context, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), false, nil
}

// Check things in order of cost

// Check local

// Check statically configured admin entries

// Check entry-based admin and downstream auth

func isAdminViaConfig(ctx context.Context, adminIDs map[spiffeid.ID]struct{}) (context.Context, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), false
}

func isAdminViaEntries(ctx context.Context, entries []*types.Entry) (context.Context, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), false
}

func isDownstreamViaEntries(ctx context.Context, entries []*types.Entry) (context.Context, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), false
}

func isAgent(ctx context.Context, agentAuthorizer AgentAuthorizer) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func setAuthorizationLogFields(ctx context.Context, as, via string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
